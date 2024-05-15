package request

type MovementRequest struct {
	Value       float64 `json:"value" binding:"required"`
	Description string  `json:"description" binding:"required"`
	OperationID int     `json:"operation_id" binding:"required"`
	TypeID      int     `json:"type_id" binding:"required"`
	UserID      int     `json:"user_id" binding:"required"`
}
