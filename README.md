# st-server

`st-server` is a comprehensive server-side solution written in Go, tailored for stock trading platforms. It encapsulates three distinct services that communicate via gRPC, offering a cohesive backend system for handling authorization, journaling trades, and interfacing with RESTful requests.

## Services Overview

- **Authorization/Users Service**: Manages user authentication and authorization, ensuring secure access to the platform.
- **Journals/Trades Service**: Records and processes all trading activities, maintaining a secure and reliable trade history.
- **Gateway Service**: Serves as the RESTful endpoint for the platform, utilizing `gin` for efficient HTTP request routing.

Each service communicates internally using gRPC to ensure high performance and efficient type-safe messaging.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.x or later)
- gRPC and Protocol Buffers

### Installing

A step-by-step series of examples that tell you how to get a development environment running:

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/st-server.git
    ```

2. Navigate to the project directory:
    ```bash
    cd st-server
    ```

3. Install dependencies:
    ```bash
    go mod download
    ```

4. Start each service separately. For example, to start the Authorization/Users Service:
    ```bash
    go run services/auth/main.go
    ```

Repeat the steps above for the Journals/Trades Service and the Gateway Service.

## Usage

A brief description of how to use the system, possibly with a few quick examples.

## Running the tests

Explain how to run the automated tests for this system.

    ```bash
    go test ./...
    ```

## Deployment

Add additional notes about how to deploy this on a live system, potentially with links to more comprehensive documentation or guides.

## Authors

- Luis Gimenez

See also the list of [contributors](https://github.com/yourusername/st-server/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.