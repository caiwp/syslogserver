package syslogserver

type Handler interface {
	Handle(data map[string]interface{})
}
