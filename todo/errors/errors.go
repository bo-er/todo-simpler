package errors

import "errors"

var (
	// ErrDbIsNil 没有数据库
	ErrDbIsNil = errors.New("db is nil") 
)
