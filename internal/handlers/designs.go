package handlers

import (
	"context"
	"encoding/json"
	"github.com/orewaee/embroidery-api/internal/database"
	"github.com/orewaee/embroidery-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type DesignsHandler struct{}

func NewDesigns() *DesignsHandler {
	return &DesignsHandler{}
}

func (*DesignsHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	designs := database.GetCollection("designs")

	cur, err := designs.Find(context.TODO(), bson.D{{"visible", true}})
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	var list []models.Design
	if err := cur.All(context.TODO(), &list); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(list)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
