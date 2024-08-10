package res

import (
	"fmt"
	"net/http"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResultData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResultHealth struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ResultError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (r Result) String() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s"}`, r.Code, r.Message)
}

func (rd ResultData) String() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s","data":%s}`, rd.Code, rd.Message, rd.Data)
}

func (rh ResultHealth) String() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s","status":"%s"}`, rh.Code, rh.Message, rh.Status)
}

func (re ResultError) String() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s","error":"%s"}`, re.Code, re.Message, re.Error)
}

func SendAuthorizationError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	result := ResultError{
		Code:    http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
		Error:   "Unauthorized access",
	}
	w.Write([]byte(result.String()))
}
