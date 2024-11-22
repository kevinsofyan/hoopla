## ERD Title: Hoopla Database System

## Entities and Their Attributes:

A. Entity: User
- Attributes:
  - UserID (PK, AI): INT
  - Role: VARCHAR(50)
  - Name: VARCHAR(50)
  - Email: VARCHAR(100)
  - Phone: VARCHAR(20)
  - Address: VARCHAR(255)
  - City: VARCHAR(50)

B. Entity: Category
- Attributes:
  - CategoryID (PK, AI): INT
  - CategoryName: VARCHAR(50)

C. Entity: Product
- Attributes:
  - ProductID (PK, AI): INT
  - ProductName: VARCHAR(100)
  - CategoryID (FK): INT
  - Price: DECIMAL(10,2)
  - StockQuantity: INT
  - Description: TEXT

D. Entity: Orders
- Attributes:
  - OrderID (PK, AI): INT
  - UserID (FK): INT
  - OrderDate: DATETIME
  - Paid: BOOLEAN
  - TotalAmount: DECIMAL(10,2)

E. Entity: OrderDetails
- Attributes:
  - OrderDetailID (PK, AI): INT
  - OrderID (FK): INT
  - ProductID (FK): INT
  - Quantity: INT
  - UnitPrice: DECIMAL(10,2)

F. Entity: Payment
- Attributes:
  - PaymentID (PK, AI): INT
  - OrderID (FK): INT
  - PaymentDate: DATETIME
  - PaymentMethod: VARCHAR(50)
  - PaymentAmount: DECIMAL(10,2)

## Relationships:
- User to Orders: One to Many
  - Description: One user can place multiple orders, but each order belongs to only one user.

- Category to Product: One to Many
  - Description: One category can have multiple products, but each product belongs to only one category.

- Product to OrderDetails: One to Many
  - Description: One product can appear in multiple order details, but each order detail line refers to only one product.

- Orders to OrderDetails: One to Many
  - Description: One order can have multiple order details, but each order detail belongs to only one order.

- Orders to Payment: One to Many
  - Description: One order can have multiple payments, but each payment belongs to only one order.

## Integrity Constraints:
- All prices and amounts should be positive decimals
- Stock quantity should be non-negative
- User email should be unique
- Product name and category name are required fields
- Order date and payment date cannot be null
- Payment amount should match the order total amount

## Additional Notes:
- The Orders table has a Paid boolean flag to track payment status
- OrderDetails table maintains the unit price at the time of purchase
- Foreign key constraints ensure referential integrity between all related tables
- The system supports multiple payments for a single order