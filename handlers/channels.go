package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	authdto "wayshub/dto/auth"
	channelsdto "wayshub/dto/channels"
	dto "wayshub/dto/result"
	"wayshub/models"
	"wayshub/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerChannel struct {
	ChannelRepository repositories.ChannelRepository
}

func HandlerChannel(ChannelRepository repositories.ChannelRepository) *handlerChannel {
	return &handlerChannel{ChannelRepository}
}

func (h *handlerChannel) FindChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channels, err := h.ChannelRepository.FindChannels()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range channels {
		channels[i].Photo = os.Getenv("PATH_FILE") + p.Photo
	}

	for i, p := range channels {
		channels[i].Cover = os.Getenv("PATH_FILE") + p.Cover
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: channels}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerChannel) GetChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channel, err := h.ChannelRepository.GetChannel(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: channel}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerChannel) EditChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ContexPhoto := r.Context().Value("dataPhoto")
	filephoto := ContexPhoto.(string)
	ContexCover := r.Context().Value("dataCover")
	filecover := ContexCover.(string)

	request := authdto.RegisterRequest{
		Channelname: r.FormValue("channelName"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channel, err := h.ChannelRepository.GetChannel(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Channelname != "" {
		channel.Channelname = request.Channelname
	}

	if request.Email != "" {
		channel.Email = request.Email
	}

	if request.Password != "" {
		channel.Password = request.Password
	}

	if filephoto != "false" {
		channel.Photo = filephoto
	}

	if filecover != "false" {
		channel.Cover = filecover
	}

	data, err := h.ChannelRepository.EditChannel(channel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerChannel) DeleteChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userID := int(userInfo["id"].(float64))

	if userID != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "you're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	channel, err := h.ChannelRepository.GetChannel(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ChannelRepository.DeleteChannel(channel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.Channel) channelsdto.DeleteResponse {
	return channelsdto.DeleteResponse{
		ID: u.ID,
	}
}
