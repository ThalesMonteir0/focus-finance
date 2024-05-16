package entity

import "gorm.io/gorm"

type Movement struct {
	gorm.Model
	Value       float64
	Description string
	UserID      int
	User        User
	TypeID      int
	Type        Type
	OperationID int
	Operation   Operation
}
