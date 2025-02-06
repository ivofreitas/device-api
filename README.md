# Device API

[![Go Version](https://img.shields.io/badge/Go-1.23.5-blue.svg)](https://golang.org/)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/ivofreitas/device-api/actions)
[![License](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](LICENSE)

## Overview
Device API is a RESTful microservice designed to manage devices, including CRUD operations, state management, and retrieval by various attributes.

## Table of Contents
- [Makefile Commands](#makefile-commands)
- [Environment Variables](#environment-variables)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Future improvements](#future-improvements)
- [License](#license)

## Makefile Commands

| Command            | Description                                         |
|--------------------|-----------------------------------------------------|
| `make all`         | Run all tests, then build and run                   |
| `make build`       | Build the binary                                    |
| `make clean`       | Remove build artifacts and tidy up dependencies     |
| `make run`         | Build and run the application                       |
| `make docker-up`   | Start the application with Docker Compose           |
| `make docker-down` | Stop and remove the Docker Compose containers       |
| `make swag`        | Generate docs                                       |
| `make lint`        | Run code linters                                    |
| `make test`        | Run all unit tests with race detection and coverage |
| `make mock`        | Generate mocks using Mockery                        |
| `make help`        | Display available commands                          |

### Endpoints

| Method   | Endpoint                 | Description                         |
|----------|--------------------------|-------------------------------------|
| `POST`   | `/devices`               | Create a new device                 |
| `PUT`    | `/devices/{id}`          | Update an existing device           |
| `PATCH`  | `/devices/{id}`          | Partially update an existing device |
| `GET`    | `/devices`               | Get all devices                     |
| `GET`    | `/devices/{id}`          | Get a device by ID                  |
| `GET`    | `/devices/brand/{brand}` | Get devices by brand                |
| `GET`    | `/devices/state/{state}` | Get devices by state                |
| `DELETE` | `/devices/{id}`          | Delete a device                     |

## Environment Variables
The following environment variables are used in the application:

| Name           | Suggested Value | Required |
|---------------|----------------|----------|
| `PORT`        | `8080`          | ✅       |
| `DB_HOST`     | `localhost`     | ✅       |
| `DB_PORT`     | `5432`          | ✅       |
| `DB_USER`     | `device_user`   | ✅       |
| `DB_PASSWORD` | `device_pass`   | ✅       |
| `DB_NAME`     | `device_db`     | ✅       |
| `DB_SSLMODE`  | `disable`       | ❌       |
| `LOG_ENABLED` | `true`          | ✅       |
| `LOG_LEVEL`   | `debug`         | ✅       |

## API Documentation
Device API uses Swagger for documentation. To view it, run the server and navigate to:
```
http://localhost:8080/swagger/index.html
```

## Future Improvements

- **Integration Testing**: Introduce end-to-end tests to cover API workflows.
- **Metrics & Monitoring**: Add metrics and dashboards for real-time monitoring.
- **Extended Logging**: Improve structured logging with trace IDs for better debugging.

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
