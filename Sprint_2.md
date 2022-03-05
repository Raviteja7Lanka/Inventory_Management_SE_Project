# Sprint 2 - Progress in  Sprint 2

## Frontend
- Developed Login Page with authentication service
- Devloped Sign Up Page to register the new user  
- Developed the proposed functionalities in the Supplier and Customer Components 
- Performed the Cypress Testing for the components 
- Perfomed the Unit Test Cases on the base components present in the project 

### Result of the Cypress Test
![image](https://user-images.githubusercontent.com/78529687/156867691-9f705afb-ff72-4dbc-9c76-f8134d34f922.png)

## Result of the Unit Test Cases
![image](https://user-images.githubusercontent.com/78529687/156867805-f34db55b-00b9-4573-abae-0b2eb8988717.png)



## Backend 
- Developed GoLang API for Customers table.
- Developed GoLang API for Products table.
- Developed GoLang API for Suppliers table.
- Developed GoLang API for Orders table.
- Developed GoLang REST API for the Staffs table to enable login authentication.
- Go lang Unit tests for User rest apis.
- Documented the new rest API
- Modified the database to accommodate new features.

## Useful links of the project
- [Inventory Management Repo](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project) 
- [Developer Stories](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/issues)


## UI Demo
[Link to video](https://www.youtube.com/watch?v=2KdnLYBD05U)

## Backend Demo
[Link to video](https://youtu.be/7t3zbaAwJ1s)



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

