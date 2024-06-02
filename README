# Exoplanets Microservice

This microservice allows users to manage and retrieve information about exoplanets and provides fuel cost estimation for space voyages.

## Architecture
The microservice architecture is designed with consideration of Domain-Driven Design (DDD) patterns and SOLID principles.

## Features

- Add an Exoplanet
- List Exoplanets
- Get Exoplanet by ID
- Update Exoplanet
- Delete Exoplanet
- Fuel Estimation for space voyages

## Endpoints

- `POST /exoplanets` - Add a new exoplanet.
- `GET /exoplanets?sortBy=<mass/radius>&order=<asc>` - List all exoplanets. Could be sorted by mass or radius, and order in ascending
- `GET /exoplanets/{id}` - Get exoplanet details by ID.
- `PUT /exoplanets/{id}` - Update exoplanet details.
- `DELETE /exoplanets/{id}` - Delete an exoplanet.
- `GET /exoplanets/{id}/fuel?crewCapacity=<int>` - Get fuel estimation of an exoplanet according to Crewcapacity


## Fuel Estimation Formula

Fuel estimation to reach an exoplanet can be calculated as:

f = d / (g^2) * c units


where:
- `d` -> distance of exoplanet from Earth
- `g` -> gravity of exoplanet
- `c` -> crew capacity (int)

## Running the Service in container

1. Build the Docker image:
    ```
    docker build -t exoplanet-service .
    ```
2. Run the Docker container:
    ```
    docker run -p 8080:8080 exoplanet-service
    ```

3. The service will be available at `http://localhost:8080`.


## Running the Service locally

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


## Dependencies

- Go 1.21
- Gorilla Mux
- Google UUID

## Error Handling

The service provides clear and consistent error messages for invalid requests.

## Extensibility

The service is designed to be easily extensible for new types of exoplanets.

## Requirements

- Go 1.21 or higher
- Docker (for containerization)
