package api

import (
	"github.com/gorilla/mux"
)

// APIs include all apis
type APIs struct {
	users Users
}

// Init : initialize all apis
func (apis *APIs) Init() {
	root := mux.NewRouter()

	apis.users.init(root, "/users")
	apis.users.RegisterHandleFunction("GET", "", getAllUsers)
}
