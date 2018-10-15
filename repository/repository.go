package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/britneyepps/fitsizecalculator/model"
	"github.com/britneyepps/fitsizecalculator/database"
)

func GetAllShoes() ([]model.Shoes, error) {
	DB, err := database.NewOpen()
	
	rows, err := DB.Query("Select productid, productname from stockx.shoes")
	if err != nil {
		panic(err)
	}
	
	defer rows.Close()
	
	shoeslist := make([]model.Shoes,0)
	
	for rows.Next() {
		var productId int
		var productName string
		shoepair := model.Shoes{}
		err := rows.Scan(&productId,&productName)
		
		if err != nil {
			print(err)
		}
		
		shoepair.ProductId = productId
		shoepair.ProductName = productName
		fmt.Println(productId,productName)
		
		shoeslist = append(shoeslist,shoepair)
	}
	
	if err = rows.Err(); err != nil {
        return nil, err
    }

    DB.Close()

    return shoeslist, nil
}

func GetFitCalc(productid int64) (model.TrueFitCalc, error) {
	fmt.Println("In Get Calc Method")
	DB, err := database.NewOpen()
	var fitcalculation float64
	
	err = DB.QueryRow("select avg(fitvalue) from stockx.truesize where productid = $1",productid).Scan(&fitcalculation)
	if err != nil {
		panic(err)
	}
	truefitcalc := model.TrueFitCalc{}
	truefitcalc.ProductId = productid
	truefitcalc.FitCalcValue = fitcalculation

	DB.Close()
	
	return truefitcalc, nil
}

func SendFitVal(productid string, fitval string) (int64, error) {
	DB, err := database.NewOpen()
	var id int64
	
	err = DB.QueryRow("INSERT INTO stockx.truesize (productid,fitvalue) VALUES ($1,$2) RETURNING id",productid,fitval).Scan(&id)
	if err != nil {
			panic(err)
	}
	
	DB.Close()
	fmt.Println("New id is",id)
	return id, nil
}