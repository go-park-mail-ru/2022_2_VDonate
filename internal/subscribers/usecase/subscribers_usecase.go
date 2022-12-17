package subscribers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ztrue/tracerr"

	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
)

type usecase struct {
	subscribersMicroservice domain.SubscribersMicroservice
	userMicroservice        domain.UsersMicroservice
}

const expire = time.Minute * 5

func New(s domain.SubscribersMicroservice, u domain.UsersMicroservice) domain.SubscribersUseCase {
	return &usecase{
		subscribersMicroservice: s,
		userMicroservice:        u,
	}
}

func (u usecase) GetSubscribers(authorID uint64) ([]models.User, error) {
	s, err := u.subscribersMicroservice.GetSubscribers(authorID)
	if err != nil {
		return nil, err
	}

	subs := make([]models.User, 0)

	for _, userID := range s {
		// Notion: if there is an error while getting user, skip it
		user, _ := u.userMicroservice.GetByID(userID)
		subs = append(subs, user)
	}

	return subs, nil
}

func (u usecase) Subscribe(subscription models.Subscription, userID uint64, as models.AuthorSubscription) (interface{}, error) {
	if subscription.AuthorID == userID {
		return nil, domain.ErrBadRequest
	}

	subscription.SubscriberID = userID
	if utils.Empty(subscription.SubscriberID, subscription.AuthorID, subscription.AuthorSubscriptionID) {
		return "", domain.ErrBadRequest
	}

	payment := models.Payment{
		ID:     uuid.New().String(),
		FromID: subscription.SubscriberID,
		ToID:   subscription.AuthorID,
		SubID:  subscription.AuthorSubscriptionID,
		Price:  as.Price,
		Time:   time.Time{},
	}

	qiwiPayment := models.QiwiPayment{
		Amount: struct {
			Currency string `json:"currency"`
			Value    string `json:"value"`
		}{
			Currency: "RUB",
			Value:    strconv.FormatUint(as.Price, 10),
		},
		Comment:            fmt.Sprintf("UserID: %d\nPrice: %d\nSubcription: %s", userID, as.Price, as.Title),
		ExpirationDateTime: time.Now().Add(expire),
		Customer: struct {
			Account string `json:"account"`
		}{
			Account: strconv.FormatUint(userID, 10),
		},
	}

	p, err := json.Marshal(qiwiPayment)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	req, err := http.NewRequest(http.MethodPut, "https://api.qiwi.com/partner/bill/v1/bills/"+payment.ID, bytes.NewBuffer(p))
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJ2ZXJzaW9uIjoiUDJQIiwiZGF0YSI6eyJwYXlpbl9tZXJjaGFudF9zaXRlX3VpZCI6Ijgyem03Ny0wMCIsInVzZXJfaWQiOiI3OTc3NDU4MjM1NiIsInNlY3JldCI6IjkyYzg2OGUwZjQ5N2VkNWFmMDc3MWI2NzkxMzg5OTJhYjY0MWJhMjRiMDE4NjAyN2EwZjJhYTIxZmNjNmNhNTkifX0=")

	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	if response.StatusCode != http.StatusOK {
		return "", domain.ErrCreatePayment
	}
	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return "", tracerr.Wrap(err)
	}
	var qiwiResp models.QiwiPaymentStatus
	var qiwiErr models.QiwiErrorPaymentStatus

	if err = json.Unmarshal(resp, &qiwiResp); err != nil {
		if err = json.Unmarshal(resp, &qiwiErr); err != nil {
			return "", tracerr.Wrap(err)
		}
		return qiwiErr, nil
	}

	qiwiResp.PayUrl += "&successUrl=https://vdonate.ml/profile?id=" + strconv.FormatUint(payment.ToID, 10)

	go u.subscribersMicroservice.Subscribe(payment)

	return qiwiResp, nil
}

func (u usecase) Unsubscribe(userID, authorID uint64) error {
	if userID == 0 || authorID == 0 {
		return domain.ErrBadRequest
	}
	return u.subscribersMicroservice.Unsubscribe(models.Subscription{
		AuthorID:     authorID,
		SubscriberID: userID,
	})
}

func (u usecase) IsSubscriber(userID, authorID uint64) (bool, error) {
	s, err := u.subscribersMicroservice.GetSubscribers(authorID)
	if err != nil {
		return false, err
	}

	for _, id := range s {
		if id == userID {
			return true, nil
		}
	}

	return false, nil
}
