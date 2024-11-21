package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection details
	connStr := "postgres://postgres:secret@localhost/saleco?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Queries to execute
	queries := map[string]string{
		"how_many_invoices": "SELECT COUNT(*) AS total_invoices FROM INVOICE;",
		"list_invoices":     "SELECT INV_NUMBER, INV_DATE FROM INVOICE;",
		"how_many_customers": `
            SELECT COUNT(*) AS total_customers FROM customer;
        `,
		"list_customers": `
            SELECT
                cus_code,
                CONCAT(cus_fname, ' ', COALESCE(cus_initial, ''), ' ', cus_lname) AS customer_name
            FROM customer;
        `,
		"vendors_in_tennessee": `
            SELECT v_code, v_name
            FROM vendor
            WHERE v_state = 'TN';
        `,
		"most_expensive_product": `
            SELECT p_code, p_descript, p_price, p_qoh
            FROM product
            ORDER BY p_price DESC
            LIMIT 1;
        `,
		"products_with_discount": `
            SELECT p_descript, p_qoh, p_price
            FROM product
            WHERE p_discount > 0.05;
        `,
		"vendor_products": `
            SELECT
                v.v_name AS vendor_name,
                p.p_code AS product_code,
                p.p_descript AS product_name
            FROM vendor v
            JOIN product p ON v.v_code = p.v_code
            ORDER BY v.v_name, p.p_code;
        `,
		"discounted_price_georgia": `
            SELECT
                v.v_code AS vendor_code,
                p.p_code AS product_code,
                p.p_descript AS product_description,
                ROUND(p.p_price * (1 - p.p_discount), 2) AS discounted_price
            FROM vendor v
            JOIN product p ON v.v_code = p.v_code
            WHERE v.v_state = 'GA';
        `,
		"products_on_hand": `
            SELECT
                v.v_name AS vendor_name,
                p.p_code AS product_code,
                p.p_descript AS product_description,
                COALESCE(p.p_qoh, 0) AS products_on_hand
            FROM vendor v
            LEFT JOIN product p ON v.v_code = p.v_code
            ORDER BY v.v_name, p.p_code;
        `,
		"customer_purchases": `
            SELECT
                c.CUS_CODE,
                i.INV_NUMBER,
                l.LINE_NUMBER,
                l.P_CODE,
                l.LINE_UNITS,
                l.LINE_PRICE,
                (l.LINE_UNITS * l.LINE_PRICE) AS total_price
            FROM CUSTOMER c
            JOIN INVOICE i ON c.CUS_CODE = i.CUS_CODE
            JOIN LINE l ON i.INV_NUMBER = l.INV_NUMBER
            ORDER BY c.CUS_CODE, i.INV_NUMBER, l.LINE_NUMBER;
        `,
		"customers_without_purchases": `
            SELECT
                c.CUS_CODE,
                c.CUS_LNAME,
                c.CUS_FNAME
            FROM CUSTOMER c
            LEFT JOIN INVOICE i ON c.CUS_CODE = i.CUS_CODE
            WHERE i.INV_NUMBER IS NULL
            ORDER BY c.CUS_CODE;
        `,
	}

	var wg sync.WaitGroup
	for name, query := range queries {
		wg.Add(1)
		go func(name, query string) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < 1000; i++ {
				rows, err := db.Query(query)
				if err != nil {
					log.Fatalf("Error executing query %s: %v", name, err)
				}
				rows.Close()
			}
			duration := time.Since(start)
			fmt.Printf("Query: %s, Total Time: %.2fs, Average Time: %.4fs\n", name, duration.Seconds(), duration.Seconds()/1000)
		}(name, query)
	}
	wg.Wait()
}
