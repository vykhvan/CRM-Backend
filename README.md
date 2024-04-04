# CRM-Backend
Golang Project CRM Backend

## Description

This project represents the back end of a CRM app. As users interact with the app via some user interface, the server that you'll build will support the following functionalities:

* Getting a list of all customers
* Getting data for a single customer
* Adding a customer
* Updating a customer's information
* Removing a customer


# How to Run

1. Make sure you have Go installed. If not, install it from the official Go website: https://golang.org/

2. Clone the repository:

    ```
    git clone https://github.com/vykhvan/CRM-Backend.git
    ```

3. Navigate to the project directory:

    ```
    cd CRM-Backend
    ```

4. Run the application using the command:

    ```
    go run main.go
    ```

5. Open your browser and go to the following URL:

    ```
    http://localhost:3000
    ```

# Dependencies

This project depends on the `github.com/gorilla/mux` package for implementing HTTP request routing.

# Consume API

## Getting a list of all customers

```
curl --location 'localhost:3000/customers'
```

## Getting data for a single customer

```
curl --location 'localhost:3000/customers/1'
```

## Adding a customer

```
curl --location 'localhost:3000/customers' \
--header 'Content-Type: application/json' \
--data-raw '{
        "ID": "7",
        "Name": "Vyacheslav Khvan",
        "Role": "Data Scientist",
        "Email": "alice.johnson@example.com",
        "Phone": 5550403,
        "Contacted": true
}'
```

## Updating a customer's information

```
curl --location --request PUT 'localhost:3000/customers/8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ID": "8",
    "Name": "Vyacheslav Khvan",
    "Role": "Software Engineer",
    "Email": "ric@example.com",
    "Phone": 5550100,
    "Contacted": true
}'
```

## Removing a customer
```
curl --location --request DELETE 'localhost:3000/customers/8'
```