package models

import (
	uuid "github.com/satori/go.uuid"
)

// IModel : interface of Model
type IModel interface {
	New()
	GetID() string
}

// IModels : slice of IModel
type IModels []IModel

// Model implements the interface
type Model struct {
	ID string `json:"ID" bson:"ID"`
}

// New : create new model
func (mod *Model) New() {
	mod.ID = uuid.NewV4().String()
}

// GetID : return ID of the model
func (mod *Model) GetID() string {
	return mod.ID
}
