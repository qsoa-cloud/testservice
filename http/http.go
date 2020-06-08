package http

import (
	"fmt"
	"net/http"
	"strconv"

	"testservice/grpc/pb"
)

type Handler struct {
	Client pb.TestClient
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n1, err := strconv.ParseInt(r.FormValue("n1"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	n2, err := strconv.ParseInt(r.FormValue("n2"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Client.Sum(r.Context(), &pb.SumReq{
		N1: n1,
		N2: n2,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprintf(w, "Sum = %d", resp.Sum)
}
