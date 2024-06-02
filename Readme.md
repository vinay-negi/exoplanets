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
