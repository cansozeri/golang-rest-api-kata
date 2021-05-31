package request

type GetInMemoryRequest struct {
	Key string `validate:"required"`
}
