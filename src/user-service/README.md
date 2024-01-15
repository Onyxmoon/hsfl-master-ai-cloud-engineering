# The user service component
[![Run tests (user-service)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-user-service.yml/badge.svg)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-user-service.yml)

The user service component provides user data and authentication functions for registration and login as REST-Service. Additionally, it provides a service via gRPC to validate authorisation tokens for other services.

## API Endpoints

### Authentication

#### POST `/api/v1/authentication/login/`
- **Description:** Processes user login requests. Users need to provide login credentials.
- **Usage:** Submit authentication details (username, password) in request body. Returns a jwt token on success.

#### POST `/api/v1/authentication/register/`
- **Description:** Handles new user registrations.
- **Usage:** Submit registration details (username, password, email, etc.). Returns confirmation upon successful account creation.

### User Management

#### GET `/api/v1/user/role/:userRole`
- **Description:** Retrieves users based on their role.
- **Usage:** Include user role as URL parameter. Returns a list of users matching the specified role.
- **Access control:** Only merchants are unrestricted for users. An admin can fetch users by every role.

#### GET `/api/v1/user/:userId`
- **Description:** Fetches details of a specific user by their unique ID.
- **Usage:** Provide user ID as URL parameter. Returns details of the corresponding user.
- **Access control:** Only the user or an admin can do this.

#### PUT `/api/v1/user/:userId`
- **Description:** Allows updating a user's information.
- **Usage:** Include user ID as URL parameter and updated information in the request body. Updates user details in the database.
- **Access control:** Only the user or an admin can do this.

#### DELETE `/api/v1/user/:userId`
- **Description:** Deletes a user from the system.
- **Usage:** Require user ID as a URL parameter. Removes the user with the given ID from the database.
- **Access control:** Only the user or an admin can do this.

## gRPC Endpoints

### ValidateUserToken
- **Server Method:** `UserServiceServer.ValidateUserToken`
- **Request Type:** `ValidateUserTokenRequest`
- **Response Type:** `ValidateUserTokenResponse`

#### Description
This endpoint validates a user's authentication token. It checks the validity of the provided token and, if valid, retrieves the user's details based on the token's claims. If the verification fails due to an invalid token, it returns a status error.

#### Usage
This endpoint is primarily used for internal authentication and authorization purposes of other services, ensuring that the client's requests are accompanied by a valid, authenticated user token. It's essential in scenarios where user identity and permissions need to be validated before granting access to specific resources or operations.


## Configuration
You can configure this service with environmental variables or an environment file (.env) relative to the main application file.

### Example configuration
```dotenv
RQLITE_HOST="database"
RQLITE_PORT=4001
RQLITE_USER="db-user"
RQLITE_PASSWORD="db-pw-changeMe!"

HTTP_SERVER_PORT=3001

GRPC_SERVER_PORT=50051

JWT_PRIVATE_KEY="<path [e.g. privateKey.pem] or inline-key>"
```

To generate a ecdsa private key in pem format you can use the following command:
```shell
ssh-keygen -t ecdsa -f privateKey.pem -m pem
```
> When no or an invalid key is provided, container will use a random but secure generated key on each start. Note: Users will be logged-out when the service restarts.

## Run service
You can operate the service either as a standalone application or within a containerized environment. The database is required. It is designed to automatically generate the necessary tables should they not already exist.

Also: See the main [README.md](/README.md) to get information for a complete containerized setup. 