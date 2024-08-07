package services

import (
	"context"

	"github.com/BigNoseCattyHome/aorb/backend/rpc/user"
	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/config"
	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/strings"
	"github.com/BigNoseCattyHome/aorb/backend/utils/logging"
	"github.com/BigNoseCattyHome/aorb/backend/utils/storage/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 使用logging库，添加字段日志 UserRpcServerName
var log = logging.LogService(config.UserRpcServerName)

func GetUserInfo(username string, fields []string) (resp *user.UserResponse, err error) {
	// 根据 username 和 fields 获取用户信息
	collection := database.MongoDbClient.Database("aorb").Collection("users")
	filter := bson.M{"username": username}
	projection := bson.M{} // 设置查询的字段
	for _, field := range fields {
		projection[field] = 1 // 1表示要查询的字段, 0表示不查询
	}
	opts := options.FindOne().SetProjection(projection)
	var queryUser user.User
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&queryUser)
	log.Debug("get the user with fields: ", &queryUser)
	if err != nil {
		// 如果查询不到用户，返回 UnableToQueryUserErrorCode "无法查询到对应用户"
		// 但是没有返回错误，只在返回的用户信息中 StatusCode 和 StatusMsg 字段中返回错误信息
		if err == mongo.ErrNoDocuments {
			return &user.UserResponse{
				StatusCode: strings.UnableToQueryUserErrorCode,
				StatusMsg:  strings.UnableToQueryUserError,
			}, nil
		}
		// 其他错误，直接返回错误
		return nil, err
	}

	// 返回用户信息
	resp = &user.UserResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		User:       &queryUser,
	}

	return resp, nil
}

// 根据 username 查询用户是否存在
func CheckUserExists(username string) (*user.UserExistResponse, error) {
	collection := database.MongoDbClient.Database("aorb").Collection("users")
	filter := bson.M{"username": username}
	var queryUser user.User
	err := collection.FindOne(context.TODO(), filter).Decode(&queryUser)
	if err != nil {
		// 当用户不存在的时候，返回 false ，没有错误
		if err == mongo.ErrNoDocuments {
			return &user.UserExistResponse{
				StatusCode: strings.ServiceOKCode,
				StatusMsg:  strings.ServiceOK,
				Existed:    false,
			}, nil
		}
		// 当在查询中出现其他的错误时，返回错误
		return nil, err
	}

	return &user.UserExistResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		Existed:    true,
	}, nil
}

// 查询一个用户是否关注另外一个用户
func IsUserFollowing(username string, target_username string) (*user.IsUserFollowingResponse, error) {
	collection := database.MongoDbClient.Database("aorb").Collection("users")
	filter := bson.M{"username": username}
	var queryUser user.User
	err := collection.FindOne(context.TODO(), filter).Decode(&queryUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &user.IsUserFollowingResponse{
				StatusCode:  strings.ServiceOKCode,
				StatusMsg:   strings.ServiceOK,
				IsFollowing: false,
			}, nil
		}
		return nil, err
	}

	// 检查 followed.userids 中是否包含 target_username
	isFollowing := false
	for _, username := range queryUser.Followed.Usernames {
		if username == target_username {
			isFollowing = true
			break
		}
	}

	return &user.IsUserFollowingResponse{
		StatusCode:  strings.ServiceOKCode,
		StatusMsg:   strings.ServiceOK,
		IsFollowing: isFollowing,
	}, nil
}
