package subscribersMicroservice

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ztrue/tracerr"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"

	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type SubscribersMicroservice struct {
	subscribersClient protobuf.SubscribersClient
}

func New(subscribersClient protobuf.SubscribersClient) domain.SubscribersMicroservice {
	return &SubscribersMicroservice{
		subscribersClient: subscribersClient,
	}
}

func (m SubscribersMicroservice) GetSubscribers(userID uint64) ([]uint64, error) {
	subscribers, err := m.subscribersClient.GetSubscribers(context.Background(), &userProto.UserID{
		UserId: userID,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	res := make([]uint64, 0)
	for _, id := range subscribers.GetIds() {
		res = append(res, id.GetUserId())
	}

	return res, nil
}

func (m SubscribersMicroservice) Subscribe(payment models.Payment) {
	log := logrus.New()

	req, err := http.NewRequest(http.MethodGet, "https://api.qiwi.com/partner/bill/v1/bills/"+payment.ID, nil)
	if err != nil {
		log.Error(tracerr.Wrap(err))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJ2ZXJzaW9uIjoiUDJQIiwiZGF0YSI6eyJwYXlpbl9tZXJjaGFudF9zaXRlX3VpZCI6Ijgyem03Ny0wMCIsInVzZXJfaWQiOiI3OTc3NDU4MjM1NiIsInNlY3JldCI6IjkyYzg2OGUwZjQ5N2VkNWFmMDc3MWI2NzkxMzg5OTJhYjY0MWJhMjRiMDE4NjAyN2EwZjJhYTIxZmNjNmNhNTkifX0=")

	client := http.Client{}
	for {
		time.Sleep(1 * time.Second)
		response, err := client.Do(req)
		if err != nil {
			log.Error(tracerr.Wrap(err))
		}

		if response.StatusCode != http.StatusOK {
			log.Error(tracerr.Wrap(err))
		}
		defer response.Body.Close()

		resp, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error(tracerr.Wrap(err))
		}
		var qiwiResp models.QiwiPaymentStatus
		var qiwiErr models.QiwiErrorPaymentStatus

		if err = json.Unmarshal(resp, &qiwiResp); err != nil {
			if err = json.Unmarshal(resp, &qiwiErr); err != nil {
				log.Error(tracerr.Wrap(err))
			}
			log.Error(tracerr.Wrap(err))
		}

		if qiwiResp.Status.Value == "PAID" {
			break
		} else if qiwiResp.Status.Value == "REJECTED" || qiwiResp.Status.Value == "EXPIRED" {
			if _, err = m.subscribersClient.ChangePaymentStatus(context.Background(), &protobuf.StatusAndID{
				Status: qiwiResp.Status.Value,
				Id:     payment.ID,
			}); err != nil {
				log.Error(err)
			}
			return
		}
	}

	_, err = m.subscribersClient.Subscribe(context.Background(), &protobuf.Payment{
		ID:     payment.ID,
		ToID:   payment.ToID,
		FromID: payment.FromID,
		SubID:  payment.SubID,
		Price:  payment.Price,
		Time:   timestamppb.New(payment.Time),
	})
	if err != nil {
		logger.GetInstance().Error(err)
	}
}

func (m SubscribersMicroservice) Unsubscribe(subscriber models.Subscription) error {
	_, err := m.subscribersClient.Unsubscribe(context.Background(), &userProto.UserAuthorPair{
		UserId:   subscriber.SubscriberID,
		AuthorId: subscriber.AuthorID,
	})

	return tracerr.Wrap(err)
}
