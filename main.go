package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func buy(cash, price float64) float64 {
	return cash / price
}

func sell(shares, price float64) float64 {
	return shares * price
}

func main() {

	cash := 1000.0
	shares := 0.0
	buyPrice := 1000000.0 // =infinity
	sellPrice := 0.0

	buyTrigger := 0.996
	sellTrigger := 1.02 // buy or sell if there is a 2% change

	file, err := os.Open("ftse.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		priceString := strings.ReplaceAll(record[2], ",", "")
		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			log.Print(err)
			continue //skip any bad values
		}
		if cash > 0 && price <= buyPrice {
			shares = buy(cash, price)
			cash = 0
			sellPrice = price * sellTrigger
			fmt.Println("Buy ", record[1], price, shares*price)
			continue
		}
		if shares > 0 && price >= sellPrice {
			cash = sell(shares, price)
			shares = 0
			buyPrice = price * buyTrigger
			fmt.Println("Sell", record[1], price, cash)
			continue
		}
	}
}
