package models

type MovementDomainInterface interface {
	GetID() int
	SetID(int)
	GetValue() float64
	SetValue(float642 float64)
	GetDesc() string
	SetDesc(string)
	GetOperation() int
	SetOperationID(int)
	GetTypeID() int
	SetTypeID(int)
	GetUserID() int
	SetUserID(int)
}

func NewMovementsDomain(description string, value float64, operationID, typeID, userID int) MovementDomainInterface {
	return &movementDomain{
		description: description,
		value:       value,
		operationID: operationID,
		typeID:      typeID,
		userID:      userID,
	}
}
