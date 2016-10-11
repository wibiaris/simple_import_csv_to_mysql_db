package main

import (
	"fmt"
	"flag"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/csv"
	"os"
)

func main() {
	csvPath := flag.String("path", "test", "Path to an CSV file")
	table := flag.String("table", "sales_order_return", "Table to insert data")
	flag.Parse()
	fmt.Println(*csvPath)

	csvFile, err := os.Open(*csvPath)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("mysql", "username:password@tcp(mysqlhost:port)/database_name?charset=utf8")
	checkErr(err)

	for n, each := range csvData {
		if n == 0 {
			continue;
		}
		soStoreNumber := each[0]
		storeId := each[1]
		salesOrderId := each[2]
		fgPaid := 1
		query := fmt.Sprintf("INSERT INTO %s (store_id, sales_order_id, so_store_number, fg_paid) VALUES(?,?,?,?)", *table)
		fmt.Println(soStoreNumber)
		_, err = db.Exec(query, storeId, salesOrderId, soStoreNumber, fgPaid)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	
}

func checkErr(err error) {
	if err !=nil {
		panic(err)
	}
}