package utils

type Validator struct {
	FieldErrors map[string]string
}

type Errors struct {
	Errors map[string]string `json:"errors"`
}
