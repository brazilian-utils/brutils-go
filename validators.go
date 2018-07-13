package brutils

type Validator interface {
	Validate(input string) bool
}
