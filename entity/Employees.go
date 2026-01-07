package entity

import (
	"gorm.io/gorm"
)

type Employees struct{
	gorm.Model
	Name         string   `valid:"length(2|80)"`
	Salary       float64  `valid:"range(15000|20000)"`
	EmployeeCode string   `valid:"alphanum"`
}