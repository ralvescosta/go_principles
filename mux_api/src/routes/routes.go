package routes

import (
	"encoding/json"
	"goapi/src/entities"
	"goapi/src/repositories"
	"math/rand"
	"net/http"
)

var (
	repository repositories.PostRepository = repositories.PostRepositoryConstructor()
)

func GetPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	posts, err := repository.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error geting posts"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var body entities.Post

	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshaling array"}`))
		return
	}

	body.Id = rand.Int63()
	repository.Save(&body)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(body)
}
