package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dimiantoni/api-clean-archi-go/usecase/user"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dimiantoni/api-clean-archi-go/api/presenter"

	"github.com/dimiantoni/api-clean-archi-go/domain/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listUsers(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*entity.User
		var err error
		email := r.URL.Query().Get("email")
		switch {
		case email == "":
			data, err = service.ListUsers()
		default:
			data, err = service.SearchUsers(email)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.User
		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:      d.ID,
				Name:    d.Name,
				Email:   d.Email,
				Address: d.Address,
				Age:     d.Age,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var input struct {
			Name     string `json:"name" bson:"name",omitempty`
			Email    string `json:"email" bson:"email",omitempty`
			Password string `json:"password" bson:"password",omitempty`
			Address  string `json:"address" bson:"address",omitempty`
			Age      int8   `json:"age" bson:"age",omitempty`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateUser(input.Name, input.Email, input.Password, input.Address, input.Age)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.User{
			ID:       id,
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
			Address:  input.Address,
			Age:      input.Age,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// func updateUser(service user.UseCase) http.Handler {
// 	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	// 	errorMessage := "Error updating user"
// 	// 	vars := mux.Vars(r)
// 	// 	id, err := primitive.ObjectIDFromHex(vars["id"])
// 	// 	if err != nil {
// 	// 		w.WriteHeader(http.StatusInternalServerError)
// 	// 		w.Write([]byte(errorMessage))
// 	// 		return
// 	// 	}
// 	// 	var input struct {
// 	// 		ID       string `json:"id" bson:"_id",omitempty`
// 	// 		Name     string `json:"name" bson:"name",omitempty`
// 	// 		Email    string `json:"email" bson:"email",omitempty`
// 	// 		Password string `json:"password" bson:"password",omitempty`
// 	// 		Address  string `json:"address" bson:"address",omitempty`
// 	// 		Age      int8   `json:"age" bson:"age",omitempty`
// 	// 	}
// 	// 	id, err := service.UpdateUser(id, input.Name, input.Email, input.Password, input.Address, input.Age)
// 	// 	err = json.NewDecoder(r.Body).Decode(&input)
// 	// 	if err != nil {
// 	// 		log.Println(err.Error())
// 	// 		w.WriteHeader(http.StatusInternalServerError)
// 	// 	}
// 	// })
// }

func getUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.User{
			ID:      data.ID,
			Name:    data.Name,
			Email:   data.Email,
			Address: data.Address,
			Age:     data.Age,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteUser(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(listUsers(service)),
	)).Methods("GET", "OPTIONS").Name("listUsers")

	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(getUser(service)),
	)).Methods("GET", "OPTIONS").Name("getUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(deleteUser(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}
