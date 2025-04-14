package response

type OperationResponse struct {
	ID      int64       `json:"id"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}

func ResultOperation(id int64, success bool, message string, data interface{}) Response {
	return Result(id, success, message, data)
}

func Result(id int64, success bool, message string, data interface{}) Response {

	return &OperationResponse{
		ID:      id,
		Success: success,
		Message: message,
		Data:    data,
	}
}

func (s *OperationResponse) GetID() int64 {
	return s.ID
}

func (s *OperationResponse) GetData() interface{} {
	return s.Data
}

func (s *OperationResponse) GetSuccess() bool {
	return s.Success
}

func (s *OperationResponse) GetMessage() string {
	return s.Message
}
