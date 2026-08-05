package benchmark

import (
	"apiProfiler/internal/service"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func New() *Handler {
	return &Handler{
		service: service.New(),
	}
}

type SumResponse struct {
	Result int `json:"result"`
}

func (h *Handler) Sum(
	w http.ResponseWriter,
	r *http.Request,
) {

	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	resp := SumResponse{
		Result: h.service.Sum(a, b),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Users(
	w http.ResponseWriter,
	r *http.Request,
) {

	users := h.service.GetUsers()

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	if err := json.NewEncoder(buf).Encode(users); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}