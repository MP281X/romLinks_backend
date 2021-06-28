package logger

import "errors"

// init error
var ErrDbInit = errors.New("db init")
var ErrConfig = errors.New("config read")
var ErrLog = errors.New("log init")
var ErrApi = errors.New("api init")

// db error
var ErrDbRead = errors.New("db read")
var ErrDbWrite = errors.New("db write")
var ErrDuplicateKey = errors.New("key already exist")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidKey = errors.New("invalid db key")

// encryption error
var ErrTokenGen = errors.New("token gen")
var ErrTokenRead = errors.New("token read")
var ErrHashGen = errors.New("password hash")
var ErrHashCompare = errors.New("password comparation")

// validation error
var ErrInvUsername = errors.New("invalid username")
var ErrInvEmail = errors.New("invalid email")
var ErrInvPassword = errors.New("invalid password")
var ErrInv = errors.New("invalid ")
