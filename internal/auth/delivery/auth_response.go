package httpAuth

type AuthResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

const (
	statusOk = "OK"
	statusFail = "FAIL"
)

const (
	successGetCSRF = "successfully get csrf"
	invalidCSRF = "invalid csrf"
)

func AuthCSRFSuccessResponse() *AuthResponse {
	return &AuthResponse{
		Status: statusOk,
		Result: successGetCSRF,
	}
}

func AuthMiddlewareErrorResponse(err error) *AuthResponse {
	return &AuthResponse{
		Status: statusFail,
		Result: invalidCSRF,
	}
}
