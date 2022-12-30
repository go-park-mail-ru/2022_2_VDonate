package models

import "time"

type Payment struct {
	ID     string    `json:"id"`
	FromID uint64    `json:"from_id"`
	ToID   uint64    `json:"to_id"`
	SubID  uint64    `json:"sub_id"`
	Price  uint64    `json:"price"`
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

type Withdraw struct {
	UserID uint64 `json:"userID"`
	Phone  string `json:"phone"`
	Card   string `json:"card"`
}

type WithdrawCard struct {
	Account string `json:"account"`
}

type WithdrawPayment struct {
	ID  string `json:"id"`
	Sum struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"sum"`
	PaymentMethod struct {
		Type      string `json:"type"`
		AccountId string `json:"accountId"`
	} `json:"paymentMethod"`
	Fields struct {
		Account string `json:"account"`
	} `json:"fields"`
}

type WithdrawInfo struct {
	Id     string `json:"id"`
	Terms  string `json:"terms"`
	Fields struct {
		Account string `json:"account"`
	} `json:"fields"`
	Sum struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"sum"`
	Transaction struct {
		Id    string `json:"id"`
		State struct {
			Code string `json:"code"`
		} `json:"state"`
	} `json:"transaction"`
	Source  string `json:"source"`
	Comment string `json:"comment"`
}

type WithdrawValidation struct {
	Elements []struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"elements"`
}

type WithdrawError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	TxnId   string `json:"txnId"`
}

func (w WithdrawError) Error() string {
	return w.Message
}

type QiwiPayment struct {
	Amount struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	} `json:"amount"`
	Comment            string    `json:"comment"`
	ExpirationDateTime time.Time `json:"expirationDateTime"`
	Customer           struct {
		Account string `json:"account"`
	} `json:"customer"`
}

type QiwiPaymentStatus struct {
	SiteId string `json:"siteId"`
	BillId string `json:"billId"`
	Amount struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	} `json:"amount"`
	Status struct {
		Value           string    `json:"value"`
		ChangedDateTime time.Time `json:"changedDateTime"`
	} `json:"status"`
	Customer struct {
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Account string `json:"account"`
	} `json:"customer"`
	CustomFields struct {
		PaySourcesFilter string `json:"paySourcesFilter"`
		ThemeCode        string `json:"themeCode"`
		YourParam1       string `json:"yourParam1"`
		YourParam2       string `json:"yourParam2"`
	} `json:"customFields"`
	Comment            string    `json:"comment"`
	CreationDateTime   time.Time `json:"creationDateTime"`
	ExpirationDateTime time.Time `json:"expirationDateTime"`
	PayUrl             string    `json:"payUrl"`
}

type QiwiErrorPaymentStatus struct {
	ServiceName string    `json:"serviceName"`
	ErrorCode   string    `json:"errorCode"`
	Description string    `json:"description"`
	UserMessage string    `json:"userMessage"`
	DateTime    time.Time `json:"dateTime"`
	TraceId     string    `json:"traceId"`
}
