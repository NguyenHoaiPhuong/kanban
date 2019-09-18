package api

import (
	"github.com/gorilla/mux"
)

type APIs struct {
	apiUsers APIUsers
}

func (apis *APIs) Init() {
	root := mux.NewRouter()

	apis.apiUsers.init(root, "/users")
	apis.apiUsers.RegisterHandleFunction("GET", "", getAllUsers)
}
