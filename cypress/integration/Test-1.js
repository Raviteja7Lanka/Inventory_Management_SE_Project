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
