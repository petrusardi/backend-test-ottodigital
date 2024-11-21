# Backend Application

This backend application provides API endpoints for managing vouchers, redemptions, transactions, and brands. The application is built with Go and uses PostgreSQL as the database.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Setting Up the Project](#setting-up-the-project)
   1. [Clone the Repository](#1-clone-the-repository)
   2. [Set Up the Database](#2-set-up-the-database)
   3. [Configure Environment Variables](#3-configure-environment-variables)
3. [Install Dependencies](#install-dependencies)
4. [Running the Application](#running-the-application)
5. [Testing the Application](#testing-the-application)
6. [Accessing API Endpoints](#accessing-api-endpoints)
7. [Troubleshooting](#troubleshooting)
8. [License](#license)

## Prerequisites

Before you can run the application, ensure you have the following installed on your local machine:

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) (version 12 or higher)
- [Git](https://git-scm.com/) (optional, for cloning the repository)
- [Make](https://www.gnu.org/software/make/) (optional, for automating some tasks)

## Setting Up the Project

### 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/petrusardi/backend-test-ottodigital.git
cd backend-test
```

### 2. Set Up the Database

1. **Create a PostgreSQL Database:**

   If you don't already have PostgreSQL installed, download and install it. After installation, create a new database for this application.

   ```bash
   psql -U postgres
   CREATE DATABASE your_database_name;
   ```

2. **Set Up Database Schema:**

   Apply the migrations to set up your tables by running the following command:

   ```bash
   go run main.go migrate
   ```

### 3. Configure Environment Variables

Set up your environment variables, particularly for the database connection. You can set this up in a `.env` file or manually set the environment variables.

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=your_database_name
```

### 4. Install Dependencies

Run the following command to install the necessary Go dependencies:

```bash
go mod tidy
```

## Running the Application

To start the application locally, run the following command:

```bash
go run main.go
```

By default, the application will start on `localhost:8080`. You can access the API endpoints through this URL.

## Accessing API Endpoints

Once the application is running, you can access the API through the following endpoints:

### Create Brand

- **POST** `/brand`
- **Request body example:**
  ```json
  {
    "name": "Brand A"
  }
  ```

### Create Voucher

- **POST** `/voucher`
- **Request body example:**
  ```json
  {
    "name": "Voucher Name",
    "cost": 100000,
    "brand_id": 2,
    "code": "INDO-001",
    "description": "top-up 100k",
    "value": 100000
  }
  ```

### Get Single Voucher

- **GET** `/voucher?id={voucher_id}`
- Example request:
  ```bash
  curl "http://localhost:8080/voucher?id=1"
  ```

### Get All Vouchers by Brand

- **GET** `/voucher/brand?id={brand_id}`
- Example request:
  ```bash
  curl "http://localhost:8080/voucher/brand?id=1"
  ```

### Create Transaction

- **POST** `/transaction`
- **Request body example:**
  ```json
  {
    "customer_name": "Budi"
  }
  ```

### Create Redemption

- **POST** `/transaction/redemption`
- **Request body example:**
  ```json
  {
    "transaction_id": 2,
    "voucher_id": 10
  }
  ```

### Redemption Detail

- **GET** `/transaction/redemption?transactionId={transactionId}`
- Example request:
  ```bash
  curl "http://localhost:8080/transaction/redemption?transactionId=1"
  ```

## Troubleshooting

- **Error: Connection to database failed**

  - Ensure your PostgreSQL database is running and that the connection details in your `.env` file are correct.

- **Error: Port already in use**

  - Check if another application is using port `8080` and stop it, or change the port number in your application code.

- **Error: Migration not applied correctly**
  - Ensure that you ran the migration commands correctly. You may need to check if your database schema is up to date by running the migration command again.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
