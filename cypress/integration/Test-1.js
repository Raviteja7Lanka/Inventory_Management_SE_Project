describe("Test for Suppliers Table Component", function () {
  it("Checks whether the add-supplier button is opening a dialog with add supplier form", function () {
    cy.visit("http://localhost:4200/supplier-table");
    cy.get("#add_supplier").click();
    cy.contains("Please Enter the Supplier Details");
    cy.get("#addSupplierFormCancel").click();
  });

  it("Checks whether the delete supplier button is opening a dialog with delete supplier form", function () {
    cy.visit("http://localhost:4200/supplier-table");
    cy.contains("Delete Supplier");
    cy.get("#delete_supplier").click();
    cy.contains("Email Address");
    cy.get("#cancelDeleteSupplierForm").click();
  });
});
describe("Test for Customer Table Component", function () {
  it("Checks whether the add customer button is opening a dialog with add customer form", function () {
    cy.visit("http://localhost:4200/customer-table");
    cy.get("#add_customer").click();
    cy.contains("Please Enter the Customer Details");
    cy.get("#addCustomerFormCancel").click();
  });

  it("Checks whether the delete customer button is opening a dialog with delete customer form", function () {
    cy.visit("http://localhost:4200/customer-table");
    cy.contains("Delete Customer");
    cy.get("#delete_customer").click();
    cy.contains("Email Address");
    cy.get("#cancelDeleteCustomerForm").click();
  });
});

describe("Test for Cards in the Home Page", function () {
  it("Checks whether the products card is leading to the products page or not", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#Products").click();
    cy.contains("Category 1");
  });

  it("Checks whether the Services Card is Leading to the Warehouse Page or not ", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#Services").click();
    cy.contains("Warehouse");
  });

  it("Checks whether the Reports Card is Leading to the Warehouse Page or not ", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#Reports").click();
    cy.contains("layout");
  });

  it("Checks whether the Orders Card is Leading to the Orders Page or not ", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#Orders").click();
    cy.contains("Order Id");
  });

  it("Checks whether the menu button in the home page is opening the side toggle or not", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#Menu_Button").click();
    cy.contains("About");
  });
});

describe("Test For Orders Page", function () {
  it("Checks whether the Add Order Button is present and bringing the dialog", function () {
    cy.visit("http://localhost:4200/orders");
    cy.get("#add-order");
  });
  it("Checks whether the Delete Order Button is present and bringing the dialog", function () {
    cy.visit("http://localhost:4200/orders");
    cy.get("#delete-order");
  });
  it("Checks whether the Edit Order Button is present and bringing the dialog", function () {
    cy.visit("http://localhost:4200/orders");
    cy.get("#edit-order");
  });
});
