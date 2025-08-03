package dto

const (
	CodeSuccess          = 0
	ErrCodePwdHashFailed = 10000

	ErrCodeRegisterFailed    = 20000
	ErrCodePhoneRequired     = 20001
	ErrCodePasswordInvalid   = 20002
	ErrCodeUserAlreadyExists = 20003

	ErrCodeLoginFailed         = 30000
	ErrCodeUserNotExists       = 30001
	ErrCodePwdIsNotRight       = 30002
	ErrCodeTokenGenerateFailed = 30003
)

var FieldErrCodeMap = map[string]int{
	"MobilePhoneNum": ErrCodePhoneRequired,
	"Password":       ErrCodePasswordInvalid,
}
