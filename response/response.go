package response

type Response interface {
	GetID() int64
	GetData() interface{}
	GetSuccess() bool
	GetMessage() string
	StatusCode() int
}
