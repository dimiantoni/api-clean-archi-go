package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/dimiantoni/api-clean-archi-go/api/presenter"
	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
	"github.com/dimiantoni/api-clean-archi-go/usecase/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_listUsers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("listUsers").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)
	u := &entity.User{
		ID: entity.NewID(),
	}
	m.EXPECT().
		ListUsers().
		Return([]*entity.User{u}, nil)
	ts := httptest.NewServer(listUsers(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listUsers_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listUsers(m))
	defer ts.Close()
	m.EXPECT().
		SearchUsers("dimi@gmail.com").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?email=dimi@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listUsers_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.User{
		ID: entity.NewID(),
	}
	m.EXPECT().
		SearchUsers("buster@gmail.com").
		Return([]*entity.User{u}, nil)
	ts := httptest.NewServer(listUsers(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?email=buster@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("createUser").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)

	m.EXPECT().
		CreateUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	h := createUser(m)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
	"name": "Buster",
	"email": "buster@gmail.com",
	"password": "123456",
	"address": "Lime street 129",
	"age": 18
	}`)
	resp, _ := http.Post(ts.URL+"/v1/user", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.User
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "Buster", fmt.Sprintf("%s", u.Name))
}

func Test_getUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("getUser").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	u := &entity.User{
		ID: entity.NewID(),
	}
	m.EXPECT().
		GetUser(u.ID).
		Return(u, nil)
	handler := getUser(m)
	r.Handle("/v1/user/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/user/" + u.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.User
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("deleteUser").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	u := &entity.User{
		ID: entity.NewID(),
	}
	m.EXPECT().DeleteUser(u.ID).Return(nil)
	handler := deleteUser(m)
	req, _ := http.NewRequest("DELETE", "/v1/user/"+u.ID.String(), nil)
	r.Handle("/v1/user/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
