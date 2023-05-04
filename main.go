package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
)

func main(){
	file, err := os.Open("content.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range rows {
		fmt.Println(v)
	}
}