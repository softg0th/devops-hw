package exceptions

import "errors"

func HandleRunTimeError(errorType string) error {
	return errors.New(errorType)
}
