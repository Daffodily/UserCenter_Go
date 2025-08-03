package dto

type RegisterDTO struct {
	Username       string `json:"userName"`
	Password       string `json:"password" binding:"required,min=6,max=10,alphanum"`
	MobilePhoneNum string `json:"mobilePhoneNum" binding:"required,len=11"`
}
