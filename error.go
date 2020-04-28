package database

import "fmt"

type DbError struct {
	Code        int
	Description string
}

func (err DbError) Error() string {
	return fmt.Sprint("Database error ", err.Code, ": ", err.Description)
}

func NewDbError(code int, desc string) DbError {
	return DbError{
		Code:        code,
		Description: desc,
	}
}