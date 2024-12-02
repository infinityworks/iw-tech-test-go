package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger *slog.Logger
}

type Authority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthorityRating struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type FSAAuthorities struct {
	Authorities []FSAAuthority `json:"authorities"`
}

type FSAAuthority struct {
	ID   int    `json:"LocalAuthorityId"`
	Name string `json:"Name"`
}

func (s Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api", s.getAuthorities)
	router.HandleFunc("/api/{authorityID}", s.getAuthority)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	s.Logger.Info("Server started", slog.String("address", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		s.Logger.Error("Server failed to start",
			slog.String("address", srv.Addr),
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}

func (s Server) getAuthorities(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Getting authorities")
	req, _ := http.NewRequest(http.MethodGet, "http://api.ratings.food.gov.uk/Authorities", nil)
	req.Header.Set("x-api-version", "2")
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)

	var fsaAuthorities FSAAuthorities
	err := json.Unmarshal(body, &fsaAuthorities)
	if err != nil {
		s.Logger.Error("error unmarshalling authorities", "error", err)
	}

	var authorities []Authority
	for _, authority := range fsaAuthorities.Authorities {
		model := Authority(authority)
		authorities = append(authorities, model)
	}

	data, _ := json.Marshal(authorities)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s Server) getAuthority(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Getting authority")
	authorityRating := []AuthorityRating{
		{Name: "5-star", Value: 22.41},
		{Name: "4-star", Value: 43.13},
		{Name: "3-star", Value: 12.97},
		{Name: "2-star", Value: 1.54},
		{Name: "1-star", Value: 17.84},
		{Name: "Exempt", Value: 2.11},
	}

	data, _ := json.Marshal(authorityRating)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
