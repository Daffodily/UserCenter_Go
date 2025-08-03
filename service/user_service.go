package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"usercenter/config"
	"usercenter/dto"
	"usercenter/model"
	"usercenter/utils"
)

// RegisterResult 服务层注册结果
type RegisterResult struct {
	Code    int         // 业务错误码
	Message string      // 错误描述信息
	User    *model.User // 成功时返回用户信息
}

type LoginResult struct {
	Code     int
	Username string
	Token    string
}

func Register(userRegister dto.RegisterDTO) RegisterResult {
	//判断用户是否已存在
	var user model.User
	result := config.DB.Where("username = ?", userRegister.Username).First(&user)
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

	hashedPwd, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		log.Println("密码加密失败:", err)
		return RegisterResult{
			Code: dto.ErrCodePwdHashFailed,
		}
	}
	//创建用户
	newUser := model.User{
		Username:    userRegister.Username,
		Password:    hashedPwd,
		PhoneNumber: userRegister.MobilePhoneNum,
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

func Login(userLogin dto.LoginDTO) LoginResult {
	var user model.User
	result := config.DB.Where("username = ?", userLogin.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return LoginResult{
				Code:     dto.ErrCodeUserNotExists,
				Username: userLogin.Username,
			}
		}
		return LoginResult{
			Code:     dto.ErrCodeLoginFailed,
			Username: userLogin.Username,
		}
	}
	if !utils.CheckPasswordHash(userLogin.Password, user.Password) {
		return LoginResult{
			Code:     dto.ErrCodePwdIsNotRight,
			Username: userLogin.Username,
		}
	}
	token, err := utils.GenerateToken(&user)
	if err != nil {
		return LoginResult{
			Code:     dto.ErrCodeTokenGenerateFailed,
			Username: userLogin.Username,
		}
	}
	return LoginResult{
		Code:     dto.CodeSuccess,
		Username: user.Username,
		Token:    token,
	}

}
