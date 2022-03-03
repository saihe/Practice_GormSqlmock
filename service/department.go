package service

import (
	"log"
	"practice_gormsqlmock/domain/entity"
	"practice_gormsqlmock/domain/repository"
)

type DepartmentService interface {
	GetDepartment() entity.Department
}

type Department struct{}

func (Department) GetDepartment() entity.Department {
	repo := &repository.Department{}
	department := repo.First()
	if len(department.Employee) == 0 {
		log.Print("取得失敗")
		return entity.Department{}
	}
	return department
}
