package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BlaiseRitchie/SakuraconGaming/server/internal/gameroom"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("port", "8080")
	port := viper.GetString("port")

	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)

	mux.HandleFunc("/admin/stations", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			stations, err := gameroom.GetStations()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(stations)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.Station
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateStation(args.ConsoleID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		case http.MethodDelete:
			var args gameroom.Station
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.DeleteStation(args.ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})

	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}

func respondErr(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	fmt.Fprint(w, err)
	log.Println(err)
}
