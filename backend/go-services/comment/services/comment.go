/*
service层, 负责处理业务逻辑
*/
package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	redisUtil "github.com/BigNoseCattyHome/aorb/backend/utils/storage/redis"
	"github.com/redis/go-redis/v9"
	"time"

	commentModels "github.com/BigNoseCattyHome/aorb/backend/go-services/comment/models"
	pollModels "github.com/BigNoseCattyHome/aorb/backend/go-services/poll/models"
	commentPb "github.com/BigNoseCattyHome/aorb/backend/rpc/comment"
	"github.com/BigNoseCattyHome/aorb/backend/rpc/poll"
	"github.com/BigNoseCattyHome/aorb/backend/rpc/user"
	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/config"
	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/strings"
	"github.com/BigNoseCattyHome/aorb/backend/utils/extra/tracing"
	grpc2 "github.com/BigNoseCattyHome/aorb/backend/utils/grpc"
	"github.com/BigNoseCattyHome/aorb/backend/utils/logging"
	"github.com/BigNoseCattyHome/aorb/backend/utils/rabbitmq"
	"github.com/BigNoseCattyHome/aorb/backend/utils/storage/database"
	"github.com/BigNoseCattyHome/aorb/backend/utils/uuid"

	_ "github.com/mbobakov/grpc-consul-resolver"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var userClient user.UserServiceClient
var pollClient poll.PollServiceClient
var actionCommentLimitKeyPrefix = config.Conf.Redis.Prefix + "comment_freq_limit"

const actionCommentMaxQPS = 3 // Maximum ActionComment query amount of an actor per second

func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func actionCommentLimitKey(username string) string {
	return fmt.Sprintf("%s-%s", actionCommentLimitKeyPrefix, username)
}

type CommentServiceImpl struct {
	commentPb.CommentServiceServer
}

var conn *amqp.Connection
var channel *amqp.Channel

func (c CommentServiceImpl) New() {
	userRpcConn := grpc2.Connect(config.UserRpcServerName)
	userClient = user.NewUserServiceClient(userRpcConn)

	pollRpcConn := grpc2.Connect(config.PollRpcServerName)
	pollClient = poll.NewPollServiceClient(pollRpcConn)

	var err error
	conn, err = amqp.Dial(rabbitmq.BuildMqConnAddr())
	exitOnError(err)

	channel, err = conn.Channel()
	exitOnError(err)

	err = channel.ExchangeDeclare(
		strings.EventExchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	exitOnError(err)
}

func CloseMQConn() {
	if err := conn.Close(); err != nil {
		panic(err)
	}
	if err := channel.Close(); err != nil {
		panic(err)
	}
}

func (c CommentServiceImpl) ActionComment(ctx context.Context, request *commentPb.ActionCommentRequest) (resp *commentPb.ActionCommentResponse, err error) {
	ctx, span := tracing.Tracer.Start(ctx, "CommentService.ActionComment")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("CommentService.ActionComment").WithContext(ctx)
	logger.WithFields(logrus.Fields{
		"username":     request.Username,
		"poll_uuid":    request.PollUuid,
		"action_type":  request.ActionType,
		"comment_text": request.GetCommentText(),
		"comment_uuid": request.GetCommentUuid(),
	}).Debugf("Process start")

	var pCommentId string
	var pCommentText string

	switch request.ActionType {
	case commentPb.ActionCommentType_ACTION_COMMENT_TYPE_ADD:
		pCommentText = request.GetCommentText()
		break
	case commentPb.ActionCommentType_ACTION_COMMENT_TYPE_DELETE:
		pCommentId = request.GetCommentUuid()
		break
	case commentPb.ActionCommentType_ACTION_COMMENT_TYPE_UNSPECIFIED:
		fallthrough
	default:
		logger.Warnf("Invalid action type")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.ActionCommentTypeInvalidCode,
			StatusMsg:  strings.ActionCommentTypeInvalid,
		}
		return
	}

	// Check if poll exists
	pollExistResp, err := pollClient.PollExist(ctx, &poll.PollExistRequest{
		PollUuid: request.PollUuid,
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Query poll existence happens error")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.PollServiceInnerErrorCode,
			StatusMsg:  strings.PollServiceInnerError,
		}
		return
	}

	if !pollExistResp.Exist {
		logger.WithFields(logrus.Fields{
			"PollUuId": request.PollUuid,
		}).Errorf("Poll Uuid does not exist")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToQueryPollErrorCode,
			StatusMsg:  strings.UnableToQueryPollError,
		}
		return
	}

	//// get target user
	userResponse, err := userClient.GetUserInfo(ctx, &user.UserRequest{
		Username: request.Username,
	})

	if err != nil || userResponse.StatusCode != strings.ServiceOKCode {
		if userResponse.StatusCode == strings.UserNotExistedCode {
			resp = &commentPb.ActionCommentResponse{
				StatusCode: strings.UserNotExistedCode,
				StatusMsg:  strings.UserNotExisted,
			}
			return
		}
		logger.WithFields(logrus.Fields{
			"err":      err,
			"userName": request.Username,
		}).Errorf("User service error")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToQueryUserErrorCode,
			StatusMsg:  strings.UnableToQueryUserError,
		}
		return
	}

	switch request.ActionType {
	case commentPb.ActionCommentType_ACTION_COMMENT_TYPE_ADD:
		resp, err = addComment(ctx, logger, span, request.Username, request.PollUuid, pCommentText)
	case commentPb.ActionCommentType_ACTION_COMMENT_TYPE_DELETE:
		resp, err = deleteComment(ctx, logger, span, request.Username, request.PollUuid, pCommentId)
	}

	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":      err,
			"userName": request.Username,
		}).Errorf("Error when executing ActionComment Service")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToQueryUserErrorCode,
			StatusMsg:  strings.UnableToQueryUserError,
		}
		return
	}

	logger.WithFields(logrus.Fields{
		"response": resp,
	}).Debugf("Process done.")

	return
}

func (c CommentServiceImpl) ListComment(ctx context.Context, request *commentPb.ListCommentRequest) (resp *commentPb.ListCommentResponse, err error) {
	ctx, span := tracing.Tracer.Start(ctx, "ListCommentService")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("CommentService.ListComment").WithContext(ctx)
	logger.WithFields(logrus.Fields{
		"poll_uuid": request.PollUuid,
	}).Debugf("Process start")

	// TODO 添加redis
	// 设置redis键
	key := request.PollUuid

	// 从redis中获取数据
	redisResult, err := redisUtil.RedisClient.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error when getting data from redis")
		logging.SetSpanError(span, err)
		resp = &commentPb.ListCommentResponse{
			StatusCode: strings.PollServiceInnerErrorCode,
			StatusMsg:  strings.PollServiceInnerError,
		}
		return
	}

	if redisResult != "" {
		// 如果存在数据
		var responseCommentList []*commentPb.Comment
		err = json.Unmarshal([]byte(redisResult), &responseCommentList)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("Error when unmarshal data from redis")
			logging.SetSpanError(span, err)
			resp = &commentPb.ListCommentResponse{
				StatusCode: strings.PollServiceInnerErrorCode,
				StatusMsg:  strings.PollServiceInnerError,
			}
			return
		}
		resp = &commentPb.ListCommentResponse{
			StatusCode:  strings.ServiceOKCode,
			StatusMsg:   strings.ServiceOK,
			CommentList: responseCommentList,
		}
		logger.WithFields(logrus.Fields{
			"response": resp,
		}).Debugf("Process done.")
		return
	}

	// redis中不存在，查找数据库
	// 查看poll是否存在
	pollExistResp, err := pollClient.PollExist(ctx, &poll.PollExistRequest{
		PollUuid: request.PollUuid,
	})
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Query poll existence happens error")
		logging.SetSpanError(span, err)
		resp = &commentPb.ListCommentResponse{
			StatusCode: strings.PollServiceInnerErrorCode,
			StatusMsg:  strings.PollServiceInnerError,
		}
		return
	}

	if !pollExistResp.Exist {
		logger.WithFields(logrus.Fields{
			"poll_uuid": request.PollUuid,
		}).Errorf("Poll ID does not exist")
		logging.SetSpanError(span, err)
		resp = &commentPb.ListCommentResponse{
			StatusCode: strings.UnableToQueryPollErrorCode,
			StatusMsg:  strings.UnableToQueryPollError,
		}
		return
	}

	var pPoll pollModels.Poll
	collections := database.MongoDbClient.Database("aorb").Collection("polls")
	filter := bson.D{{"pollUuid", request.PollUuid}}
	cursor := collections.FindOne(ctx, filter)

	err = cursor.Decode(&pPoll)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error When Decoding Cursor of Poll")
		logging.SetSpanError(span, err)
		resp = &commentPb.ListCommentResponse{
			StatusCode: strings.UnableToQueryCommentErrorCode,
			StatusMsg:  strings.UnableToQueryCommentError,
		}
		return
	}

	pCommentList := pPoll.CommentList

	rCommentList := make([]*commentPb.Comment, 0)

	for _, comment := range pCommentList {
		rCommentList = append(rCommentList, BuildCommentPbModel(&comment))
	}

	resp = &commentPb.ListCommentResponse{
		StatusCode:  strings.ServiceOKCode,
		StatusMsg:   strings.ServiceOK,
		CommentList: rCommentList,
	}

	// 将数据存入redis
	jsonBytes, err := json.Marshal(&rCommentList)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error when marshalling commentList to json")
		logging.SetSpanError(span, err)
		resp = &commentPb.ListCommentResponse{
			StatusCode: strings.UnableToQueryCommentErrorCode,
			StatusMsg:  strings.UnableToQueryCommentError,
		}
		return
	}
	err = redisUtil.RedisClient.Set(ctx, key, string(jsonBytes), time.Hour).Err()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error when inserting commentList into redis")
		logging.SetSpanError(span, err)
		resp = &commentPb.ListCommentResponse{
			StatusCode: strings.UnableToQueryCommentErrorCode,
			StatusMsg:  strings.UnableToQueryCommentError,
		}
		return
	}

	logger.WithFields(logrus.Fields{
		"response": resp,
	}).Debugf("Process done.")

	return
}

func (c CommentServiceImpl) CountComment(ctx context.Context, request *commentPb.CountCommentRequest) (resp *commentPb.CountCommentResponse, err error) {
	ctx, span := tracing.Tracer.Start(ctx, "CountCommentService")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("CommentService.CountComment").WithContext(ctx)
	logger.WithFields(logrus.Fields{
		"poll_uuid": request.PollUuid,
	}).Debugf("Process start")

	// TODO 使用缓存
	//countStringKey := fmt.Sprintf("CommentCount-%s", request.PollUuid)
	//countString, err := cached.GetWithFunc(ctx, countStringKey,
	//	func(ctx context.Context, key string) (string, error) {
	//		rCount, err := count(ctx, request.PollUuid)
	//		return strconv.FormatInt(rCount, 10), err
	//	})

	collection := database.MongoDbClient.Database("aorb").Collection("polls")
	filter := bson.D{{"pollUuid", request.PollUuid}}
	cursor := collection.FindOne(ctx, filter)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":       err,
			"poll_uuid": request.PollUuid,
		}).Errorf("Error when searching poll")
		logging.SetSpanError(span, err)
		resp = &commentPb.CountCommentResponse{
			StatusCode:   strings.UnableToQueryCommentErrorCode,
			StatusMsg:    strings.UnableToQueryCommentError,
			CommentCount: 0,
		}
		return
	}

	var pPoll pollModels.Poll
	cursor.Decode(&pPoll)

	resp = &commentPb.CountCommentResponse{
		StatusCode:   strings.ServiceOKCode,
		StatusMsg:    strings.ServiceOK,
		CommentCount: uint32(len(pPoll.CommentList)),
	}
	logger.WithFields(logrus.Fields{
		"response": resp,
	}).Debugf("Process done.")
	return
}

func deleteComment(ctx context.Context, logger *logrus.Entry, span trace.Span, username string, pPollUuId string, commentUuid string) (resp *commentPb.ActionCommentResponse, err error) {
	pPoll := pollModels.Poll{}

	collections := database.MongoDbClient.Database("aorb").Collection("polls")

	filter := bson.D{
		{Key: "pollUuid", Value: pPollUuId},
		{"commentList.commentUuid", commentUuid},
	}

	// 先查询
	collections.FindOne(ctx, filter).Decode(&pPoll)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":          err,
			"poll_uuid":    pPollUuId,
			"comment_uuid": commentUuid,
		}).Errorf("Failed when searching for comment")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToQueryCommentErrorCode,
			StatusMsg:  strings.UnableToQueryCommentError,
		}
		return
	}

	flag := false

	for _, comment := range pPoll.CommentList {
		if comment.CommentUserName == username {
			flag = true
			break
		}
	}

	if flag == false {
		// 没找到对应提问
		logger.WithFields(logrus.Fields{
			"err":          err,
			"poll_uuid":    pPollUuId,
			"comment_uuid": commentUuid,
			"username":     username,
		}).Errorf("user information does not match")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToQueryCommentErrorCode,
			StatusMsg:  strings.UnableToQueryCommentError,
		}
		return
	}

	filter = bson.D{{"pollUuid", pPollUuId}}
	update := bson.D{
		{"$pull", bson.D{
			{"commentList", bson.D{
				{"commentUuid", commentUuid},
			}},
		}},
	}
	_, err = collections.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":          err,
			"poll_uuid":    pPollUuId,
			"comment_uuid": commentUuid,
			"username":     username,
		}).Errorf("Failed to delete comment")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToQueryCommentErrorCode,
			StatusMsg:  strings.UnableToQueryCommentError,
		}
		return
	}

	resp = &commentPb.ActionCommentResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		Comment:    nil,
	}
	return
}

func addComment(ctx context.Context, logger *logrus.Entry, span trace.Span, username string, pPollUuId string, pCommentText string) (resp *commentPb.ActionCommentResponse, err error) {

	collections := database.MongoDbClient.Database("aorb").Collection("polls")

	pComment := commentModels.Comment{
		CommentUuid:     uuid.GenerateUuid(),
		CommentUserName: username,
		Content:         pCommentText,
		CreateAt:        time.Now(),
	}

	newComment := bson.D{
		{"commentUuid", pComment.CommentUuid},
		{"commentUserName", pComment.CommentUserName},
		{"content", pComment.Content},
		{"createAt", pComment.CreateAt},
	}

	update := bson.D{
		{"$push", bson.D{
			{"commentList", newComment},
		}},
	}

	filter := bson.D{
		{"pollUuid", pPollUuId},
	}

	_, err = collections.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":       err,
			"poll_uuid": pPollUuId,
		}).Errorf("CommentService add comment action failed to response when adding comment")
		logging.SetSpanError(span, err)
		resp = &commentPb.ActionCommentResponse{
			StatusCode: strings.UnableToCreateCommentErrorCode,
			StatusMsg:  strings.UnableToCreateCommentError,
		}
		return
	}

	// TODO 一段意义不明的代码
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	//productComment(ctx, models.RecommendEvent{
	//	//	ActorId: pUser.Id,
	//	//	PollId:  []uint32{pPollId},
	//	//	Type:    2,
	//	//	Source:  config.CommentRpcServerName,
	//	//})
	//}()
	//wg.Wait()

	// TODO 后续可能会添加的 向user的userans中插入该提问的uuid
	//userCollection := database.MongoDbClient.Database("aorb").Collection("users")
	//filter4InsertPollUuid2UserAns := bson.D{
	//	{"username", username},
	//}
	//update4InsertPollUuid2Userans := bson.D{
	//	{"$push", bson.D{
	//		{"pollans.poll_ids", pPollUuId},
	//	}},
	//}
	//
	//_, err = userCollection.UpdateOne(ctx, filter4InsertPollUuid2UserAns, update4InsertPollUuid2Userans)
	//if err != nil {
	//	logger.WithFields(logrus.Fields{
	//		"err":       err,
	//		"poll_uuid": pPollUuId,
	//		"username":  username,
	//	}).Errorf("Error when inserting poll_uuid into user %s's pollansList", username)
	//	logging.SetSpanError(span, err)
	//	resp = &commentPb.ActionCommentResponse{
	//		StatusCode: strings.UnableToCreateCommentErrorCode,
	//		StatusMsg:  strings.UnableToCreateCommentError,
	//	}
	//	return
	//}

	resp = &commentPb.ActionCommentResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		Comment: &commentPb.Comment{
			CommentUuid:     pComment.CommentUuid,
			CommentUsername: username,
			Content:         pComment.Content,
			CreateAt:        timestamppb.New(pComment.CreateAt),
		},
	}
	return
}

func BuildCommentPbModel(comment *commentModels.Comment) *commentPb.Comment {
	return &commentPb.Comment{
		CommentUuid:     comment.CommentUuid,
		CommentUsername: comment.CommentUserName,
		Content:         comment.Content,
		CreateAt:        timestamppb.New(comment.CreateAt),
	}
}
