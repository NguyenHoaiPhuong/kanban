package api

import (
	"github.com/gorilla/mux"
)

// APIs include all apis
type APIs struct {
	Root  *mux.Router
	users Users
}

// Init : initialize all apis
func (apis *APIs) Init() {
	apis.Root = mux.NewRouter()

	apis.users.init(apis.Root, "/users")
	apis.users.RegisterHandleFunction("GET", "", getAllUsers)
}
