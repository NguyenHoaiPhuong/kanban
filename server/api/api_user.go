package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Users users api
type Users struct {
	API
}

func (api *Users) init(root *mux.Router, path string) {
	log.Println("Init Users API")
	api.Router = root.PathPrefix(path).Subrouter()
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Start getting all users from database")

	// Set up header
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// docs, err := a.Database.GetAllDocuments(*a.Config.MongoDBConfig.DBName, string(model.BookCol))
	// if err != nil {
	// 	utils.RespondError(w, http.StatusInternalServerError, error.ErrorDB)
	// 	var errNew error.Imp
	// 	errNew.SetErrorMessage(err.Error())
	// 	errNew.InsertErrorMessage(error.ErrorAppGetAllBooks)
	// 	log.Printf("%v\n", errNew.Error())
	// 	return
	// }

	// books := make([]*model.Book, len(docs))
	// for i, doc := range docs {
	// 	book := new(model.Book)
	// 	bsonBytes, _ := bson.Marshal(doc)
	// 	bson.Unmarshal(bsonBytes, book)
	// 	books[i] = book
	// }

	// err = utils.RespondJSON(w, http.StatusOK, books)
	// if err != nil {
	// 	utils.RespondError(w, http.StatusInternalServerError, error.ErrorJSON)
	// 	var errNew error.Imp
	// 	errNew.SetErrorMessage(err.Error())
	// 	errNew.InsertErrorMessage(error.ErrorAppGetAllBooks)
	// 	log.Printf("%v\n", errNew.Error())
	// 	return
	// }

	log.Println("Finish getting all users from database")
}
