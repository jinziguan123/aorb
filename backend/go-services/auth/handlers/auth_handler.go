package handlers

import (
	"context"

	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/strings"

	"github.com/BigNoseCattyHome/aorb/backend/go-services/auth/services"
	"github.com/BigNoseCattyHome/aorb/backend/rpc/auth"
	"github.com/BigNoseCattyHome/aorb/backend/rpc/user"
	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/config"
	"github.com/BigNoseCattyHome/aorb/backend/utils/logging"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

var log = logging.LogService(config.AuthRpcServerName) // 使用logging库，添加字段日志AuthRpcServer
var conn *amqp.Connection                              // amqp.Connection用于连接RabbitMQ服务器
var channel *amqp.Channel                              // amqp.Channel用于与RabbitMQ服务器通信

// exitOnError 如果err不为nil，则panic
func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// CloseMQConn 关闭RabbitMQ连接
func CloseMQConn() {
	if err := conn.Close(); err != nil {
		panic(err)
	}
	if err := channel.Close(); err != nil {
		panic(err)
	}
}

// AuthServiceImpl AuthService服务实现
type AuthServiceImpl struct {
	auth.AuthServiceServer
}

// 初始化
func (a AuthServiceImpl) Init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	exitOnError(err)
	channel, err = conn.Channel()
	exitOnError(err)
}

// 创建AuthService服务实例
func (a AuthServiceImpl) New() {

}

// 登录
func (a AuthServiceImpl) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {

	// 解析请求
	login_request := auth.LoginRequest{
		Username:  request.Username,
		Password:  request.Password,
		DeviceId:  request.DeviceId,
		Nonce:     request.Nonce,
		Timestamp: request.Timestamp,
	}

	// 调用服务
	token, exp_token, refresh_token, simple_user, err := services.AuthenticateUser(&login_request)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "login failed: %v", err)
	}

	// 返回响应
	loginResponse := &auth.LoginResponse{
		StatusCode:   strings.ServiceOKCode,
		StatusMsg:    strings.ServiceOK,
		Token:        token,
		TokenType:    "Bearer",
		ExpiresAt:    exp_token,
		RefreshToken: refresh_token,
		SimpleUser: &auth.SimpleUser{
			Avatar:    simple_user.Avatar,
			Username:  simple_user.Username,
			Ipaddress: simple_user.Ipaddress,
			Nickname:  simple_user.Nickname,
		},
	}
	return loginResponse, nil
}

// Verify 验证
func (a AuthServiceImpl) Verify(context context.Context, request *auth.VerifyRequest) (*auth.VerifyResponse, error) {

	// 调用服务
	claims, err := services.VerifyAccessToken(request.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token verification failed: %v", err)
	}

	// 返回响应
	verifyResponse := &auth.VerifyResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		UserId:     claims.UserId,
		Username:   claims.Username,
		ExpiresAt:  claims.ExpiresAt,
	}
	return verifyResponse, nil
}

// Refresh 刷新
func (a AuthServiceImpl) Refresh(context context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {

	// 调用服务
	newToken, exp_token, err := services.RefreshAccessToken(request.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token failed: %v", err)
	}

	// 返回响应
	refreshResponse := &auth.RefreshResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		Token:      newToken,
		ExpiresAt:  exp_token,
	}
	return refreshResponse, nil
}

// Logout 登出
func (a AuthServiceImpl) Logout(context context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	// 解析参数
	accessToken := request.AccessToken
	refreshToken := request.RefreshToken

	// 调用服务
	// 验证访问令牌，确保合法用户的操作
	claim, err := services.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid access token: %v", err)
	}
	// 撤销刷新令牌
	err = services.RevokeRefreshToken(claim.UserId, refreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "revoke refresh token failed: %v", err)
	}

	// 返回响应
	logoutResponse := &auth.LogoutResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
	}
	return logoutResponse, nil

}

// Register 注册
func (a AuthServiceImpl) Register(context context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Infof("Received Register request: %v", request)

	// 解析参数
	user := user.User{
		Username:  request.Username,
		Password:  &request.Password,
		Nickname:  request.Nickname,
		Avatar:    request.Avatar,
		Ipaddress: &request.Ipaddress,
	}

	// 调用服务
	err := services.RegisterUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "register failed: %v", err)
	}

	// 返回响应
	registerResponse := &auth.RegisterResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
	}
	return registerResponse, nil
}
