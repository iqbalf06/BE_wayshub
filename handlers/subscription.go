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
	userId := int(userInfo["id"].(float64))

	// subscribe, _ := strconv.Atoi(r.FormValue("subscribe"))

	userchanelId, _ := strconv.Atoi(r.FormValue("user_channel_id"))
	println("ini apa ? ", userchanelId)
	request := subscriptiondto.Subscriber{
		UserChannelId: userchanelId,
	}

	// UserChannelId := 4
	subscription := models.Subscription{
		UserChannelId: request.UserChannelId,
		ChannelId:     userId,
	}
	fmt.Println(subscription)

	subscription, _ = h.SubscriptionRepository.AddSubscription(subscription)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	subscription, _ = h.SubscriptionRepository.GetSubscription(subscription.ID)

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
