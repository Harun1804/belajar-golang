package iomanager

type IOManager interface {
	ReadFile() ([]string, error)
	WriteJson(data interface{}) error
}