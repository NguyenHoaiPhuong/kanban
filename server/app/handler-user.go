package app

import (
	"context"
	"log"
	"net/http"
	"reflect"

	"github.com/NguyenHoaiPhuong/warehouse/server/auth"

	"github.com/NguyenHoaiPhuong/warehouse/server/jsonfunc"
	"github.com/NguyenHoaiPhuong/warehouse/server/models"
	"github.com/NguyenHoaiPhuong/warehouse/server/utils"
)

func (a *App) authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("Start authenticating user")

	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user models.User
	err := jsonfunc.ConvertFromJSON(r.Body, &user)
	defer r.Body.Close()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		log.Printf("JSON conversion error: %v\n", err.Error())
		return
	}

	colName := "Users"
	fieldName := "UserName"
	mod, err := a.mdb.GetDocumentByKey(context.Background(), colName, reflect.TypeOf(user), fieldName, user.UserName)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		log.Printf("Getting user %s in database error: %v\n", user.UserName, err.Error())
		return
	}
	userInDB := mod.(*models.User)
	if userInDB.Password == user.Password {
		token, err := auth.GenerateTokenPair(userInDB, "secret")
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			log.Printf("Generating token pair error: %v\n", err.Error())
			return
		}
		utils.RespondJSON(w, http.StatusOK, token)
	}
}
