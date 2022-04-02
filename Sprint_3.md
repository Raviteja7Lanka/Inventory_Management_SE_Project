# Sprint 3 - Progress in  Sprint 3

FrontEnd Team: Ravi Teja Lanka, Nithin Meela 

BackEnd Team: Surya Mandapti, Sai Kiran Madipadige

## Frontend
- Developed the Reports Page as the dashboard.
- Developed services page with fully integration from the backend.
- Developed Orders page with cred funtionality.
- Developed The Products Page with including the Catagories.
- Fixed the CRED operations across the UI.
- Performed the Cypress testing for all the new components.
- Performed the Unit Testing For all the new components.

### Result of the Cypress Test

![cypresstest_sprint3](https://user-images.githubusercontent.com/94930984/161364471-a47856a2-237e-442d-8617-857bed01ca8d.jpeg)

## How to run Cypress testing
- ./node_modules/.bin/cypress open
- Click on Test1.js after opening cypress window
## Result of the Unit Test Cases
<img width="1439" alt="KarmaScreenshot" src="https://user-images.githubusercontent.com/94930984/161365933-b1b74263-e2fa-4d1f-90b1-b26280e48666.png">



## How to run Karma Unit Testing

- ng test

## Backend 
- Created Warehouses table and developed its API to accomodate the changes. 
- Developed GoLang API for Category table.
- Developed GoLang API for Order_Details table.
- Developed GoLang API for Role table (used to tell the hierachy of the existing staff.)
- Refactored the old API for better readability
- Validated the new API using Postman.
- In the Server folder, the APIs are placed in three folders
  <details>
        <summary>/Models</summary>
 
          - This folder contains the main implementation functions of ADD, UPDATE, DELETE, GET. 
 
   </details>
   <details>
        <summary>/Routes</summary>
 
          - This folder contains the routes of the API. 
 
   </details>
   <details>
        <summary>/Controllers</summary>
 
          - This functions in this folder will call actual implementation functions in the Models folder. 
 
   </details>

## Useful links of the project
- [Inventory Management Repo](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project) 
- [Developer Stories](https://github.com/Raviteja7Lanka/Inventory_Management_SE_Project/issues)


## UI Demo
[Link to video](https://youtu.be/yVaqCzsJV7w)

## Backend Demo
[Link to video](https://youtu.be/CR6iwblbIS4)



