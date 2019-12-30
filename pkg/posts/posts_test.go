package posts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {

	newPost := &Post{
		Title: "Test Create",
		Body:  "Hello World!",
	}
	jsonBody, err := json.Marshal(newPost)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/posts", createPost)
	router.ServeHTTP(res, req)

	resPost := &Post{}
	json.Unmarshal(res.Body.Bytes(), resPost)
	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, newPost.Title, resPost.Title, "Title is as expected")
	assert.Equal(t, newPost.Body, resPost.Body, "Body is as expected")
	posts = []Post{}
}

func TestGetPost(t *testing.T) {

	newPost := &Post{
		Title: "Test Get",
		Body:  "Hello World!",
		ID:    "12345",
	}
	posts = append(posts, *newPost)
	jsonBody, err := json.Marshal(newPost)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/posts/12345", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/posts/{id}", getPost)
	router.ServeHTTP(res, req)

	resPost := &Post{}
	json.Unmarshal(res.Body.Bytes(), resPost)
	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, newPost.Title, resPost.Title, "Body title is as expected")
	assert.Equal(t, newPost.Body, resPost.Body, "Body title is as expected")
	posts = []Post{}
}

func TestUpdatePost(t *testing.T) {

	newPost := &Post{
		Title: "Test Create",
		Body:  "Hello World!",
		ID:    "12345",
	}
	posts = append(posts, *newPost)

	updateP := &Post{
		Title: "Test Update",
		Body:  "Hello World!",
		ID:    "12345",
	}
	jsonBody, err := json.Marshal(updateP)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/posts/12345", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/posts/{id}", updatePost)
	router.ServeHTTP(res, req)

	resPost := &Post{}
	json.Unmarshal(res.Body.Bytes(), resPost)
	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, updateP.Title, resPost.Title, "Body title is as expected")
	assert.Equal(t, updateP.Body, resPost.Body, "Body title is as expected")
	posts = []Post{}
}

func TestDeletePost(t *testing.T) {

	newPost := &Post{
		Title: "Test Create",
		Body:  "Hello World!",
		ID:    "12345",
	}
	posts = append(posts, *newPost)

	req, err := http.NewRequest("DELETE", "/posts/12345", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/posts/{id}", deletePost)
	router.ServeHTTP(res, req)

	resPost := &Post{}
	json.Unmarshal(res.Body.Bytes(), resPost)
	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, 0, len(posts), "length of psts is zero")
	posts = []Post{}
}
