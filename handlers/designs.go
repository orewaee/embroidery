package handlers

import (
	"context"
	"encoding/json"
	"github.com/orewaee/embroidery-api/database"
	"github.com/orewaee/embroidery-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func Designs(writer http.ResponseWriter, request *http.Request) {
	designs := database.GetCollection("designs")

	cur, err := designs.Find(context.TODO(), bson.D{{"visible", true}})
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	var list []models.Design
	if err := cur.All(context.TODO(), &list); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(list)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
