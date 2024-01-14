# The http proxy service component
[![Run tests (http proxy service)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-http-proxy-service.yml/badge.svg)](https://github.com/Onyxmoon/hsfl-master-ai-cloud-engineering/actions/workflows/run-tests-http-proxy-service.yml)

The http proxy service component is the gateway component as between all services and serves as main entry point for clients. This service routes incoming HTTP requests to the appropriate internal services based on the predefined proxy routes.

## Configuration
The service is configured through a YAML file which details the proxy routes and additional settings. A typical configuration looks as follows:
```yaml
listenAddress: 0.0.0.0:8080
proxyRoutes:
  - name: User Service (User Info)
    context: /api/v1/user
    target: http://users:3001/api/v1/user
  - name: User Service (Login)
    context: /api/v1/authentication/login
    target: http://users:3001/api/v1/authentication/login
  - name: User Service (Registration)
    context: /api/v1/authentication/register
    target: http://users:3001/api/v1/authentication/register
  - name: Shoppinglist Service (Lists)
    context: /api/v1/shoppinglist
    target: http://shoppinglists:3002/api/v1/shoppinglist
  - name: Shoppinglist Service (Entries)
    context: /api/v1/shoppinglistentries
    target: http://shoppinglists:3002/api/v1/shoppinglistentries
  - name: Product Service (Products)
    context: /api/v1/product
    target: http://products:3003/api/v1/product
  - name: Product Service (Prices)
    context: /api/v1/price
    target: http://products:3003/api/v1/price
  - name: Web Service
    context: /
    target: http://web:3000
```

### Configuring the Path to Configuration File
You can specify the path to the configuration file using the `PROXY_CONFIG_PATH` environment variable. If this variable is set, the service will use the specified path to load the configuration file, when not it defaults to `./config/proxyConfig.yaml`. 

For example:
```shell
export PROXY_CONFIG_PATH=/path/to/your/config.yaml
```

### Explanation
- `listenAddress`: This defines the network address where the proxy service will listen for incoming requests.
- `proxyRoutes`: A list of routing definitions, each containing:
  - `name`: A descriptive name for the route.
  - `context`: The URL path that triggers this route.
  - `target`: The internal service URL to which the request should be forwarded.

## Run service
After setting up the configuration file and the `PROXY_CONFIG_PATH` environment variable, run the HTTP Proxy Service. It will start listening on the specified `listenAddress`. Incoming requests are then analyzed and routed according to the matching `context` found in the `proxyRoutes` settings.

Also: See the main [README.md](/README.md) to get information for a complete containerized setup.