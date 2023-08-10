package error

import (
	"fmt"
)

const (
	CodeNormal               = 0  // normal return code
	CodeInternalError        = -1 // omni internal server error
	CodeInvalidParam         = 1  // invalid parameter in request
	CodeModelNotExist        = 2  // types(e.g. checkpoint) not exist
	CodeTaskIdNotExist       = 3  // task_id not exist
	CodeInvalidAuth          = 4  // key is invalid
	CodeParamRangeOutOfLimit = 6  // parameter(e.g. height\size) out of range
	CodeCostBalanceFailure   = 7  // balance is not enought
	CodeSamplerNotExist      = 8  // sampler not found
	CodeNotSupport           = 10 // feature not supported
	CodePromptIllegal        = 11 // prompt is illegal(e.g. child porn)
)

type OmniError struct {
	Code int
	Msg  string
}

func NewOmniError(code int, msg string) *OmniError {
	return &OmniError{
		Code: code,
		Msg:  msg,
	}
}

func (o OmniError) Error() string {
	return fmt.Sprintf("omni's api returns error, code = %d, msg = %s, please refer to error/omni.go for details", o.Code, o.Msg)
}
