# The product service component
[![Run tests (product-service)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-product-service.yml/badge.svg)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-product-service.yml)

The product service component allows to manage entities (modified CRUD) like the products and their prices. For future extensions, this service also makes the products available internally to other services via gRPC.

This service has access restrictions and requires a connection via gRPC to the [user service](/src/user-service) in order to authorize transmitted JWT tokens.

## API Endpoints

### Product Management

#### GET `/api/v1/product/`
- **Description:** Retrieves a list of all products.
- **Usage:** Returns an array of all products in the system.
- **Access control:**: Not restricted

#### GET `/api/v1/product/:productId`
- **Description:** Retrieves details of a product by its unique ID.
- **Usage:** Include the product ID as a URL parameter. Returns the product's details.
- **Access control:**: Not restricted

#### GET `/api/v1/product/ean/:productEan`
- **Description:** Fetches a specific product using its EAN (European Article Number).
- **Usage:** Requires the product's EAN as a URL parameter. Returns details of the product associated with the given EAN.
- **Access control:**: Not restricted

#### PUT `/api/v1/product/:productId`
- **Description:** Updates details of an existing product.
- **Usage:** Requires the product ID as a URL parameter and the updated product information in the request body.
- **Access control:**: Only merchants or admins can do this.

#### POST `/api/v1/product/`
- **Description:** Adds a new product to the system.
- **Usage:** Include product details in the request body.
- **Access control:**: Only merchants or admins can do this.

#### DELETE `/api/v1/product/:productId`
- **Description:** Deletes a specific product from the system.
- **Usage:** Requires the product ID as a URL parameter.
- **Access control:**: Only admins can do this.

### Price Management

#### GET `/api/v1/price/`
- **Description:** Retrieves pricing information for all products.
- **Usage:** Returns a list of prices for all products.
- **Access control:**: Not restricted

#### GET `/api/v1/price/user/:userId`
- **Description:** Fetches prices specifically tailored for a merchant user.
- **Usage:** Include the user ID as a URL parameter. Returns user-specific pricing details.
- **Access control:**: Not restricted

#### GET `/api/v1/price/product/:productId`
- **Description:** Retrieves all prices associated with a specific product.
- **Usage:** Requires the product ID as a URL parameter.
- **Access control:**: Not restricted

#### GET `/api/v1/price/:productId/:userId`
- **Description:** Fetches the price of a product for a specific merchant user.
- **Usage:** Include both the product ID and user ID as URL parameters.
- **Access control:**: Not restricted

#### PUT `/api/v1/price/:productId/:userId`
- **Description:** Updates the price of a product for a specific merchant user.
- **Usage:** Requires the product ID and user ID as URL parameters, and the updated price details in the request body.
- **Access control:**: Merchants can only update their own prices for a product.

#### POST `/api/v1/price/:productId/:userId`
- **Description:** Adds a new price for a product and merchant user combination.
- **Usage:** Include the product ID, user ID, and price details in the request body.
- **Access control:**: Merchants can only add their own prices for a product.

#### DELETE `/api/v1/price/:productId/:userId`
- **Description:** Removes a specific price entry for a product and merchant user.
- **Usage:** Requires both the product ID and user ID as URL parameters.
- **Access control:**: Merchants can only delete their own prices for a product.


## Configuration
You can configure this service with environmental variables or an environment file (.env). relative to the main application file.

### Example configuration
```dotenv
RQLITE_HOST="database"
RQLITE_PORT=4001
RQLITE_USER="db-user"
RQLITE_PASSWORD="db-pw-changeMe!"

HTTP_SERVER_PORT=3003

GRPC_SERVER_PORT=50053
GRPC_USER_SERVICE_TARGET="users:50051"
```

## Run service
You can operate the service either as a standalone application or within a containerized environment. The database is required. It is designed to automatically generate the necessary tables should they not already exist.

Also: See the main [README.md](/README.md) to get information for a complete containerized setup. 