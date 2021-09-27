package controllers

import "github.com/cepot-blip/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Routes
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Routes
	s.Router.HandleFunc("/user_login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routess
	s.Router.HandleFunc("/users_create", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users_read", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users_update/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users_delete/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	s.Router.HandleFunc("/users_find/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")

	//Posts routes
	s.Router.HandleFunc("/posts_create", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts_read", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts_update/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts_delete/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
	s.Router.HandleFunc("/posts_find/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("POST")

}