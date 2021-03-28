# Metrix

Calculate DORA metrics from CI server

## Structure

Metrix uses a hexagonal architecture (ports and adapters) pattern and is structured as follows.

### `/core/app`

Core application logic. You shouldn't need to change any of this code when swapping in and out adapters.

### `/core/domain`

Core domain model for deployments. When building a new adapter this is the model you need to align with.

### `/core/ports`

Each of the ports exposed from the core:

- **ciserver** - represents a CI server where the data about deployments will be queried from
- **repository** - represents the database for storing deployment and incident data
- **service** - represents the API that can be consumed by the Grafana dashboard or any other external tool

### `/adapters`

Each of the currently implemented adapters for CI Server, Database and API. If you use a different CI server or database please feel free to build and add a new adapter here.

### `/internal`

Some common resources used throughout the application such as global constants and enums.

### `/cmd`

Contains the `main()` function and is where you can plug in the different adapters for your use case.

## Running Tests

Unit tests are implemented using the standard Go library so can be executed with the following command:

```
go test ./...
```
