package models

type MovementDomainInterface interface {
	GetID() int
	SetID(int)
	GetName() string
	SetName(string)
	GetValue() int64
	SetValue(int64)
	GetDesc() string
	SetDesc(string)
	GetOperation() int
	SetOperationID(int)
	GetTypeID() int
	SetTypeID(int)
}

func NewMovementsDomain(name, description string, value int64, operationID, typeID int) MovementDomainInterface {
	return &movementDomain{
		name:        name,
		description: description,
		value:       value,
		operationID: operationID,
		typeID:      typeID,
	}
}
