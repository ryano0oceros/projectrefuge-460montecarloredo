# projectrefuge-460montecarloredo
Assignment 4. Monte Carlo Methods---Benchmark Design and Analysis

## Description

This project benchmarks the performance of various SQL queries executed against a PostgreSQL database using both Go and Python. The benchmarks are designed to measure the total and average execution times of each query when run multiple times (n=1000, n=10000, n=100000). The results provide insights into the efficiency and performance characteristics of the queries in different programming environments.

## Usage

### Prerequisites

- PostgreSQL database with the required schema and data.
- Go 1.23.2 or later.
- Python 3.8 or later.
- Required dependencies for Go listed in `go.mod`

### Running the Benchmarks

#### Go

1. Ensure the PostgreSQL database is running and accessible.
2. Update the connection string in `main.go` with the correct database credentials.
3. Install the required Go dependencies:
    ```sh
    go mod tidy
    ```
4. Run the Go benchmarks:
    ```sh
    go run main.go
    ```

#### Python

1. Ensure the PostgreSQL database is running and accessible.
2. Update the connection string in `python.py` with the correct database credentials.
3. Run the Python benchmarks:
    ```sh
    python3 python.py
    ```

The benchmark results will be printed to the console and can be reviewed for performance analysis.

## Go Benchmarks

### n=1000
- Query: list_invoices, Total Time: 0.31s, Average Time: 0.0003s
- Query: how_many_invoices, Total Time: 0.32s, Average Time: 0.0003s
- Query: how_many_customers, Total Time: 0.32s, Average Time: 0.0003s
- Query: vendors_in_tennessee, Total Time: 0.32s, Average Time: 0.0003s
- Query: products_with_discount, Total Time: 0.32s, Average Time: 0.0003s
- Query: list_customers, Total Time: 0.32s, Average Time: 0.0003s
- Query: most_expensive_product, Total Time: 0.32s, Average Time: 0.0003s
- Query: customers_without_purchases, Total Time: 0.34s, Average Time: 0.0003s
- Query: discounted_price_georgia, Total Time: 0.34s, Average Time: 0.0003s
- Query: vendor_products, Total Time: 0.35s, Average Time: 0.0004s
- Query: products_on_hand, Total Time: 0.37s, Average Time: 0.0004s
- Query: customer_purchases, Total Time: 0.38s, Average Time: 0.0004s

### n=10000
- Query: how_many_invoices, Total Time: 2.26s, Average Time: 0.0023s
- Query: vendors_in_tennessee, Total Time: 2.26s, Average Time: 0.0023s
- Query: how_many_customers, Total Time: 2.31s, Average Time: 0.0023s
- Query: list_invoices, Total Time: 2.32s, Average Time: 0.0023s
- Query: list_customers, Total Time: 2.33s, Average Time: 0.0023s
- Query: most_expensive_product, Total Time: 2.35s, Average Time: 0.0024s
- Query: products_with_discount, Total Time: 2.35s, Average Time: 0.0024s
- Query: customers_without_purchases, Total Time: 2.48s, Average Time: 0.0025s
- Query: discounted_price_georgia, Total Time: 2.52s, Average Time: 0.0025s
- Query: vendor_products, Total Time: 2.65s, Average Time: 0.0027s
- Query: customer_purchases, Total Time: 2.73s, Average Time: 0.0027s
- Query: products_on_hand, Total Time: 2.76s, Average Time: 0.0028s

### n=100000
- Query: list_invoices, Total Time: 20.65s, Average Time: 0.0207s
- Query: how_many_invoices, Total Time: 20.94s, Average Time: 0.0209s
- Query: how_many_customers, Total Time: 21.12s, Average Time: 0.0211s
- Query: vendors_in_tennessee, Total Time: 21.15s, Average Time: 0.0212s
- Query: products_with_discount, Total Time: 21.43s, Average Time: 0.0214s
- Query: list_customers, Total Time: 21.61s, Average Time: 0.0216s
- Query: most_expensive_product, Total Time: 21.69s, Average Time: 0.0217s
- Query: customers_without_purchases, Total Time: 22.68s, Average Time: 0.0227s
- Query: discounted_price_georgia, Total Time: 23.24s, Average Time: 0.0232s
- Query: vendor_products, Total Time: 24.55s, Average Time: 0.0246s
- Query: customer_purchases, Total Time: 25.68s, Average Time: 0.0257s
- Query: products_on_hand, Total Time: 25.89s, Average Time: 0.0259s

## Python Benchmarks

### n=1000
- Query: how_many_invoices, Total Time: 0.02s, Average Time: 0.0000s
- Query: list_invoices, Total Time: 0.02s, Average Time: 0.0000s
- Query: how_many_customers, Total Time: 0.02s, Average Time: 0.0000s
- Query: list_customers, Total Time: 0.03s, Average Time: 0.0000s
- Query: vendors_in_tennessee, Total Time: 0.02s, Average Time: 0.0000s
- Query: most_expensive_product, Total Time: 0.03s, Average Time: 0.0000s
- Query: products_with_discount, Total Time: 0.03s, Average Time: 0.0000s
- Query: vendor_products, Total Time: 0.07s, Average Time: 0.0001s
- Query: discounted_price_georgia, Total Time: 0.05s, Average Time: 0.0000s
- Query: products_on_hand, Total Time: 0.09s, Average Time: 0.0001s
- Query: customer_purchases, Total Time: 0.10s, Average Time: 0.0001s
- Query: customers_without_purchases, Total Time: 0.04s, Average Time: 0.0000s

### n=10000
- Query: how_many_invoices, Total Time: 0.22s, Average Time: 0.0002s
- Query: list_invoices, Total Time: 0.21s, Average Time: 0.0002s
- Query: how_many_customers, Total Time: 0.22s, Average Time: 0.0002s
- Query: list_customers, Total Time: 0.29s, Average Time: 0.0003s
- Query: vendors_in_tennessee, Total Time: 0.24s, Average Time: 0.0002s
- Query: most_expensive_product, Total Time: 0.33s, Average Time: 0.0003s
- Query: products_with_discount, Total Time: 0.26s, Average Time: 0.0003s
- Query: vendor_products, Total Time: 0.74s, Average Time: 0.0007s
- Query: discounted_price_georgia, Total Time: 0.49s, Average Time: 0.0005s
- Query: products_on_hand, Total Time: 0.93s, Average Time: 0.0009s
- Query: customer_purchases, Total Time: 0.96s, Average Time: 0.0010s
- Query: customers_without_purchases, Total Time: 0.42s, Average Time: 0.0004s

### n=100000
- Query: how_many_invoices, Total Time: 2.20s, Average Time: 0.0022s
- Query: list_invoices, Total Time: 2.15s, Average Time: 0.0022s
- Query: how_many_customers, Total Time: 2.34s, Average Time: 0.0023s
- Query: list_customers, Total Time: 2.66s, Average Time: 0.0027s
- Query: vendors_in_tennessee, Total Time: 2.21s, Average Time: 0.0022s
- Query: most_expensive_product, Total Time: 2.85s, Average Time: 0.0029s
- Query: products_with_discount, Total Time: 2.75s, Average Time: 0.0027s
- Query: vendor_products, Total Time: 6.78s, Average Time: 0.0068s
- Query: discounted_price_georgia, Total Time: 4.80s, Average Time: 0.0048s
- Query: products_on_hand, Total Time: 8.93s, Average Time: 0.0089s
- Query: customer_purchases, Total Time: 9.88s, Average Time: 0.0099s
- Query: customers_without_purchases, Total Time: 4.20s, Average Time: 0.0042s