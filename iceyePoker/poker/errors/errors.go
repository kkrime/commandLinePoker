package errors

import "fmt"

type UserInputError struct {
	error string
}

func NewUserInputError(errorMessage string, vars ...interface{}) error {
	return &UserInputError{
		error: fmt.Sprintf(errorMessage, vars...),
	}
}

func (uie *UserInputError) Error() string {
	return uie.error
}
