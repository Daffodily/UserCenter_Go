package dto

const (
	ErrCodeRegisterFailed    = 20000
	ErrCodePhoneRequired     = 20001
	ErrCodePasswordInvalid   = 20002
	ErrCodeUserAlreadyExists = 20003
)

var FieldErrCodeMap = map[string]int{
	"MobilePhoneNum": ErrCodePhoneRequired,
	"Password":       ErrCodePasswordInvalid,
}
