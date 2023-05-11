package router

import (
	"historycznymonolog/api"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	//	Users requests
	router.HandleFunc("/users", api.GetAllUsers).Methods("GET","OPTIONS")
	router.HandleFunc("/users/username/{username}", api.GetUserByUsername).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/{id}", api.GetUserByID).Methods("GET", "OPTIONS")

	router.HandleFunc("/users/{id}", api.DeleteUserByID).Methods("DELETE", "OPTIONS")

	//	Comments requests
	router.HandleFunc("/comments", api.GetAllComments).Methods("GET", "OPTIONS")
	router.HandleFunc("/comments/post/{postId}", api.GetCommentsForPost).Methods("GET","OPTIONS")
	router.HandleFunc("/comments/user/{authorId}", api.GetCommentsForUser).Methods("GET","OPTIONS")

	router.HandleFunc("/comments", api.AddComment).Methods("POST", "OPTIONS")

	//	Categories requests
	router.HandleFunc("/categories", api.GetAllCategories).Methods("GET", "OPTIONS")
	router.HandleFunc("/categories/{id}", api.GetCategoryByID).Methods("GET", "OPTIONS")

	router.HandleFunc("/categories", api.AddCategory).Methods("POST", "OPTIONS")

	//	Posts requests
    router.HandleFunc("/posts", api.GetAllPosts).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts/{id}", api.GetPostById).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts/title/{title}", api.GetPostByTitle).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts/author/{authorId}", api.GetPostByAuthor).Methods("GET", "OPTIONS")
	//router.HandleFunc("/posts/tag/{tag}", api.GetPostByTag).Methods("GET", "OPTIONS")	TODO

    router.HandleFunc("/posts", api.AddSinglePost).Methods("POST", "OPTIONS")
	router.HandleFunc("/posts/{id}", api.DeletePost).Methods("DELETE", "OPTIONS")
	
	//	Contents requests
	router.HandleFunc("/contents", api.GetAllContents).Methods("GET", "OPTIONS")
	router.HandleFunc("/contents/author/{authorId}", api.GetContentsForUser).Methods("GET","OPTIONS")
	router.HandleFunc("/contents/id/{id}", api.GetContentByID).Methods("GET","OPTIONS")

	router.HandleFunc("/contents", api.AddContent).Methods("POST", "OPTIONS")


	//	Redundant requests
	//router.HandleFunc("/users/password/{id}", api.ChangeUserPassword).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/users/email/{id}", api.ChangeUserEmail).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/register", api.RegisterUser).Methods("POST")
	//router.HandleFunc("/login", api.Login).Methods("POST")

	return router
}