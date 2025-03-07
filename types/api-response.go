package types

type APIResponse[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}
