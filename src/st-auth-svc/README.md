# Authentication Service

## Description

The Authentication Service is a core component of the stock trading platform, responsible for managing user
authentication and authorization. It ensures secure access to the platform's features and services, leveraging advanced
security protocols to protect sensitive data.

## Installation

To set up the Authentication Service for development or testing, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/menezmethod/st-server.git
    ```

2. Navigate to the Authentication Service directory:

    ```bash
    cd st-server/src/st-auth-svc
    ```

3. Install the necessary dependencies and generate gRPC code:

    ```bash
    make proto
    ```

This command will generate all the necessary gRPC code required for the service to run.

## Running the App

To start the Authentication Service in a development environment, use the following command:

```bash
make server
```