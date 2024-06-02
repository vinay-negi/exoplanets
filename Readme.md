# Exoplanet Microservice

This microservice allows users to manage and retrieve information about exoplanets.

## Endpoints

- `POST /exoplanets` - Add a new exoplanet.
- `GET /exoplanets` - List all exoplanets.
- `GET /exoplanets/{id}` - Get exoplanet details by ID.
- `PUT /exoplanets/{id}` - Update exoplanet details.
- `DELETE /exoplanets/{id}` - Delete an exoplanet.
- `POST /exoplanets/{id}/

## Fuel Estimation Formula

Fuel estimation to reach an exoplanet can be calculated as:

f = d / (g^2) * c units


where:
- `d` -> distance of exoplanet from Earth
- `g` -> gravity of exoplanet
- `c` -> crew capacity (int)

## Running the Service

1. Build the Docker image:
    ```
    docker build -t exoplanet-service .
    ```

2. Run the Docker container:
    ```
    docker run -p 8080:8080 exoplanet-service
    ```

3. The service will be available at `http://localhost:8080`.

## Dependencies

- Go 1.16
- Gorilla Mux
- Google UUID

## Error Handling

The service provides clear and consistent error messages for invalid requests.

## Extensibility

The service is designed to be easily extensible for new types of exoplanets.



# Exoplanet Microservice

This microservice manages exoplanets and provides fuel cost estimation for space voyages.

## Features

- Add an Exoplanet
- List Exoplanets
- Get Exoplanet by ID
- Update Exoplanet
- Delete Exoplanet
- Fuel Estimation for space voyages

## Requirements

- Go 1.18 or higher
- Docker (for containerization)

## Getting Started

### Running Locally

1. Clone the repository:

    ```sh
    git clone https://github.com/your-repo/exoplanet-service.git
    cd exoplanet-service
    ```

2. Run the application:

    ```sh
    go run cmd/main.go
    ```

3. Access the API at `http://localhost:8080`.

### Running with Docker

1. Build the Docker image:

    ```sh
    docker build -t exoplanet-service .
    ```

2. Run the Docker container:

    ```sh
    docker run -p 8080:8080 exoplanet-service
    ```

3. Access the API at `http://localhost:8080`.

## API Endpoints

- `POST /exoplanets`: Add a new exoplanet.
- `GET /exoplanets`: List all exoplanets.
- `GET /exoplanets/:id`: Get exoplanet by ID.
- `PUT /exoplanets/:id`: Update exoplanet.
- `DELETE /exoplanets/:id`: Delete exoplanet.
- `POST /exoplanets/:id/fuel-estimation`: Get fuel estimation for a trip to an exoplanet.

## Fuel Estimation Formula

Fuel estimation to reach an exoplanet can be calculated as:

