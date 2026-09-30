package tools

type Tool interface {
	Name() string
	Execute(input string) (string, error)
} // how is interface used , what is the use of it?
