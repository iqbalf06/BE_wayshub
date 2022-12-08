package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	dto "wayshub/dto/result"
	subscriptiondto "wayshub/dto/subscription"
	"wayshub/models"
	"wayshub/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerSubscription struct {
	SubscriptionRepository repositories.SubscriptionRepository
}

func HandlerSubscription(SubscriptionRepository repositories.SubscriptionRepository) *handlerSubscription {
	return &handlerSubscription{SubscriptionRepository}
}

func (h *handlerSubscription) AddSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["Id"].(float64))
	fmt.Println(userId)

	subscribe, _ := strconv.Atoi(r.FormValue("subscribe"))
	request := subscriptiondto.Subscriber{
		Subscribe: subscribe,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	subscription := models.Subscription{
		ChannelID: userId,
	}

	subscription, err = h.SubscriptionRepository.AddSubscription(subscription)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	subscription, _ = h.SubscriptionRepository.GetSubscription(userId)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: subscription}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscription) GetSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	subscription, err := h.SubscriptionRepository.GetSubscription(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: subscription}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscription) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	subscription, err := h.SubscriptionRepository.GetSubscription(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.SubscriptionRepository.Unsubscribe(subscription)
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
