package errors

import (
	"fmt"
)

type CustomError struct {
	Arg     int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Arg, e.Message)

}

// User errors

// Register errors
var (
	ErrUserAlreadyExistsWithEmail = &CustomError{Arg: 409, Message: "User already exists with this email"}
	ErrUserAlreadyExistsWithPhone = &CustomError{Arg: 409, Message: "User already exists with this phone number"}
)

var (
	ErrUserNotFoundWithPhone = &CustomError{Arg: 409, Message: "User was not found with phone number"}
)

var (
	ErrTokenIsNotValid = &CustomError{Arg: 409, Message: "Token is not valid"}
)

// Warehouse errors

var (
	ErrWarehouseAlreadyExist = &CustomError{Arg: 409, Message: "Warehouse Already Exist with name"}
	ErrWareHouseNotFound     = &CustomError{Arg: 409, Message: "Warehouse not found with name"}
)

// Zone errors

var ()
