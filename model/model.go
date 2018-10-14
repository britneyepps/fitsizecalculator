package model

import (
)

type Shoes struct {
	ProductId int	`json:"productid" bson:"productid"`
	ProductName string `json:"productname" bson:"productname"`
}

func PairOfShoes( productid int, productname string) Shoes{
	shoes := new(Shoes)
	shoes.ProductId = productid
	shoes.ProductName = productname
	
	return *shoes
}	

type TrueFitCalc struct {
	ProductId int64 `json:"productid" bson:"productid"`
	FitCalcValue float64 `json:"fitcalcvalue" bson:"fitcalcvalue"`
}

func TrueFit( productid int64, fitcalcval float64) TrueFitCalc{
	truefitcalc := new(TrueFitCalc)
	truefitcalc.ProductId = productid
	truefitcalc.FitCalcValue = fitcalcval
	
	return *truefitcalc
}