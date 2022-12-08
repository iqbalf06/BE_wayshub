package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	dto "wayshub/dto/result"
	videodto "wayshub/dto/video"
	"wayshub/models"
	"wayshub/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerVideo struct {
	VideoRepository repositories.VideoRepository
}

func HandlerVideo(VideoRepository repositories.VideoRepository) *handlerVideo {
	return &handlerVideo{VideoRepository}
}

func (h *handlerVideo) AddVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get dataFile from midleware and store to filethumbnail variable here ...
	dataContex := r.Context().Value("dataThumbnail")
	filethumbnail := dataContex.(string)

	videoContex := r.Context().Value("dataVideo")
	filevideo := videoContex.(string)

	request := videodto.VideoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video := models.Video{
		Title:       request.Title,
		Thumbnail:   filethumbnail,
		Description: request.Description,
		Video:       filevideo,
		CreatedAt:   time.Now(),
	}

	video, err = h.VideoRepository.AddVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video, _ = h.VideoRepository.GetVideo(video.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: video}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) FindVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	videos, err := h.VideoRepository.FindVideos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range videos {
		videos[i].Video = os.Getenv("PATH_FILE") + p.Video
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: videos}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var video models.Video
	video, err := h.VideoRepository.GetVideo(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video.Video = os.Getenv("PATH_FILE") + video.Video

	videos, err := h.VideoRepository.GetVideo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: videos}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) EditVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataThumbnail")
	filethumbnail := dataContex.(string)

	videoContex := r.Context().Value("dataVideo")
	filevideo := videoContex.(string)

	request := videodto.EditVideoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	video, _ := h.VideoRepository.GetVideo(id)

	if request.Title != "" {
		video.Title = request.Title
	}

	if filethumbnail != "false" {
		video.Thumbnail = filethumbnail
	}

	if request.Description != "" {
		video.Description = request.Description
	}

	if filevideo != "false" {
		video.Video = filevideo
	}

	video, err = h.VideoRepository.EditVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: video}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	video, err := h.VideoRepository.GetVideo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = h.VideoRepository.DeleteVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data := models.Video{
		ID: video.ID,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: convertResponseVideo(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseVideo(u models.Video) videodto.DeleteResponse {
	return videodto.DeleteResponse{
		ID: u.ID,
	}
}
