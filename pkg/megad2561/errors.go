package megad2561

import (
	"errors"
)

var (
	ErrUpdateConfig = errors.New("cant update config")
	ErrMaxLengthID  = errors.New("device id maximum 5 length")
)
