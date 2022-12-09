package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	commentdto "wayshub/dto/comment"
	dto "wayshub/dto/result"
	"wayshub/models"
	"wayshub/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerComment struct {
	CommentRepository repositories.CommentRepository
}

func HandlerComment(CommentRepository repositories.CommentRepository) *handlerComment {
	return &handlerComment{CommentRepository}
}

func (h *handlerComment) AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	// println(r.Context())
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	fmt.Println(userInfo, " ini user info")
	userId := int(userInfo["id"].(float64))
	fmt.Println(userId, "masuk sini ?")

	request := commentdto.CommentRequest{
		Comment: r.FormValue("comment"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	comment := models.Comment{
		ChannelID: userId,
		Comment:   request.Comment,
		CreatedAt: time.Now(),
	}

	comment, err = h.CommentRepository.AddComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	comment, _ = h.CommentRepository.GetComment(comment.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: comment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerComment) FindComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments, err := h.CommentRepository.FindComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: comments}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerComment) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	comment, err := h.CommentRepository.GetComment(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: comment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerComment) EditComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := commentdto.CommentRequest{
		Comment: r.FormValue("comment"),
	}

	comment, err := h.CommentRepository.GetComment(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Comment != "" {
		comment.Comment = request.Comment
	}

	data, err := h.CommentRepository.EditComment(comment)
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

func (h *handlerComment) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	comment, err := h.CommentRepository.GetComment(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CommentRepository.DeleteComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: convertResponseComment(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseComment(u models.Comment) commentdto.DeleteResponse {
	return commentdto.DeleteResponse{
		ID: u.ID,
	}
}
