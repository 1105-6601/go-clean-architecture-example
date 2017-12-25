package context

type HTTP interface {
	Param   (string)           (string)
	Query   (string)           (string)
	Bind    (interface{})      (error)
	Status  (int)
	JSON    (int, interface{})
	JSONRaw (int, string)
}
