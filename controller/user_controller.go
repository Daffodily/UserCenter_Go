package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"usercenter/dto"
	"usercenter/service"
)

func Register(c *gin.Context) {
	var user dto.RegisterDTO

	//绑定参数，利用validator判断参数是否合格
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("绑定失败：", err)
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			field := errs[0].Field()
			errorcode := dto.FieldErrCodeMap[field]
			dto.Fail(c, errorcode)
			return
		}
	}

	result := service.Register(user)
	switch result.Code {
	case 0:
		dto.Success(c, gin.H{
			"userid":   result.User.ID,
			"username": result.User.Username,
		})
	case dto.ErrCodeUserAlreadyExists:
		dto.Fail(c, dto.ErrCodeUserAlreadyExists)
	case dto.ErrCodeRegisterFailed:
		dto.Fail(c, dto.ErrCodeRegisterFailed)
	}
}

func Login(c *gin.Context) {

}
