# HTTP Stress Test

## Overview
This HTTP stress test tool is designed to simulate high traffic and analyze the performance of web servers under different load conditions. It allows for configuration to customize the test scenarios.

## Configuration
The tool supports configuration in YAML, providing a simple and readable format. Below are the key components of the YAML configuration:

```yaml
phases:
    - targetrps: 0
      timeidx: 0
    - targetrps: 500
      timeidx: 10
    - targetrps: 500
      timeidx: 20
targets:
    - url: https://google.de:443
users: 10
```

### Phases
The phases section defines the different segments of the stress test. Each phase specifies the target number of requests per second (RPS) and the time index at which this RPS should be reached.

- `targetrps`: The target number of requests per second for this phase.
- `timeidx`: The time index in seconds when the specified RPS should be reached, starting from 0.

### Targets
The `targets` section lists the URLs or endpoints to be tested. The tool will distribute the generated traffic across these targets.
- `url`: The specific URL or endpoint to be included in the stress test

### Users
The users field indicates the number of simulated users or concurrent workers generating the traffic.

## Additional Information
- Requests in config are per second and not an amount
- If you're not waiting for a response, the metric still prints responses, but the response time you see is only the time it takes to send one and all the requests will be considered successful whether they get a response or not
- You can choose whether you want the test to wait for a response (using fasthttp) or not wait for a response (using a TCP client)
- If you wait for responses and the server crashes or does not respond anymore, the test might not exit because it's still waiting for a response