// handler/response.go
package handler

type Response struct {
	Status  bool        `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data,omitempty"`
}
