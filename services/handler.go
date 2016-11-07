package services

import (
	"encoding/json"
	"fmt"
	Const "github.com/gauravtayal0/sysutil/constants"
	"net/http"
	"time"
)

type APIResponse struct {
	Data interface{} `json:"data"`
}

func NewAPIResponse(data interface{}) *APIResponse {
	apiResp := new(APIResponse)
	apiResp.Data = data
	return apiResp
}

func WriteResponse(w http.ResponseWriter, resp interface{}) {
	fmt.Println(resp)
	jsonResp, err := json.MarshalIndent(resp, Const.JsonIndent, Const.JsonPrefix)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(jsonResp)
	return
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	timestamp := fmt.Sprintf(" (%v) ", time.Now().UnixNano())
	http.Error(w, http.StatusText(statusCode)+timestamp, statusCode)
	return
}
