package models

type movementDomain struct {
	id          int
	name        string
	value       int64
	description string
	operationID int
	typeID      int
}

func (md *movementDomain) GetID() int {
	return md.id
}
func (md *movementDomain) SetID(id int) {
	md.id = id
}
func (md *movementDomain) GetName() string {
	return md.name
}
func (md *movementDomain) SetName(name string) {
	md.name = name
}
func (md *movementDomain) GetValue() int64 {
	return md.value
}
func (md *movementDomain) SetValue(value int64) {
	md.value = value
}
func (md *movementDomain) GetDesc() string {
	return md.description
}
func (md *movementDomain) SetDesc(desc string) {
	md.description = desc
}
func (md *movementDomain) GetOperation() int {
	return md.operationID
}
func (md *movementDomain) SetOperationID(operationID int) {
	md.operationID = operationID
}
func (md *movementDomain) GetTypeID() int {
	return md.typeID
}
func (md *movementDomain) SetTypeID(typeID int) {
	md.typeID = typeID
}
