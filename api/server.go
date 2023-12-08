package api

import (
	"encoding/json"
	hygiene "github.com/aviva-verde/tech-test-backend-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct{}

func (s Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api", s.getAuthorities)
	router.HandleFunc("/api/{authorityID}", s.getAuthority)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(srv.ListenAndServe(), router)
}

func (s Server) getAuthorities(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, "http://api.ratings.food.gov.uk/Authorities", nil)
	req.Header.Set("x-api-version", "2")
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	var fsaAuthorities hygiene.FSAAuthorities
	err := json.Unmarshal(body, &fsaAuthorities)
	if err != nil {
		panic("error")
	}

	var authorities []hygiene.Authority
	for _, authority := range fsaAuthorities.Authorities {
		authorities = append(authorities, hygiene.Authority{
			ID:   authority.ID,
			Name: authority.Name,
		})
	}

	data, _ := json.Marshal(authorities)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s Server) getAuthority(w http.ResponseWriter, r *http.Request) {
	authorityRating := []hygiene.AuthorityRating{
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
