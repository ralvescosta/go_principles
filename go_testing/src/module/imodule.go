package module

// SomethingStruct ...
type SomethingStruct struct {
	id   uint16
	name string
}

// IModule ...
type IModule interface {
	DoSomething(id uint16) *SomethingStruct
}
