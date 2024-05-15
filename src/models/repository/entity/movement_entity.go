package entity

import "gorm.io/gorm"

type Movement struct {
	gorm.Model
	Value       float64
	Description string
	UserID      int
	TypeID      int
	OperationID int
}
