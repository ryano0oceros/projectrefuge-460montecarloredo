import psycopg2
import time

# Database connection details
connection_string = "dbname=saleco user=postgres password=secret"

# Connect to the database
conn = psycopg2.connect(connection_string)
cursor = conn.cursor()

# Queries to execute
queries = {
    "how_many_invoices": "SELECT COUNT(*) AS total_invoices FROM INVOICE;",
    "list_invoices": "SELECT INV_NUMBER, INV_DATE FROM INVOICE;",
    "how_many_customers": "SELECT COUNT(*) AS total_customers FROM customer;",
    "list_customers": """
        SELECT
            cus_code,
            CONCAT(cus_fname, ' ', COALESCE(cus_initial, ''), ' ', cus_lname) AS customer_name
        FROM customer;
    """,
    "vendors_in_tennessee": """
        SELECT v_code, v_name
        FROM vendor
        WHERE v_state = 'TN';
    """,
    "most_expensive_product": """
        SELECT p_code, p_descript, p_price, p_qoh
        FROM product
        ORDER BY p_price DESC
        LIMIT 1;
    """,
    "products_with_discount": """
        SELECT p_descript, p_qoh, p_price
        FROM product
        WHERE p_discount > 0.05;
    """,
    "vendor_products": """
        SELECT
            v.v_name AS vendor_name,
            p.p_code AS product_code,
            p.p_descript AS product_name
        FROM vendor v
        JOIN product p ON v.v_code = p.v_code
        ORDER BY v.v_name, p.p_code;
    """,
    "discounted_price_georgia": """
        SELECT
            v.v_code AS vendor_code,
            p.p_code AS product_code,
            p.p_descript AS product_description,
            ROUND(p.p_price * (1 - p.p_discount), 2) AS discounted_price
        FROM vendor v
        JOIN product p ON v.v_code = p.v_code
        WHERE v.v_state = 'GA';
    """,
    "products_on_hand": """
        SELECT
            v.v_name AS vendor_name,
            p.p_code AS product_code,
            p.p_descript AS product_description,
            COALESCE(p.p_qoh, 0) AS products_on_hand
        FROM vendor v
        LEFT JOIN product p ON v.v_code = p.v_code
        ORDER BY v.v_name, p.p_code;
    """,
    "customer_purchases": """
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
    """,
    "customers_without_purchases": """
        SELECT
            c.CUS_CODE,
            c.CUS_LNAME,
            c.CUS_FNAME
        FROM CUSTOMER c
        LEFT JOIN INVOICE i ON c.CUS_CODE = i.CUS_CODE
        WHERE i.INV_NUMBER IS NULL
        ORDER BY c.CUS_CODE;
    """
}

# Execute each query 1000 times
for query_name, query in queries.items():
    start_time = time.time()
    for _ in range(100000):
        cursor.execute(query)
        cursor.fetchall()
    duration = time.time() - start_time
    print(f"Query: {query_name}, Total Time: {duration:.2f}s, Average Time: {duration / 1000:.4f}s")

# Cleanup
cursor.close()
conn.close()
