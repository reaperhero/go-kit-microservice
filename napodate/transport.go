package napodate

import (
	"context"
	"encoding/json"
	"net/http"
)

// 第一部分中，我们将请求和响应映射到它们的 JSON 实体。对于 statusRequest 和 getRequest 我们并不需要，因为没有有效载荷被发送到服务器。而 validateRequest 我们要传递一个要验证的日期，所以这里是 date 字段。
type getRequest struct{}

type getResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

// 我们将为传入的请求编写“解码器”，告诉服务他应该如何转换请求并将它们映射到正确的请求结构。我知道 get 和 status 是空的，但他们在那里为完整起见。
func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	return req, nil
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
