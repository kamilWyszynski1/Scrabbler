package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scrabble"
	"scrabble/game"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	engine, err := game.NewGameEngine(logrus.New())
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/put", func(writer http.ResponseWriter, request *http.Request) {
		var req scrabble.PutRequest

		err := json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		points, err := engine.Put(req)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = fmt.Fprint(writer, points)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", r)
}
