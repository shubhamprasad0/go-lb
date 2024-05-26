# go-lb: Load Balancer Project in Go

## Overview

This project implements a simple load balancer in Go. It allows distributing incoming requests to multiple HTTP application servers based on a specified configuration. The load balancer also includes health check functionality to ensure that traffic is only directed to healthy servers.

## Features

- Configurable through command-line arguments and YAML configuration files.
- Supports health checks to monitor the status of application servers.
- Customizable buffer size for reading request data.
- Easily extensible to add more features.

## Installation

### Prerequisites

- Go 1.22 or higher
- Make

### Steps

1. Clone the repository:
   ```sh
   git clone git@github.com:shubhamprasad0/go-lb.git
   cd go-lb
   ```

2. Build the application using the provided Makefile:
   ```sh
   make
   ```

## Usage

### Command-Line Arguments

You can configure the load balancer using command-line arguments:

- `-config`: Path to the YAML configuration file.
- `-port`: Port on which the load balancer runs.
- `-buffer-size`: Size of buffer (in bytes) used while reading request data.
- `-health-check-route`: Health check endpoint on the application servers.
- `-health-check-interval`: Number of seconds after which health check is performed periodically.
- `-servers`: Comma-separated list of application server addresses.

Example:
```sh
./bin/lb -config=config.yaml -port=8080 -buffer-size=1024 -health-check-route=/health -health-check-interval=30 -servers=server1:8080,server2:8080
```

### YAML Configuration File

You can also configure the load balancer using a YAML file. The configuration from the YAML file can be overridden by command-line arguments.

Example `config.yaml`:
```yaml
port: 8080
bufferSize: 1024
healthCheckRoute: "/health"
healthCheckInterval: 30
servers:
  - "server1:8080"
  - "server2:8080"
```

### Running the Load Balancer

To start the load balancer, use the following command:

```sh
./bin/lb -config=config.yaml
```

## Development

### Directory Structure

- `cmd/`: Contains the main package for the load balancer.
- `pkg/`: Contains the package with the load balancer logic.
- `bin/`: Output directory for the compiled binary.

### Building the Project

Use the Makefile to build the project:

```sh
make build
```

### Cleaning Up

To clean up the build artifacts:

```sh
make clean
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for details. 