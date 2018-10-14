package controller

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"fitsizecalculator/repository"
	"github.com/gorilla/mux"
)

func GetShoeList(w http.ResponseWriter, r *http.Request) {
	
	shoeslist, err := repository.GetAllShoes()
	
	if err != nil {
		panic(err)
	}

	for _, shoes := range shoeslist {
		fmt.Println(shoes.ProductId,shoes.ProductName)
	}

	b, err := json.Marshal(shoeslist)
	if err != nil {
		panic(err)
	}
	
	w.Header().Set("Contenty-Type", "application/json")
	w.Write(b)
}

func GetFitCalc (w http.ResponseWriter, r *http.Request) {
	params :=mux.Vars(r)
	fmt.Println("ProductId is:",params["productid"])
	productid, err := strconv.ParseInt(params["productid"],10,64)
	if err != nil {
		panic(err)
	}
	truefitcalc, err := repository.GetFitCalc(productid)
	
	if err != nil {
		panic(err)
	}
	
	b, err := json.Marshal(truefitcalc)
	
	if err != nil {
		panic(err)
	}
	
	w.Header().Set("Contenty-Type", "application/json")
	w.Write(b)
}

func SendFitVal (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	productid := params["productid"]
	fitval := params["fitval"]
	
	id,err := repository.SendFitVal(productid,fitval)
	if err != nil {
		panic(err)
	}
	
	type Message struct {
		rid int64 `json:"id" bson:"id"`
	}
	
	m := Message {id}
	b, err := json.Marshal(m)
	
	if err != nil {
		panic(err)
	}
	
	w.Header().Set("Contenty-Type", "application/json")
	w.Write(b)
	
}