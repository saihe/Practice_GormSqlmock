package entity

import "github.com/gofrs/uuid"

type Employee struct {
	ID           uuid.UUID
	Name         string
	DepartmentId uuid.UUID
}
