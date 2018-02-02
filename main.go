package main

import (
	"log"
	"net/http"

	"github.com/bosemian/go-webbord/app"
	"github.com/bosemian/go-webbord/ctrl"
)

func main() {
	mux := http.NewServeMux()
	api := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", middleware(api)))

	userCtrl := ctrl.NewUserController()
	forumCtrl := ctrl.NewForumController()
	topicCtrl := ctrl.NewTopicController()

	app.MountUserController(api, userCtrl)
	app.MountForumController(api, forumCtrl)
	app.MountTopicController(api, topicCtrl)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
