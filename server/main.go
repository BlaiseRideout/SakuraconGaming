package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/BlaiseRitchie/SakuraconGaming/server/internal/gameroom"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.AutomaticEnv()
	viper.SetDefault("port", "8080")
	port := viper.GetString("port")

	router := mux.NewRouter().StrictSlash(true)
	handler := handlers.CORS()(router)
	handler = handlers.LoggingHandler(os.Stdout, handler)

	handler = handlers.RecoveryHandler()(handler)

	imagePath := "/images/"
	s := http.StripPrefix(imagePath, http.FileServer(http.Dir("."+imagePath)))
	router.PathPrefix(imagePath).Handler(s)

	// TODO: break all this out into package
	// Stations

	router.HandleFunc("/admin/stations", func(w http.ResponseWriter, r *http.Request) {
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
			err = validateNonZero(args)
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
		}
	})
	router.HandleFunc("/admin/stations/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteStation(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.Station
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateStation(ID, args.ConsoleID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		}
	})

	// Consoles
	router.HandleFunc("/admin/consoles", func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			consoles, err := gameroom.GetConsoles()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(consoles)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.Console
			args.Name, args.Image, err = acceptImage(r, "/images/consoles/")
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateConsole(args.Name, args.Image)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})
	router.HandleFunc("/admin/consoles/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteConsole(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.Console
			args.Name, args.Image, err = acceptImage(r, "/images/consoles/")
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateConsole(ID, args.Name, args.Image)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		}
	})

	// Controllers
	router.HandleFunc("/admin/controllers", func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			controllers, err := gameroom.GetControllers()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(controllers)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.Controller
			args.Name, args.Image, err = acceptImage(r, "/images/controllers/")
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateController(args.Name, args.Image, args.Count)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})
	router.HandleFunc("/admin/controllers/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteController(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.Controller
			args.Name, args.Image, err = acceptImage(r, "/images/controllers/")
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateController(ID, args.Name, args.Image, args.Count)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		}
	})

	//ConsoleControllers

	router.HandleFunc("/admin/console_controllers", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			consoleControllers, err := gameroom.GetConsoleControllers()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(consoleControllers)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.ConsoleController
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateConsoleController(args.ConsoleID, args.ControllerID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})
	router.HandleFunc("/admin/console_controllers/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteConsoleController(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.ConsoleController
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateConsoleController(ID, args.ConsoleID, args.ControllerID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		}
	})

	// Games
	router.HandleFunc("/admin/games", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			games, err := gameroom.GetGames()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(games)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.Game
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateGame(args.Name, args.ConsoleID, args.Count)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})
	router.HandleFunc("/admin/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteGame(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.Game
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateGame(ID, args.Name, args.ConsoleID, args.Count)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		}
	})

	// Barcodes
	router.HandleFunc("/admin/barcodes", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			barcodes, err := gameroom.GetBarcodes()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(barcodes)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.Barcode
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateBarcode(args.GameID, args.Barcode)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})
	router.HandleFunc("/admin/barcodes/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteBarcode(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.Barcode
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateBarcode(ID, args.GameID, args.Barcode)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")

		}
	})

	//Transactions

	router.HandleFunc("/admin/transactions", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case http.MethodGet:
			transactions, err := gameroom.GetTransactions()
			if err != nil {
				respondErr(w, err)
				return
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(true)
			err = enc.Encode(transactions)
			if err != nil {
				respondErr(w, err)
				return
			}
		case http.MethodPost:
			var args gameroom.Transaction
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.CreateTransaction(args.Type, args.BadgeID, args.StationID, args.GameID, args.ControllerID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		}
	})
	router.HandleFunc("/admin/transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

		if err != nil {
			respondErr(w, err)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			err := gameroom.DeleteTransaction(ID)
			if err != nil {
				respondErr(w, err)
				return
			}
			fmt.Fprint(w, "ok")
		case http.MethodPut:
			var args gameroom.Transaction
			err := json.NewDecoder(r.Body).Decode(&args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = validateNonZero(args)
			if err != nil {
				respondErr(w, err)
				return
			}
			err = gameroom.UpdateTransaction(ID, args.Type, args.BadgeID, args.StationID, args.GameID, args.ControllerID)
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

func validateNonZero(v interface{}) error {
	vs := reflect.ValueOf(v)
	ts := reflect.TypeOf(v)
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		vs = reflect.ValueOf(v).Elem()
		ts = reflect.TypeOf(v).Elem()
	}

	for i := 0; i < vs.NumField(); i++ {
		if ts.Field(i).Name == "ID" {
			continue
		}
		f := vs.Field(i)
		// log.Println(path+"."+tag, "value:", f.Interface())
		name := ts.Field(i).Name
		var err error
		switch f.Kind() {
		case reflect.Int:
			err = mustInt(name, f.Interface().(int))
		case reflect.String:
			err = mustString(name, f.Interface().(string))
		case reflect.Slice:
			switch ts.Field(i).Type.Elem().Kind() {
			case reflect.String:
				err = mustListString(name, f.Interface().([]string))
			default:
				return fmt.Errorf("UNHANDLED SLICE TYPE: %s at %s",
					ts.Field(i).Type.Elem().Kind(),
					name)
			}
		case reflect.Struct:
			err = validateNonZero(f.Interface())
			// Return early since we already wrapped the error at the lowest level
			if err != nil {
				return err
			}
		case reflect.Ptr:
			if f.IsNil() {
				return fmt.Errorf("%s missing from config", name)
			}
			err = validateNonZero(f.Interface())
			// Return early since we already wrapped the error at the lowest level
			if err != nil {
				return err
			}
		case reflect.Map:
			for _, k := range f.MapKeys() {
				err = validateNonZero(f.MapIndex(k).Interface())
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("UNHANDLED TYPE: %s at %s", f.Kind(), name)
		}
		if err != nil {
			return errors.Wrap(err, "At key: "+name)
		}
	}

	return nil
}

// mustString ensures that a string is not empty.
func mustString(key, val string) error {
	if val == "" {
		return fmt.Errorf("key %q is empty", key)
	}
	return nil
}

// mustListString ensures that a list of string values has a length of greater
// than zero, and ensures that each value is not an empty string.
func mustListString(key string, vals []string) error {
	if len(vals) == 0 {
		return fmt.Errorf("list of values %q has zero length", key)
	}
	for _, val := range vals {
		if err := mustString(key, val); err != nil {
			return err
		}
	}
	return nil
}

// mustInt ensures that a value is not zero.
func mustInt(key string, val int) error {
	if val == 0 {
		return fmt.Errorf("key %q has zero value", key)
	}
	return nil
}

func acceptImage(r *http.Request, imagePathPrefix string) (string, string, error) {
	if !(strings.HasPrefix(imagePathPrefix, "/") && strings.HasSuffix(imagePathPrefix, "/")) {
		return "", "", errors.New("imagePathPrefix must start and end with '/'")
	}
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("Image")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	imagePath := imagePathPrefix + handler.Filename

	f, err := os.OpenFile("."+imagePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	defer f.Close()
	io.Copy(f, file)
	return r.FormValue("Name"), imagePath, nil
}
