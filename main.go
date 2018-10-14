package main

import (
	"fmt"
	"log"
	
	"net/http"
	"fitsizecalculator/controller"
	
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the StockX True Fit Calculator")
	router:= mux.NewRouter()
	router.HandleFunc("/shoelist",controller.GetShoeList).Methods("GET")
	router.HandleFunc("/getfitcalc/{productid}",controller.GetFitCalc).Methods("GET")
	router.HandleFunc("/sendfitval/{productid}/{fitval}",controller.SendFitVal).Methods("POST")
	
	staticFileDirectory := http.Dir("./view/")
	staticFileHandler := http.StripPrefix("/view/",http.FileServer(staticFileDirectory))
	router.PathPrefix("/view/").Handler(staticFileHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",router))	
   
}