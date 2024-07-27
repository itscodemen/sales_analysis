# Sales Analysis API

A RESTful API for analyzing historical sales data using Go, Gin framework, and MySQL.

## Prerequisites

- **Go 1.16+**
- **MySQL**

## Setup

1. **Clone the Repository:**

    ```bash
    git clone <repository-url>
    cd sales-analysis
    ```

2. **Set Up Environment Variables:**

    Create a `.env` file with:

    ```env
    DB_USER=<your_db_user>
    DB_PASSWORD=<your_db_password>
    DB_NAME=<your_db_name>
    DB_HOST=<your_db_host>
    DB_PORT=<your_db_port>
    ```

3. **Install Dependencies:**

    ```bash
    go mod tidy
    ```

4. **Run the Application:**

    ```bash
    go run main.go
    ```

    The server will start on port `8080`.

## API Endpoints

### Get Total Revenue

- **URL:** `/revenue`
- **Method:** `GET`
- **Query Parameters:**
  - `start_date`: `YYYY-MM-DD`
  - `end_date`: `YYYY-MM-DD`

### Sample Request

```http
GET /revenue?start_date=2023-01-01&end_date=2023-12-31
