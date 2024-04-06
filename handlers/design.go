package handlers

import (
	"context"
	"github.com/orewaee/embroidery-api/database"
	"github.com/orewaee/embroidery-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strconv"
)

type DesignHandler struct{}

func NewDesign() *DesignHandler {
	return &DesignHandler{}
}

func (*DesignHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	id := request.PathValue("id")

	designs := database.GetCollection("designs")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var design models.Design
	err = designs.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&design)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if !design.Visible {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	path := "./" + design.Path
	bytes, err := os.ReadFile(path)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	length := len(bytes)

	if _, err := writer.Write(bytes); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.Header().Add("Content-Type", "image/"+design.Extension)
		writer.Header().Add("Content-Length", strconv.Itoa(length))
	}
}
