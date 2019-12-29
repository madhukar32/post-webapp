package posts

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Post data to store it in mem
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

// CreateRouter creates an httprouter and serves at port
func CreateRouter(port int32) {
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id:[0-9]+}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", deletePost).Methods("DELETE")
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

// get all posts present in the memory
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// get post with id given in the request
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, post := range posts {
		if post.ID == params["id"] {
			// post id found return the data
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	//return 404
	w.WriteHeader(http.StatusNotFound)
}

// create a new post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{}
	_ = json.NewDecoder(r.Body).Decode(post)
	// generate new ID
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, *post)
	json.NewEncoder(w).Encode(&post)
}

// update existing post
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			// post id found delete the exisitng item from posts array
			posts = append(posts[:index], posts[index+1:]...)
			post := &Post{}
			_ = json.NewDecoder(r.Body).Decode(post)
			post.ID = params["id"]
			posts = append(posts, *post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	//return 404
	w.WriteHeader(http.StatusNotFound)
}

// delete existing post
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			// post id found delete the exisitng item from posts array
			posts = append(posts[:index], posts[index+1:]...)
			return
		}
	}
	//return 404
	w.WriteHeader(http.StatusNotFound)
}
