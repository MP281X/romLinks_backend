package logger

import (
	"errors"
)

// db error
var ErrDbRead = errors.New("unable to get the data")
var ErrDbWrite = errors.New("unable to save the data")
var ErrDbEdit = errors.New("unable to edit the data")
var ErrDeviceAlreadyExist = errors.New("there is already a device with the same info")
var ErrUserAlreadyExist = errors.New("there is already a user with the same username or email")
var ErrReqAlreadyExist = errors.New("there is already a rom request with the same data")
var ErrUnauthorized = errors.New("unauthorized")

// encryption error
var ErrTokenGen = errors.New("unable to generate the user token")
var ErrTokenRead = errors.New("unable to validate the current user")
var ErrHashGen = errors.New("unable to encrypt the password")
var ErrHashCompare = errors.New("invalid password")

// validation error
var ErrInvUsername = errors.New("invalid username")
var ErrInvEmail = errors.New("invalid email")
var ErrInvPassword = errors.New("invalid password")

func ResCode(err error) int {
	switch err {
	case ErrDbRead:
		return 404
	case ErrDbWrite:
		return 500
	case ErrDbEdit:
		return 500
	case ErrDeviceAlreadyExist:
		return 409
	case ErrUserAlreadyExist:
		return 409
	case ErrUnauthorized:
		return 401
	case ErrTokenGen:
		return 500
	case ErrTokenRead:
		return 401
	case ErrHashGen:
		return 500
	case ErrHashCompare:
		return 400
	case ErrInvUsername:
		return 400
	case ErrInvEmail:
		return 400
	case ErrInvPassword:
		return 400
	default:
		return 400
	}
}
