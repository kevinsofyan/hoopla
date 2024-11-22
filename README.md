# Hoopla

## Description
Hoopla is a project that provides an engaging and interactive experience for users. It includes various features and functionalities to enhance user interaction, such as product management, order processing, and reporting.

## Installation
To install the project, follow these steps:

1. Clone the repository:
   ```sh
   git clone <repository-url>
   ```

2. Navigate to the project directory:
   ```sh
   cd hoopla
   ```

3. Install the dependencies:
   ```sh
   go mod tidy
   ```

## Usage
To run the project, execute the following command:
```sh
go run main.go
```

## Features

### User Roles
- **Admin**: Can access reporting, update item stock, and manage categories.
- **Customer**: Can buy items, view orders, and process payments.

### Menus

#### Main Menu
- **Admin**:
  - Reporting
  - Update Item Stock
  - Exit
- **Customer**:
  - Buy Item
  - Show Orders
  - Process Payment
  - Exit

#### Reporting Menu
- Total Sales Report
- Most Popular Product Report
- Total Revenue Per Product Report
- Customer Count Per City Report
- Back to Main Menu

#### Update Item Stock Menu
- List All Products
- Create New Product
- Update Product Stock
- Delete a Product
- Return to Main Menu

#### Buy Item Menu
- Displays all products available for purchase
- Allows customers to buy items by entering product details

#### Payment Menu
- Displays unpaid orders
- Allows customers to process payments for orders

## Database Configuration
The database connection is configured in `config/config.go`. Update the connection string as needed:

```go
DB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hoopla_db")
```

## Components

### Handlers
- **Product Handler**: Manages product-related operations such as showing, creating, updating, and deleting products. Implemented in `handler/handler_product.go`.
- **Payment Handler**: Manages order and payment-related operations. Implemented in `handler/payment_handler.go`.
- **General Handler**: Manages general operations such as reporting, user management, and category management. Implemented in `handler/handler.go`.

### CLI
The command-line interface (CLI) is implemented in `cli/cli.go` and other files in the `cli` directory. It provides an interactive menu for users to perform various operations.