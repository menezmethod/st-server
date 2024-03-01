# st-server

`st-server` is a comprehensive server-side application designed for stock trading platforms. It's composed of three
microservices that handle various aspects of the trading platform's backend operations.

## Services Overview

- **Authorization/Users Service (`st-auth-svc`)**: Manages user authentication and authorization, ensuring secure access
  to the platform.
- **Journals/Trades Service (`st-journal-svc`)**: Handles recording and processing of trade activities, maintaining a
  reliable trade history.
- **Gateway Service (`st-gateway`)**: Provides a RESTful interface for client interactions, leveraging the `gin`
  framework for efficient HTTP request handling.

Each service is designed to operate independently, communicating with the others through gRPC for high-performance and
type-safe operations.

## Getting Started

Follow these instructions to set up the project on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.x or later)
- Docker (optional for containerization)
- Minikube (optional for local Kubernetes deployment)
- Skaffold (optional for continuous development with Kubernetes)

### Installation and Local Development

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/st-server.git
    ```

2. Navigate to the project directory:
    ```sh
    cd st-server
    ```

3. Install dependencies:
    ```sh
    make install_deps
    ```

4. Generate gRPC code for services:
    ```sh
    make proto_gw_auth proto_gw_journal proto_auth proto_journal
    ```

5. Build services:
    ```sh
    make build_gateway build_auth build_journal
    ```

6. Run services:
    ```sh
    make run_auth run_journal run_gateway
    ```

## Testing

To run automated tests for this system:

```sh
make test
```