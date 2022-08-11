package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
	GetUserRequest struct {
		Id string `json:"id"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, respone any) error {
	return json.NewEncoder(w).Encode(respone)
}

func decodeUserReq(ctx context.Context, r *http.Request) (any, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (any, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}

	return req, nil
}
