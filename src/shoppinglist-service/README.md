# The data management service component
[![Run tests (shoppinglist-service)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-shoppinglist-service.yml/badge.svg)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-shoppinglist-service.yml)

The shopping list service component allows clients to manage entities (modified CRUD) like the users shopping lists and their entries. 

This service has access restrictions and requires a connection via gRPC to the [user service](/src/user-service) in order to authorize transmitted JWT tokens.

## API Endpoints

### Shopping List Management

#### GET `/api/v1/shoppinglist/:userId`
- **Description:** Retrieves all shopping lists associated with a specific user.
- **Usage:** Requires the user's ID as a URL parameter. Returns a collection of shopping lists for the user.
- **Access control:** Users can only fetch their own shopping lists.

#### GET `/api/v1/shoppinglist/:listId/:userId`
- **Description:** Fetches a specific shopping list for a user.
- **Usage:** Include both the list ID and user ID as URL parameters. Returns the details of the specified shopping list.
- **Access control:** Users can only fetch their own shopping lists.

#### PUT `/api/v1/shoppinglist/:listId/:userId`
- **Description:** Updates an existing shopping list.
- **Usage:** Requires the list ID and user ID as URL parameters, and the updated list details in the request body.
- **Access control:** Users can only update their own shopping lists.

#### POST `/api/v1/shoppinglist/:userId`
- **Description:** Creates a new shopping list for a user.
- **Usage:** Requires the user's ID as a URL parameter and the list details in the request body.
- **Access control:** Users can only add shopping lists for themselves.

#### DELETE `/api/v1/shoppinglist/:listId`
- **Description:** Deletes a specific shopping list.
- **Usage:** Requires the list ID as a URL parameter.
- **Access control:** Users can only delete their own shopping lists.

### Shopping List Entry Management

#### GET `/api/v1/shoppinglistentries/:listId`
- **Description:** Retrieves all entries in a specific shopping list.
- **Usage:** Requires the list ID as a URL parameter.
- **Access control:** Users can only fetch entries from their own shopping lists.

#### GET `/api/v1/shoppinglistentries/:listId/:productId`
- **Description:** Fetches a specific entry in a shopping list.
- **Usage:** Include both the list ID and product ID as URL parameters.
- **Access control:** Users can only fetch entries from their own shopping lists.

#### PUT `/api/v1/shoppinglistentries/:listId/:productId`
- **Description:** Updates an existing entry in a shopping list.
- **Usage:** Requires the list ID and product ID as URL parameters, and the updated entry details in the request body.
- **Access control:** Users can only update entries to their own shopping lists.

#### POST `/api/v1/shoppinglistentries/:listId/:productId`
- **Description:** Adds a new entry to a shopping list.
- **Usage:** Requires the list ID and product ID as URL parameters, and the entry details in the request body.
- **Access control:** Users can only add entries from their own shopping lists.

#### DELETE `/api/v1/shoppinglistentries/:listId/:productId`
- **Description:** Removes an entry from a shopping list.
- **Usage:** Requires both the list ID and product ID as URL parameters.
- **Access control:** Users can only delete entries from their own shopping lists.


## Configuration
You can configure this service with environmental variables or an environment file (.env). relative to the main application file.

### Example configuration
```dotenv
RQLITE_HOST="database"
RQLITE_PORT=4001
RQLITE_USER="db-user" # Change this!
RQLITE_PASSWORD="db-pw-changeMe!" # Change this!

HTTP_SERVER_PORT=3002

GRPC_USER_SERVICE_TARGET="users:50051"
```

## Run service
You can operate the service either as a standalone application or within a containerized environment. The database is required. It is designed to automatically generate the necessary tables should they not already exist.

Also: See the main [README.md](/README.md) to get information for a complete containerized setup. 