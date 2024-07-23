package result

import "errors"

var (
	ErrUnauthorized = errors.New("未授权")
	ErrForbidden    = errors.New("权限不足")
)
