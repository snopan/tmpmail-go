package tmpmailgo

import "errors"

const (
	host1SecMail = "https://www.1secmail.com/api/v1"

	emailUsernameLength = 10
)

var (
	errorInvalidEmail  = errors.New("not a valid email")
	errorInvalidDomain = errors.New("not a valid domain")
)
