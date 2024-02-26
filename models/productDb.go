package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpass = ""
const dbname = "go_crud"

func GetProducts() []Product {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	products := []Product{}

	for results.Next() {

		var prod Product

		err = results.Scan(&prod.ID, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, prod)

		//fmt.Println("product.id :", prod.id+" : "+prod.Name)
	}

	return products

}

func GetProduct(id int) *Product {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	prod := &Product{}

	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product where id=?", id)

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil
	}

	if results.Next() {

		err = results.Scan(&prod.ID, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return prod

}

func AddProduct(product Product) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO product (name,qty,last_updated) VALUES (?,?, now())",
		product.Name, product.Qty)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
