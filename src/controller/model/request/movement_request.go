package request

type MovementRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Value       int64  `json:"value" binding:"required"`
	Description string `json:"description" binding:"required"`
	OperationID int    `json:"operation_id" binding:"required"`
	TypeID      int    `json:"type_id" binding:"required"`
}
