package errors

import (
	"fmt"
)

type ErrorCreatingRecordInDB struct{}

func (e ErrorCreatingRecordInDB) Error() string {
	return fmt.Sprintf("Error creating creating a new entry in DB")
}

type ErrorUpdatingRecordInDB struct{}

func (e ErrorUpdatingRecordInDB) Error() string {
	return fmt.Sprintf("Error updating entry in DB")
}

type ErrorFindingOneRecordInDB struct {
	Model string
	ID    int
}

func (e ErrorFindingOneRecordInDB) Error() string {
	return fmt.Sprintf("failed to find %s with id '%d'", e.Model, e.ID)
}
