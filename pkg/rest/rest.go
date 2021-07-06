package rest

import (
	"encoding/json"
	"fmt"
	"github.com/b-nova-openhub/gostapagen/pkg/gen"
	"github.com/b-nova-openhub/gostapagen/pkg/repo"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/page", getPage).Methods("GET")
	router.HandleFunc("/pages", getPages).Methods("GET")
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/generate", getGenerate).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("appPort"), router))
}

func getPage(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	pages := gen.GeneratedPages
	page := getPageById(v.Get("id"), pages)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&page)
}

func getPages(w http.ResponseWriter, r *http.Request) {
	pages := gen.GeneratedPages
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	status := gen.CurrentStatus
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func getGenerate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Generate Request: %+v\n", r)
	generated := gen.Generate(repo.RepoContents())
	fmt.Printf("Generate Response: %+v\n", generated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generated)
}

func getPageById(id string, pages []gen.StaticPage) *gen.StaticPage {
	var page *gen.StaticPage
	for _, p := range pages {
		if p.Permalink == id {
			page = &p
		}
	}
	return page
}
