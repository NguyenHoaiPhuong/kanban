package api

import (
	"github.com/gorilla/mux"
)

// APIs include all apis
type APIs struct {
	Root *mux.Router
	Users
	User
}

// Init : initialize all apis
func (apis *APIs) Init() {
	apis.Root = mux.NewRouter()

	apis.Users.init(apis.Root, "/apis/internal/users")
	apis.User.init(apis.Root, "/apis/internal/user")
}
