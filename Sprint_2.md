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
[Link to video](https://youtu.be/1reyv-ckDNU)

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
Body:
```java
{
    "username":"",
    "password":"123456",
    "addressLine1":"",
    "addressLine2":"1",
    "phone":"",
    "email":".com"
}
```
#### *Example Response*

```java

{
    
}

```


### 2. getCustomerOrderByID

```java
GET    http://localhost:8085/customer/order/{ordId}
```
To get a specific record from orders table which is identified by its id(orders_id)

#### *Example Request*

```java
GET    http://localhost:8085/customer/order/1
```

#### *Example Response*

```java
{
    "code": 200,
    
}
```


### 3. addCustomerOrder

```java
POST    http://localhost:8085/customer/orders/add
```
TO add a record to the order table

#### *Example Request*

```java
POST    http://localhost:8085/customer/orders/add
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

### 4. updateCustomerOrder

```java
PUT    http://localhost:8085/customer/orders/{ordId}
```
TO modify an existing record in the order table

#### *Example Request*

```java
PUT    http://localhost:8085/customer/orders/{ordId}
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

### 5. deleteCustomerOrder

```java
DELETE    http://localhost:8085/customer/orders/{ordId}
```
TO delete an existing record in the order table

#### *Example Request*

```java
DELETE    http://localhost:8085/customer/orders/{ordId}
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
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
Body:
```java
{
    "username":"",
    "password":"123456",
    "addressLine1":"",
    "addressLine2":"1",
    "phone":"",
    "email":".com"
}
```
#### *Example Response*

```java

{
    
}

```


### 2. getCustomerByID

```java
GET    http://localhost:8085/customer/{cusId}
```
To get a specific record from Customers table which is identified by its id(customers_id)

#### *Example Request*

```java
GET    http://localhost:8085/customer/{cusId}
```

#### *Example Response*

```java
{
    "code": 200,
    
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
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

### 4. updateCustomer

```java
PUT    http://localhost:8085/customer/{cusId}
```
TO modify an existing record in the customers table

#### *Example Request*

```java
PUT    http://localhost:8085/customer/{cusId}
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

### 5. deleteCustomer

```java
DELETE    http://localhost:8085/customer/{cusId}
```
TO delete an existing record in the Customers table

#### *Example Request*

```java
DELETE    http://localhost:8085/customer/{cusId}
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
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
Body:
```java
{
    "username":"",
    "password":"123456",
    "addressLine1":"",
    "addressLine2":"1",
    "phone":"",
    "email":".com"
}
```
#### *Example Response*

```java

{
    
}

```


### 2. getSupplierByID

```java
GET    http://localhost:8085/supplier/{supId}
```
To get a specific record from Suppliers table which is identified by its id(suppliers_id)

#### *Example Request*

```java
GET    http://localhost:8085/supplier/{supId}
```

#### *Example Response*

```java
{
    "code": 200,
    
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

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

### 4. updateSupplier

```java
PUT    http://localhost:8085/supplier/{supId}
```
TO modify an existing record in the suppliers table

#### *Example Request*

```java
PUT    http://localhost:8085/supplier/{supId}
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

### 5. deleteSupplier

```java
DELETE    http://localhost:8085/customer/{cusId}
```
TO delete an existing record in the Suppliers table

#### *Example Request*

```java
DELETE    http://localhost:8085/customer/{cusId}
```

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
``
























## Rating API

### 1. Rating Create API

```java
POST    http://localhost:1016/rating/create
```
Create a new rating with the provided information

#### *Example Request*
```java
POST    http://localhost:1016/rating/create
```
Body:
```java
{
    "username": "ZhongkaiSun",
    "restaurantName": "Popeyes",
    "star": 5,
    "comment": "Taste so good !!!!",
    "ratingDate": "2022/01/06"
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 2. Rating GET API

//


## Cuisine API

### 1. Cuisine Create API

```java
POST    http://localhost:1016/cuisine/create
```
Create a new cuisine for a restaurant

#### *Example Request*
```java
POST    http://localhost:1016/cuisine/create
```
Request Body:
```java
{
"name":"banana pie",        
"restaurantName":"Popeyes", 
"price":2.0,        
"calories":3
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 2. Cuisine Read API


```java
GET    localhost:1016/cuisine/read?restaurantName=Popeyes
```
Read a cuisine from a restaurant

#### *Example Request*
```java
GET    http://localhost:1016/cuisine/create
```

#### *Example Response*

```java

{
    "code": 200,
    "data": [
        {
            "name": "8PC Nuggets A La Carte",
            "restaurantName": "Popeyes",
            "price": 4.79,
            "calories": 0
        },
        {
            "name": "apple pie",
            "restaurantName": "Popeyes",
            "price": 2,
            "calories": 3
        },
        {
            "name": "banana pie",
            "restaurantName": "Popeyes",
            "price": 2,
            "calories": 3
        },
        {
            "name": "Mixed Chicken",
            "restaurantName": "Popeyes",
            "price": 22.79,
            "calories": 0
        },
        {
            "name": "Red Beans and Rice",
            "restaurantName": "Popeyes",
            "price": 4.19,
            "calories": 0
        }
    ],
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 3. Cuisine Delete API

```java
POST    http://localhost:1016/cuisine/delete
```
Delete the cuisine from a restaurant

#### *Example Request*
```java
POST    http://localhost:1016/cuisine/delete
```
Request Body:
```java
{
"name":"banana pie",        
"restaurantName":"Popeyes", 
"price":2.0,        
"calories":3
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

## Order API

### 1. Order Create API

```java
POST    http://localhost:1016/order/create
```
Create a new order for a customer

#### *Example Request*
```java
POST    http://localhost:1016/order/create
```
Request Body:
```java
{
    "userName":"Raindrop",
	"restaurantName":"Checkers",
	"orderDate":"03/02/2022",
	"price":8.88,
	"cuisineName":"burger"
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error
  
### 2. Order Read API

```java
GET    http://localhost:1016/order/read?username=Raindrop
```
Create a new cuisine for a restaurant

#### *Example Request*
```java
GET    http://localhost:1016/order/read?username=Raindrop
```

#### *Example Response*

```java

{
    "code": 200,
    "data": [
        {
            "username": "Raindrop",
            "restaurantName": "Checkers",
            "orderDate": "03/02/2022",
            "price": 8.88,
            "cuisineName": "burger"
        }
    ],
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error
- 
## Restaurant API

### 1. Restaurant Read API
```java
GET    http://localhost:1016/restaurant/read
```
Read a restaurant or get the restaurant list 

#### *Example Request*
```java
GET    http://localhost:1016/restaurant/read?name=Popeyes
```
#### *Example Response*

```java

{
    "code": 200,
    "data": {
        "name": "Popeyes",
        "address": "1412N Main St, Gainesville, FL 32601, USA",
        "deliveryFee": 3.99,
        "imgPath": "",
        "typeofMeal": "Chicken",
        "rating": 4.1
    },
    "msg": "Successfully"
}

```

#### *Example Request*
```java
GET    http://localhost:1016/restaurant/read
```
#### *Example Response*

```java

{
    "code": 200,
    "data": [
        {
            "name": "Popeyes",
            "address": "1412N Main St, Gainesville, FL 32601, USA",
            "deliveryFee": 3.99,
            "imgPath": "",
            "typeofMeal": "Chicken",
            "rating": 4.1
        },
        {
            "name": "Checkers",
            "address": "3325 W University Ave, Gainesville, FL 32607, USA",
            "deliveryFee": 1.99,
            "imgPath": "",
            "typeofMeal": "Burgers",
            "rating": 4.2
        }
    ],
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

