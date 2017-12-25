package context

type Command interface {
	Echo (string)
	Args ()       ([]interface{})
}
