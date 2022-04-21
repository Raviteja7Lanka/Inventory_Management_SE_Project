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

  it("Checks whether the Reports Card is Leading to the Reports Page or not ", function () {
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
  // it("Checks whether the Delete Order Button is present and bringing the dialog", function () {
  //   cy.visit("http://localhost:4200/orders");
  //   cy.get("#delete-order");
  // });
  // it("Checks whether the Edit Order Button is present and bringing the dialog", function () {
  //   cy.visit("http://localhost:4200/orders");
  //   cy.get("#edit-order");
  // });
});

describe("Test For Going Back to Login Page upon Clicking logout button", function () {
  it("Checks whether the logout button is present in the toggle or not", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#menutoggleright").click();
    cy.contains("Logout");
  });
  it("Checks whether the profile button is present in the toggle or not", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#menutoggleright").click();
    cy.contains("profile");
    cy.get("#profile").click();
  });
  it("checks whether the logout button is leading to sigin page or not", function () {
    cy.visit("http://localhost:4200/home");
    cy.get("#menutoggleright").click();
    cy.get("#logout").click();
    cy.contains("Login");
  });
});

describe("Test for Login Page", function () {
  it("Checks whether the Create new account is leading to the signup page or not", function () {
    cy.visit("http://localhost:4200/login");
    cy.get("#createnewaccount").click();
    cy.contains("Sign");
  });
  it("Checks whether the Forgot password is leading to the forgot password ", function () {
    cy.visit("http://localhost:4200/login");
    cy.get("#forgotpassword").click();
    cy.contains("Forgot");
  });
});

describe("Test for Signup Page", function () {
  it("Checks whether the already registered is leading back to login or not", function () {
    cy.visit("http://localhost:4200/signup");
    cy.contains("Forgot");
    cy.get("#forgotpassword").click();
    cy.contains("Forgot");
  });
  it("Checks whether the already registered is leading back to login", function () {
    cy.visit("http://localhost:4200/signup");
    cy.contains("Already");
    cy.get("#alreadyregistered").click();
    cy.contains("Login");
  });
});
