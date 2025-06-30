package addsvc

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(ep Endpoints) http.Handler {
	mux := http.NewServeMux()
	addHandler := httptransport.NewServer(
		ep.AddEndpoint,
		decodeAddRequest,
		encodeResponse,
	)
	mux.Handle("/add", addHandler)
	return mux
}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	query := r.URL.Query()
	x := query.Get("a")
	y := query.Get("b")
	a, err := strconv.Atoi(x)
	if err != nil {
		return nil, err
	}
	b, err := strconv.Atoi(y)
	if err != nil {
		return nil, err
	}
	return AddRequest{A: a, B: b}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
