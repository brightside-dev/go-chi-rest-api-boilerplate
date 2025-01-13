package validators

type Validator interface {
	ValidateRequest(req interface{}) error
	FormatErrorMessage(err error) string
}
