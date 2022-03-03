package entity

import "github.com/gofrs/uuid"

type Department struct {
	ID       uuid.UUID
	Name     string
	Employee []Employee `gorm:"foreignKey:DepartmentId"`
}
