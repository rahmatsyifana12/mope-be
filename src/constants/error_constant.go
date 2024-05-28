package constants

import "errors"

var (
	ERR_NOT_LOGGED_IN = errors.New("not logged in and no permission")
)