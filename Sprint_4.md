# Sprint 4 - Summary of the Implemented Tasks

## Description of the application
Inventory management, also known as stock control, is the process of managing a company’s inventory levels in one or multiple locations or warehouses. This is developed using golang and angular. Basically, tracking items, orders from customers, orders placed to suppliers and making sure the business keeps the optimal number of parts or products in stock, whether it’s online, on the shelf, or both. An inventory control system also watches your products’ usage and storage, its stock room arrival up to customer shipment arrival.

It helps in the removal of slow-selling products and increasing the number of high-selling, so that a company caan maximize their profits, while saving on time, resources, and money with optimal inventory sitting in the warehouse.

The functionalities our application provides are as follows:

- **Signup:** For first-timers, the app provides a UI, takes in the first name, last name, user name, email, and password, and creates an account.
- **Login:**  Once the users are registered, they can log in using email and password. The details are validated using login API.
- **WareHouse:** Describes about the warehouses present in the database.
- **Staff:** Staff contains details of the authorized personnel as managers and workers and their login credentials.
- **Order:** 
  Creating an order, getting all the orders,in-progress, outstanding, and completed orders by a partcular customer id.
- **About Us:** 
  We have an About page which tell users about the web application.
  
  
## Group Members
SNO | Name                                 | Github username                         |           Role                           |
--- | -------------                        |:-------------:                          | :---------------------------------------:
1   | Saikiran Madipadige                  | SaiKiran Madipadige and smadipad        | Backend (Go lang) and Frontend (Angular) |
2   | Venkata Surya Narayana Raju Mandapati| Surya-Mandapati-19 and smadipad         | Backend (Go lang)                        |
3   | Ravi Teja Lanka                      | Raviteja7lanka                          | Frontend(Angular)                        |
4   | Nithin Meela                          | nithin-meela                           | Frontend(Angular)                         |


## UI Tasks achieved
- Added crud functionality for the warehouse. To create different warehouses with different locations.
- Add all apis for the report page.
- Implemented catalog page with different approach to see the products of all the warehouses.
- implemented full place order functionality in order page.
- Created cypress test for all  the components.
- Implemented view oders for the particular customer.
- Created test case for unit testing using for karma unit testing.


## Backend Tasks achieved

- Implemented the API's in model, controller, and router structure.
- UNIT tests written for all the API's
- Documented all the API's
- Developed GoLang REST API for the Staffs table to enable login authentication.


## Link to Project board: <br />
&nbsp;&nbsp;&nbsp;**Sprint 1:** https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/projects/1 <br />
&nbsp;&nbsp;&nbsp;**Sprint 2:** https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/projects/2 <br />
&nbsp;&nbsp;&nbsp;**Sprint 3:** https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/projects/3 <br />
&nbsp;&nbsp;&nbsp;**Sprint 4:** https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/projects/4 <br />

## Link to Sprint4 deliverables:  
&nbsp;https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/projects/4


## BONUS
- Deployed Frontend App Url: [https://raviteja7lanka.github.io/Inventory_Management_SE_Project/login](https://raviteja7lanka.github.io/Inventory_Management_SE_Project/login)  
- GitHub pages: (https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/deployments/activity_log?environment=github-pages)  

## Useful links of the project

- [Inventory Management Repo Link](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project)
- [Sprint 4 User stories progress board link](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/projects/4)
- [All user stories link](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/issues?q=is%3Aissue+)

## Sprint 4 Demo
-[Demo](https://youtu.be/Mr3VQCEiDh0)

## Backend GoLang unit tests demo
-[Backend](https://www.youtube.com/watch?v=lRBTuu2CCIk)

## Cypress automation tests Demo
-[cypress](https://www.youtube.com/watch?v=6PS4V9Guc10)

## Karma unit tests Demo
-[karma](https://www.youtube.com/watch?v=EWUQrRcEp4M)

### Result of the Unit Test Cases
<img width="1439" alt="KarmaScreenshot" src="https://user-images.githubusercontent.com/94930984/161365933-b1b74263-e2fa-4d1f-90b1-b26280e48666.png">

### How to run Karma Unit Testing

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).


## Api documentation of backend services

# Backend API
## Orders API

### 1. getAllCustomerOrders

```java
GET    http://localhost:8085/customer/orders/all
```
TO get all the existing records of the Orders table

#### *Example Request*
```java
GET    http://localhost:8085/customer/orders/all
```
```
#### *Example Response*

```java

{
        "order_id": 14,
        "date_of_order": "01/25/2022",
        "order_details_id": "",
        "customer_id": 1,
        "supplier_id": 0,
        "status": ""
    },
    {
        "order_id": 16,
        "date_of_order": "01/29/2022",
        "order_details_id": "",
        "customer_id": 2,
        "supplier_id": 0,
        "status": ""
    },
    {
        "order_id": 15,
        "date_of_order": "01/15/2022",
        "order_details_id": "",
        "customer_id": 1,
        "supplier_id": 0,
        "status": ""
    },
    {
        "order_id": 23,
        "date_of_order": "01/25/2022",
        "order_details_id": "",
        "customer_id": 1,
        "supplier_id": 0,
        "status": "o"
    },
    {
        "order_id": 230,
        "date_of_order": "02/25/2022",
        "order_details_id": "",
        "customer_id": 10,
        "supplier_id": 10,
        "status": "o"
    }
```


### 2. getCustomerOrderByID

```java
GET    http://localhost:8085/customer/order/{ordId}
```
To get a specific record from orders table which is identified by its id(orders_id)

#### *Example Request*

```java
GET    http://localhost:8085/customer/order/14
```

#### *Example Response*

```java
{"order_id":14,"date_of_order":"01/25/2022","order_details_id":"","customer_id":1,"supplier_id":0,"status":""}
```


### 3. addCustomerOrder

```java
POST    http://localhost:8085/customer/orders/add
```
TO add a record to the order table

#### *Example Request*

```java
POST    http://localhost:8085/customer/orders/add

Body

{
        "order_id": 10,
        "date_of_order": "01/25/2022",
        "order_details_id": "1",
        "customer_id": 1,
        "supplier_id": 0,
        "status": "1"
    }
```

#### *Example Response*

```java
Status 200 OK
```

### 4. updateCustomerOrder

```java
PUT    http://localhost:8085/customer/orders/{ordId}
```
TO modify an existing record in the order table

#### *Example Request*

```java
PUT    http://localhost:8085/customer/orders/{ordId}

{
        "order_id": 10,
        "date_of_order": "01/25/2022",
        "order_details_id": "changed by put",
        "customer_id": 1,
        "supplier_id": 0,
        "status": "1"
    }
```

#### *Example Response*

```java
Status: 200 OK

{"order_id":10,"date_of_order":"01/25/2022","order_details_id":"changed by put","customer_id":1,"supplier_id":0,"status":"1"}
```

### 5. deleteCustomerOrder

```java
DELETE    http://localhost:8085/customer/orders/{ordId}
```
TO delete an existing record in the order table

#### *Example Request*

```java
DELETE    http://localhost:8085/customer/orders/10
```

#### *Example Response*

```java
"{Status:200, Message: Deletion successful}"
``




## Customers API

### 1. getAllCustomer
```java
GET    http://localhost:8085//customer/all
```
TO get all the existing records of the Customers table

#### *Example Request*
```java
GET    http://localhost:8085//customer/all
```
```
#### *Example Response*

```java
{
        "customer_id": 1,
        "first_name": "SNbyput",
        "last_name": "MANDAPATI",
        "address": "GAINESVILLE",
        "phone": 8489415,
        "email": "VMANDAPATI@UFL.EDU",
        "staff_id": 1
    },
    {
        "customer_id": 2,
        "first_name": "SMADI",
        "last_name": "MADI",
        "address": "FLORIDA",
        "phone": 444855,
        "email": "SMADI@GMAIL.COM",
        "staff_id": 2
    },
    {
        "customer_id": 10,
        "first_name": "Added_by_post",
        "last_name": "MANDAPATI",
        "address": "GAINESVILLE",
        "phone": 8489415,
        "email": "VMANDAPATI@UFL.EDU",
        "staff_id": 1
    }

```


### 2. getCustomerByID

```java
GET    http://localhost:8085/customer/{cusId}
```
To get a specific record from Customers table which is identified by its id(customers_id)

#### *Example Request*

```java
GET    http://localhost:8085/customer/2
```

#### *Example Response*

```java
 {
        "customer_id": 2,
        "first_name": "SMADI",
        "last_name": "MADI",
        "address": "FLORIDA",
        "phone": 444855,
        "email": "SMADI@GMAIL.COM",
        "staff_id": 2
    }
```


### 3. addCustomer

```java
POST    http://localhost:8085/customer/add
```
TO add a record to the Customer table

#### *Example Request*

```java
POST    http://localhost:8085/customer/add

{
        "customer_id": 12,
        "first_name": "Post_SMADI",
        "last_name": "MADI",
        "address": "Gainesville",
        "phone": 444855,
        "email": "SMADI@GMAIL.COM",
        "staff_id": 2
    }
```

#### *Example Response*

```java
Status: 200 OK
```

### 4. updateCustomer

```java
PUT    http://localhost:8085/customer/{cusId}
```
TO modify an existing record in the customers table

#### *Example Request*

```java
PUT    http://localhost:8085/customer/12

{
        "customer_id": 12,
        "first_name": "Modified by Post",
        "last_name": "MADI",
        "address": "Gainesville",
        "phone": 444855,
        "email": "SMADI@GMAIL.COM",
        "staff_id": 2
    }
```

#### *Example Response*

```java
Status: 200 OK
{"customer_id":12,"first_name":"Modified by Post","last_name":"MADI","address":"Gainesville","phone":444855,"email":"SMADI@GMAIL.COM","staff_id":2}
```

### 5. deleteCustomer

```java
DELETE    http://localhost:8085/customer/{cusId}
```
TO delete an existing record in the Customers table

#### *Example Request*

```java
DELETE    http://localhost:8085/customer/12
```

#### *Example Response*

```java
"{Status:200, Message: Deletion successful}"
``



## Suppliers API

### 1. getAllSupplier
```java
GET    http://localhost:8085/supplier/all
```
TO get all the existing records of the Suppliers table

#### *Example Request*
```java
GET    http://localhost:8085/supplier/all
```

#### *Example Response*

```java

 {
        "supplier_id": 0,
        "name": "",
        "address": "",
        "phone": 0,
        "email": ""
    },
    {
        "supplier_id": 3,
        "name": "PUBLIX",
        "address": "488",
        "phone": 44,
        "email": "4445"
    },
    {
        "supplier_id": 113,
        "name": "PUBLIXby_post",
        "address": "488",
        "phone": 44,
        "email": "4445"
    },
    {
        "supplier_id": 1001,
        "name": "WALMART",
        "address": "BUTLER PLAZA",
        "phone": 1526,
        "email": "WALMART@GMAIL.COM"
    }

```


### 2. getSupplierByID

```java
GET    http://localhost:8085/supplier/{supId}
```
To get a specific record from Suppliers table which is identified by its id(suppliers_id)

#### *Example Request*

```java
GET    http://localhost:8085/supplier/1001

```

#### *Example Response*

```java
{
        "supplier_id": 1001,
        "name": "WALMART",
        "address": "BUTLER PLAZA",
        "phone": 1526,
        "email": "WALMART@GMAIL.COM"
    }
```


### 3. addSupplier

```java
POST    http://localhost:8085/supplier/add
```
TO add a record to the Supplier table

#### *Example Request*

```java
POST    http://localhost:8085/supplier/add
```
```java

{
        "supplier_id": 99,
        "name": "I want A",
        "address": "BUTLER PLAZA",
        "phone": 1526,
        "email": "WALMART@GMAIL.COM"
    }
```

#### *Example Response*

```java
Status 200 OK

{"supplier_id":99,"name":"I want A","address":"BUTLER PLAZA","phone":1526,"email":"WALMART@GMAIL.COM"}

```

### 4. updateSupplier

```java
PUT    http://localhost:8085/supplier/{supId}
```
TO modify an existing record in the suppliers table

#### *Example Request*

```java
PUT    http://localhost:8085/supplier/99
{
        "supplier_id": 99,
        "name": "I want A_ modified by put",
        "address": "BUTLER PLAZA",
        "phone": 1526,
        "email": "WALMART@GMAIL.COM"
    }

```

#### *Example Response*

```java
Status 200 OK
{
        "supplier_id": 99,
        "name": "I want A_ modified by put",
        "address": "BUTLER PLAZA",
        "phone": 1526,
        "email": "WALMART@GMAIL.COM"
    }
```

### 5. deleteSupplier

```java
DELETE    http://localhost:8085/customer/{cusId}
```
TO delete an existing record in the Suppliers table

#### *Example Request*

```java
DELETE    http://localhost:8085/customer/113
```

#### *Example Response*

```java
"{Status:200, Message: Deletion successful}"
``

- Similarly the Payments, Products have that functionality.
