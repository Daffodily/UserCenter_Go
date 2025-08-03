package service

import (
	"errors"
	"gorm.io/gorm"
	"usercenter/config"
	"usercenter/dto"
	"usercenter/model"
)

// RegisterResult 服务层注册结果
type RegisterResult struct {
	Code    int         // 业务错误码
	Message string      // 错误描述信息
	User    *model.User // 成功时返回用户信息
}

func Register(userinfo dto.RegisterDTO) RegisterResult {
	//判断用户是否已存在
	var user model.User
	result := config.DB.Where("username = ?", userinfo.Username).First(&user)
	if result.Error == nil {
		return RegisterResult{
			Code: dto.ErrCodeUserAlreadyExists,
		}
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return RegisterResult{
			Code: dto.ErrCodeRegisterFailed,
		}
	}

	//创建用户
	newUser := model.User{
		Username:    userinfo.Username,
		Password:    userinfo.Password,
		PhoneNumber: userinfo.MobilePhoneNum,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		return RegisterResult{
			Code: dto.ErrCodeRegisterFailed,
		}
	} else {
		//注册成功
		return RegisterResult{
			Code: 0,
			User: &newUser,
		}
	}
}
