package handlers

import (
	"context"
	"fmt"
	"github.com/orewaee/embroidery-api/internal/database"
	"github.com/orewaee/embroidery-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strconv"
)

type DesignHandler struct{}

func NewDesignHandler() *DesignHandler {
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

	path := fmt.Sprintf("./designs/%s.jpg", id)
	bytes, err := os.ReadFile(path)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if _, err := writer.Write(bytes); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.Header().Add("Content-Type", "image/jpeg")
		writer.Header().Add("Content-Length", strconv.Itoa(len(bytes)))
	}
}
