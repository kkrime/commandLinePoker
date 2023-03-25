// package errors is used to store custom error types
package errors

import "fmt"

// UserInputError is not a fatal error, but an error in the user input
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
