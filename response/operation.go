package response

import "net/http"

type OperationResponse struct {
	ID      int64       `json:"id"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func Success(id int64, message string, data interface{}) Response {
	return Result(id, true, message, http.StatusOK, data)
}

func Error(message string, status int, data interface{}) Response {
	return Result(0, false, message, status, data)
}

func Result(id int64, success bool, message string, status int, data interface{}) Response {

	return &OperationResponse{
		ID:      id,
		Success: success,
		Message: message,
		Data:    data,
		Status:  status,
	}
}

func (s OperationResponse) GetID() int64 {
	return s.ID
}

func (s OperationResponse) GetData() interface{} {
	return s.Data
}

func (s OperationResponse) GetSuccess() bool {
	return s.Success
}

func (s OperationResponse) GetMessage() string {
	return s.Message
}

func (e OperationResponse) StatusCode() int {
	return e.Status
}
