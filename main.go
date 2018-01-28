package main

import (
	"log"
	"net/http"

	"github.com/bosemian/go-webbord/app"
	"github.com/bosemian/go-webbord/ctrl"
)

func main() {
	mux := http.NewServeMux()
	userCtrl := ctrl.NewUserController()
	forumCtrl := ctrl.NewForumController()

	app.MountUserController(mux, userCtrl)
	app.MountForumController(mux, forumCtrl)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
