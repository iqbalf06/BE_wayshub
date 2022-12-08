package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	authdto "wayshub/dto/auth"
	dto "wayshub/dto/result"
	"wayshub/models"
	"wayshub/pkg/bcrypt"
	jwtToken "wayshub/pkg/jwt"
	"wayshub/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.Channel{
		Email:       request.Email,
		Password:    password,
		Channelname: request.Channelname,
		Description: request.Description,
	}

	_, err = h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	channelName := authdto.RegisterResponse{
		Email: request.Email,
		Channelname: request.Channelname,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: channelName}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	channels := models.Channel{
		Email:    request.Email,
		Password: request.Password,
	}

	channel, err := h.AuthRepository.Login(channels.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
fmt.Println(channel)
	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, channel.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = channel.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorized")
		return
	}

	loginResponse := authdto.LoginResponse{
		Channelname: channel.Channelname,
		Email:       channel.Email,
		Token:       token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: "success", Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	channelID := int(userInfo["id"].(float64))

	// Check User by Id
	channel, err := h.AuthRepository.Getchannel(channelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       channel.ID,
		Channelname:     channel.Channelname,
		Email:    channel.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: "success", Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}