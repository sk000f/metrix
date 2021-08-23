# Metrix

Calculate DORA metrics from your CI server:

- Deployment Frequency
- Lead Time
- Change Fail Rate
- Mean Time to Resolve (MTTR)

Codebase includes a sample Grafana dashboard, but the application exposes an API so data can be consumed by any external system of your choice.

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

## Command Line Flags

```
-update=y
```

Setting the update flag as above will cause the app to load updated data from GitLab on startup. Default value is `n` and default behaviour is to not load data on startup.

## Running Tests

Unit tests are implemented using the standard Go library so can be executed with the following command and flag:

```
go test ./... -run 'Unit'
```

Integration tests also use the standard library and can be executed with the following command and flag:

```
go test ./... -run 'Integration'
```

## Grafana

To run Grafana in a Docker container and configure the required plugin for the JSON API data source:

```
docker run -d -p 3000:3000 --name grafana -e "GF_INSTALL_PLUGINS=marcusolsson-json-datasource" grafana/grafana
```
