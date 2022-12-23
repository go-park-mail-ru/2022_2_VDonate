package subscribers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	token string
}

const (
	expire             = time.Minute * 5
	minSum             = 75
	commission         = 0.95
	qiwiBankCommission = 50
)

func New(s domain.SubscribersMicroservice, u domain.UsersMicroservice, token string) domain.SubscribersUseCase {
	return &usecase{
		subscribersMicroservice: s,
		userMicroservice:        u,
		token:                   token,
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

func (u usecase) Follow(subscriberID, authorID uint64) error {
	if subscriberID == 0 || authorID == 0 {
		return domain.ErrBadRequest
	}
	err := u.subscribersMicroservice.Follow(subscriberID, authorID)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("QIWI_PRIVATE")))

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

func (u usecase) CardValidation(card string) (models.WithdrawValidation, error) {
	p, err := json.Marshal(models.WithdrawCard{Account: card})
	if err != nil {
		return models.WithdrawValidation{}, tracerr.Wrap(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://edge.qiwi.com/sinap/api/refs/bd6fb248-2bdf-49ed-bcb2-9b0a789cfde8/containers", bytes.NewBuffer(p))
	if err != nil {
		return models.WithdrawValidation{}, tracerr.Wrap(err)
	}
	req.Header.Set("Accept", "application/vnd.qiwi.v1+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", u.token))

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.WithdrawValidation{}, tracerr.Wrap(err)
	}
	if resp.StatusCode != http.StatusOK {
		return models.WithdrawValidation{}, domain.ErrCreatePayment
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.WithdrawValidation{}, tracerr.Wrap(err)
	}

	var v models.WithdrawValidation

	if err = json.Unmarshal(body, &v); err != nil {
		return models.WithdrawValidation{}, tracerr.Wrap(err)
	}

	return v, nil
}

func (u usecase) WithdrawCard(userID uint64, card, provider string) (models.WithdrawInfo, error) {
	if userID == 0 || card == "" || provider == "" {
		return models.WithdrawInfo{}, domain.ErrBadRequest
	}

	acc, err := u.userMicroservice.GetByID(userID)
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	if acc.Balance < minSum {
		return models.WithdrawInfo{}, domain.ErrNotEnoughMoney
	}

	p, err := json.Marshal(models.WithdrawPayment{
		ID: strconv.FormatInt(1000*time.Now().Unix(), 10),
		Sum: struct {
			Amount   float64 `json:"amount,required"`
			Currency string  `json:"currency,required"`
		}(struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		}{
			Amount:   float64(acc.Balance-qiwiBankCommission)*commission + qiwiBankCommission,
			Currency: "643",
		}),
		PaymentMethod: struct {
			Type      string `json:"type,required"`
			AccountId string `json:"accountId,required"`
		}(struct {
			Type      string `json:"type"`
			AccountId string `json:"accountId"`
		}{
			Type:      "Account",
			AccountId: "643",
		}),
		Fields: struct {
			Account string `json:"account,required"`
		}(struct {
			Account string `json:"account"`
		}{
			Account: card,
		}),
	})
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://edge.qiwi.com/sinap/api/v2/terms/"+provider+"/payments", bytes.NewBuffer(p))
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", u.token))

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}
	if resp.StatusCode != http.StatusOK {
		return models.WithdrawInfo{}, domain.ErrCreatePayment
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	var v models.WithdrawInfo
	var e models.WithdrawError

	if err = json.Unmarshal(body, &v); err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	if v.Id == "" {
		if err = json.Unmarshal(body, &e); err != nil {
			return models.WithdrawInfo{}, tracerr.Wrap(err)
		}

		return models.WithdrawInfo{}, tracerr.Wrap(e)
	}

	return v, nil
}

func (u usecase) WithdrawQiwi(userID uint64, phone string) (models.WithdrawInfo, error) {
	if userID == 0 || phone == "" {
		return models.WithdrawInfo{}, domain.ErrBadRequest
	}

	acc, err := u.userMicroservice.GetByID(userID)
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	p, err := json.Marshal(models.WithdrawPayment{
		ID: strconv.FormatInt(1000*time.Now().Unix(), 10),
		Sum: struct {
			Amount   float64 `json:"amount,required"`
			Currency string  `json:"currency,required"`
		}(struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		}{
			Amount:   float64(acc.Balance) * commission,
			Currency: "643",
		}),
		PaymentMethod: struct {
			Type      string `json:"type,required"`
			AccountId string `json:"accountId,required"`
		}(struct {
			Type      string `json:"type"`
			AccountId string `json:"accountId"`
		}{
			Type:      "Account",
			AccountId: "643",
		}),
		Fields: struct {
			Account string `json:"account,required"`
		}(struct {
			Account string `json:"account"`
		}{
			Account: "+" + phone,
		}),
	})
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://edge.qiwi.com/sinap/api/v2/terms/99/payments", bytes.NewBuffer(p))
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", u.token))

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	var v models.WithdrawInfo
	var e models.WithdrawError

	if err = json.Unmarshal(body, &v); err != nil {
		return models.WithdrawInfo{}, tracerr.Wrap(err)
	}

	if v.Id == "" {
		if err = json.Unmarshal(body, &e); err != nil {
			return models.WithdrawInfo{}, tracerr.Wrap(err)
		}

		return models.WithdrawInfo{}, tracerr.Wrap(e)
	}

	return v, nil
}

func (u usecase) Withdraw(userID uint64, phone, card string) (models.WithdrawInfo, error) {
	var provider string

	if len(card) != 0 {
		val, err := u.CardValidation(card)
		if err != nil {
			return models.WithdrawInfo{}, tracerr.Wrap(err)
		}

		provider = val.Elements[0].Value

		info, err := u.WithdrawCard(userID, card, provider)
		if err != nil {
			return models.WithdrawInfo{}, tracerr.Wrap(err)
		}

		if err = u.userMicroservice.DropBalance(userID); err != nil {
			return models.WithdrawInfo{}, err
		}

		return info, nil
	}

	if len(phone) != 0 {
		info, err := u.WithdrawQiwi(userID, phone)
		if err != nil {
			return models.WithdrawInfo{}, err
		}

		if err = u.userMicroservice.DropBalance(userID); err != nil {
			return models.WithdrawInfo{}, err
		}

		return info, nil
	}

	return models.WithdrawInfo{}, domain.ErrBadRequest
}
