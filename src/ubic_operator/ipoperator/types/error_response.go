package types

import (
	"errors"
	"fmt"
)

//Define error
var (
	ErrIPOperator = errors.New("ip operator error")

	ErrNoSuchReservation = ipOperatorError{
		message: "no such reservation with that lease ID",
		code:    1000,
	}

	ErrReservationQuantityCannotBeZero = ipOperatorError{
		message: "reservation request cannot have a quantity of zero",
		code:    1001,
	}

	errNoRegisteredError = errors.New("no registered error")
)

//IPOperatorError is interface
type IPOperatorError interface {
	error
	GetCode() int
}

type ipOperatorError struct {
	message string
	code    int
}

func (ipoe ipOperatorError) Error() string {
	return fmt.Sprintf("%s: %s", ErrIPOperator.Error(), ipoe.message)
}

func (ipoe ipOperatorError) Unwrap() error {
	return ErrIPOperator
}

func (ipoe ipOperatorError) GetCode() int {
	return ipoe.code
}

//IPOperatorErrorResponse is struct
type IPOperatorErrorResponse struct {
	Error string
	Code  int
}

var registry map[int]error

func registerError(err IPOperatorError) {
	existing, exists := registry[err.GetCode()]
	if exists {
		panic(fmt.Sprintf("error already exists with code %d: %v", err.GetCode(), existing))
	}

	registry[err.GetCode()] = err
}

//Init is function
func Init() {
	registry = make(map[int]error)
	registerError(ErrNoSuchReservation)
	registerError(ErrReservationQuantityCannotBeZero)
}
