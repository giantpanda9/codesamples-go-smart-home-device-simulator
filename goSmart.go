package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"time"
)

type Smartdev struct {
    Turnedon bool `json:"turnedon"`
    Temp int `json:"temp,omitempty"`
}

var acTemp, heaterTemp int = 1,0
var heaterSwitch, lightSwitch, acSwitch, comingBackSoon bool = false, false, false, false

func getHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(map[string]bool{"comingHome": comingBackSoon})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getReady(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getPostParams(r *http.Request) (string,string,string) {
	if err := r.ParseForm(); err != nil {
		log.Println("ParseForm() err: %v", err)
	}
	signal := r.Form.Get("signal")
	comfTemp := r.Form.Get("comfTemp")
	goingHome := r.FormValue("goingHome")
	return signal,comfTemp,goingHome
}

func smartSwitch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		smartDevice := Smartdev{Turnedon: lightSwitch,Temp: 0}
		json.NewEncoder(w).Encode(smartDevice)
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		signal,_,_ := getPostParams(r)
		if signal == "hot" {
			lightSwitch = false
		} else {
			lightSwitch = true
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func smartAC(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		smartDevice := Smartdev{Turnedon: acSwitch,Temp: acTemp}
		json.NewEncoder(w).Encode(smartDevice)		
	} else if r.Method == "POST" {
		signal,comfTemp,_ := getPostParams(r)
		comfTempInt, err := strconv.Atoi(comfTemp)
		if err != nil {
			fmt.Println(err)
		}
		if signal == "hot" {
			acSwitch = true
			if acTemp > 10 { acTemp = acTemp - 10 }
		} else if signal == "cold" {
			acSwitch = true
			if acTemp < comfTempInt { acTemp = acTemp + 13 }
		} else {
			acSwitch = false
			acTemp = comfTempInt
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func smartHeater(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		smartDevice := Smartdev{Turnedon: heaterSwitch,Temp: heaterTemp}
		json.NewEncoder(w).Encode(smartDevice)
	} else if r.Method == "POST" {
		signal,_,goingHome:= getPostParams(r)
		goingHomeBool, err := strconv.ParseBool(goingHome)
		if err != nil {
			fmt.Println(err)
		}
		if (goingHomeBool) {
			comingBackSoon = goingHomeBool
		}
		if signal == "cold" && comingBackSoon {
			if (heaterSwitch == false) {
				heaterSwitch = true
				time.AfterFunc(25 * time.Minute, func() {
					heaterSwitch = false
				})
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", getHome).Methods("GET")
	muxRouter.HandleFunc("/smartdevice", getHome).Methods("GET")
	muxRouter.HandleFunc("/smartdevice/ping", getReady).Methods("GET")
	muxRouter.HandleFunc("/smartdevice/switch", smartSwitch).Methods("GET")
	muxRouter.HandleFunc("/smartdevice/ac", smartAC).Methods("GET")
	muxRouter.HandleFunc("/smartdevice/heater", smartHeater).Methods("GET")
	muxRouter.HandleFunc("/smartdevice/switch", smartSwitch).Methods("POST")
	muxRouter.HandleFunc("/smartdevice/ac", smartAC).Methods("POST")
	muxRouter.HandleFunc("/smartdevice/heater", smartHeater).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", muxRouter))
}
