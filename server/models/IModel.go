package models

// IModel : interface of Model
type IModel interface {
	New()
}

// IModels : slice of IModel
type IModels []IModel

// Model implements the interface
type Model struct {
}

// New : create new model
func (mod *Model) New() {
}
