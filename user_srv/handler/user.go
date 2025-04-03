package handler

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	"strings"
	"time"
)

type UserServer struct {
	userProto.UnimplementedUserServer
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ModelToResponse(user *model.User) *userProto.UserInfoResponse {
	userInfoRsp := &userProto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

// 获取用户列表
func (u UserServer) GetUserList(ctx context.Context, req *userProto.PageInfo) (*userProto.UserListResponse, error) {
	rsp := &userProto.UserListResponse{}

	var count int64
	global.DB.Model(&model.User{}).Count(&count)
	rsp.Total = int32(count)

	var users []model.User
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userInfoRsp := ModelToResponse(&user)
		rsp.Data = append(rsp.Data, userInfoRsp)
	}

	return rsp, nil
}

// 通过手机查询用户
func (u UserServer) GetUserByMobile(ctx context.Context, req *userProto.MobileRequest) (*userProto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	userInfoRsp := ModelToResponse(&user)
	return userInfoRsp, nil
}

// 通过ID查询用户
func (u UserServer) GetUserById(ctx context.Context, req *userProto.IdRequest) (*userProto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	userInfoRsp := ModelToResponse(&user)
	return userInfoRsp, nil
}

// 创建用户
func (u UserServer) CreateUser(ctx context.Context, req *userProto.CreateUserInfo) (*userProto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = req.Mobile
	user.NickName = req.Nickname
	options := &password.Options{16, 100, 32, sha256.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha256$%s$%s", salt, encodedPwd)
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoRsp := ModelToResponse(&user)
	return userInfoRsp, nil
}

// 更新用户
func (u UserServer) UpdateUser(ctx context.Context, req *userProto.UpdateUserInfo) (*emptypb.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	user.NickName = req.Nickname
	user.Gender = req.Gender
	birthday := time.Unix(int64(req.Birthday), 0)
	user.Birthday = &birthday
	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

// 校验密码
func (u UserServer) CheckPassword(ctx context.Context, req *userProto.CheckInfo) (*userProto.CheckResponse, error) {
	options := &password.Options{16, 100, 32, sha256.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &userProto.CheckResponse{Success: check}, nil
}
