package repository

import (
	"practice_gormsqlmock/database"
	"practice_gormsqlmock/domain/entity"

	"gorm.io/gorm/clause"
)

type DepartmentRepository interface {
	FindAll() []entity.Department
}

type Department struct{}

func (Department) First() entity.Department {
	var department entity.Department
	database.Db.Debug().Preload(clause.Associations).First(&department)
	return department
}
