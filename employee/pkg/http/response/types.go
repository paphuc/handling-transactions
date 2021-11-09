package response

type (
	appError interface {
		Error() string
		Code() uint32
		Message() string
	}
	Base struct {
		Result interface{} `json:"result,omitempty"`
		Page   interface{} `json:"page,omitempty"`
		ID     interface{} `json:"id,omitempty"`
	}
)
