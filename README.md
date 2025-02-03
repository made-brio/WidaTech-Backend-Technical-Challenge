# WidaTech Backend Technical Challenge

## Overview
This repository contains the solution for WidaTech's Backend Engineer Technical Challenge, focusing on software integration for IoT applications. The challenge is divided into three main sections:

1. **Invoice CRUD API**  
2. **CSV/XLSX Import and Parsing**  
3. **Problem-Solving Algorithm**  

---

## Table of Contents
- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Setup Instructions](#setup-instructions)
- [API Documentation](#api-documentation)
  - [Invoice CRUD API](#invoice-crud-api)
  - [CSV/XLSX Import API](#csvxlsx-import-api)
- [Problem-Solving Algorithm](#problem-solving-algorithm)
- [License](#license)

---

## Tech Stack
- **Backend:** Golang  
- **Database:** MySQL / PostgreSQL  
- **API Type:** RESTful API  

---

## Setup Instructions

1. **Clone the Repository:**  
   ```bash
   git clone https://github.com/your-username/widatech-backend-challenge.git
   cd widatech-backend-challenge
   ```

2. **Configure Environment Variables:**  
   Create a `.env` file and add your database configurations:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database
   ```

3. **Run the Application:**  
   ```bash
   go run main.go
   ```

---

## API Documentation

### Invoice CRUD API

1. **Create Invoice**  
   - **Endpoint:** `POST /api/invoices`  
   - **Request Body:**
     ```json
     {
       "invoice_no": "INV001",
       "date": "2023-01-01",
       "customer_name": "John Doe",
       "salesperson_name": "Jane Smith",
       "payment_type": "CASH",
       "notes": "First transaction",
       "products": [
         {
           "item_name": "Product A",
           "quantity": 2,
           "total_cost_of_goods_sold": 50,
           "total_price_sold": 70
         }
       ]
     }
     ```

2. **Read Invoices**  
   - **Endpoint:** `GET /api/invoices`  
   - **Query Parameters:** `date`, `size`, `page`

3. **Update Invoice**  
   - **Endpoint:** `PUT /api/invoices/:invoice_no`

4. **Delete Invoice**  
   - **Endpoint:** `DELETE /api/invoices/:invoice_no`

---

### CSV/XLSX Import API

- **Endpoint:** `POST /api/import`
- **Description:** Upload an XLSX file with two sheets: `invoice` and `product_sold`. The API validates the data and saves valid entries while returning errors for faulty records.

- **Example Response for Errors:**
  ```json
  {
    "errors": [
      {
        "invoice_no": "INV002",
        "error": "Missing required field: customer_name"
      }
    ]
  }
  ```

---

## Problem-Solving Algorithm

This function generates all possible unique combinations of non-repeating digits (1-9) based on the provided length `l` and total sum `t`.

- **Example Usage:**
  ```go
  combinations := findCombinations(3, 8)
  fmt.Println(combinations) // Output: [[1,2,5], [1,3,4]]
  ```

---

## License
This project is licensed under the MIT License.

