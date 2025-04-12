package response

type OperationResponse struct {
	ID      int64       `json:"id"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}

func (s *OperationResponse) ResultOperation(p OperationResponse) Response {

	return &OperationResponse{
		ID:      p.ID,
		Success: p.Success,
		Message: p.Message,
		Data:    p.Data,
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
