package database

type Handler interface {
	FindByID (result interface{}, id int) (error)
	Close    ()
}
