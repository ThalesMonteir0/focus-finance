package response

type MovementResponse struct {
	ID            int     `json:"id"`
	Value         float64 `json:"value"`
	Description   string  `json:"description"`
	OperationID   int     `json:"operation_id"`
	TypeID        int     `json:"type_id"`
	UserID        int     `json:"user_id"`
	TypeName      string  `json:"type_name"`
	OperationName string  `json:"operation_name"`
}
