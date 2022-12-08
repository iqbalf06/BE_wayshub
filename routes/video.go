package routes

import (
	"wayshub/handlers"
	"wayshub/pkg/middleware"
	"wayshub/pkg/mysql"
	"wayshub/repositories"

	"github.com/gorilla/mux"
)

func VideoRoutes(r *mux.Router) {
	videoRepository := repositories.RepositoryVideo(mysql.DB)
	h := handlers.HandlerVideo(videoRepository)

	r.HandleFunc("/videos", middleware.Auth(h.FindVideos)).Methods("GET")
	r.HandleFunc("/video/{id}", middleware.Auth(h.GetVideo)).Methods("GET")
	r.HandleFunc("/video", middleware.Auth(middleware.UploadThumbnail(middleware.UploadVideo(h.AddVideo)))).Methods("POST")
	r.HandleFunc("/video/{id}", middleware.Auth(middleware.UploadThumbnail(middleware.UploadVideo(h.EditVideo)))).Methods("PATCH")
	r.HandleFunc("/video/{id}", middleware.Auth(h.DeleteVideo)).Methods("DELETE")
}
