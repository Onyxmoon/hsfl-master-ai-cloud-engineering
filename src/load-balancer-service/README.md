# Load Balancer Service

This directory contains a Load Balancer designed to orchestrate Docker containers and distribute network load. It functions as a proxy, offering selectable algorithms to optimize the distribution of requests among containers.



## Algorithms
The load balancer supports several algorithms to optimize request distribution. Configure your choice of algorithm in the `config.yml` file. Available algorithms include:

- **Round Robin** (`round_robin`): Distributes requests evenly across all containers. This is a good default choice for general use.
- **Least Connections** (`least_connections`): Directs traffic to the container with the fewest active connections. Ideal for situations where session persistence is important.
- **Least Response Time** (`least_response_time`): Routes requests to the container with the lowest average response time. This is beneficial when you need to optimize for speed and performance.


## Configuration
To configure the load balancer, edit the `config.yml` file in the root of the directory. Here's an example configuration:

```yaml
image: onyxmoon/pw-web-service:latest
network: internal
replicas: 5
algorithm: round_robin
```
In this example:
- Use the Docker image onyxmoon/pw-web-service:latest for the load balanced containers. 
- Connect the containers to specified network in docker. Please make sure that the service can access the network for health checking.
- Create 5 replicas of the container for load balancing.
- The selected algorithm is Round Robin