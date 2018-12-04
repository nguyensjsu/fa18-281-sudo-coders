package main

import (
	// "fmt"
	"net/http"
	// "encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
)

// MongoDB Config
var mongodb_server = "localhost:27017"
var mongodb_database = "superclipper"
var mongodb_collection = "payment"

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		formatter.JSON(writer, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

func getPaymentByCardId(formatter *render.Render) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve the cardId sent as parameter 
        params := mux.Vars(request)
		var cardId string = params["cardId"]

		//Start MongoDB session
        session, error := mgo.Dial(mongodb_server)
        if error != nil {
			formatter.JSON(writer, http.StatusServiceUnavailable, "")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)
		
		//Find the Document present in MongoDB collection with matching cardId
        var cardPayment = CardPayment{}
        error = collection.Find(bson.M{"cardid" : cardId}).One(&cardPayment)
        if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return		
		}
        formatter.JSON(writer, http.StatusOK, cardPayment)
    }
}