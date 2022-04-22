package common

import "github.com/pkg/errors"

var (
	CanceledError = errors.New("ctx has canceled")
)
