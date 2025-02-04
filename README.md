# golang-starter

## Overview

`golang-starter` is a boilerplate project for starting new Go projects. It includes a basic setup for a web server, database migrations, and testing. The project uses Docker for containerization and includes a Makefile for common tasks.

## Prerequisites

- Docker
- Docker Compose
- Go

## Getting Started

### Running the Project

1. **Start the project:**

   ```sh
   make up
   ```

   This command will start the Docker containers for the project.

2. **Access the server:**

   Open your browser and navigate to [`http://localhost:8080`](http://localhost:8080). Swagger documentation is available at [`http://localhost:8080/swagger`](http://localhost:8080/swagger).

3. **Open a container bash window:**

   ```sh
   make bash
   ```

### Stopping the Project

To stop the running containers:

```sh
make stop
```

To remove the containers and associated volumes:

```sh
make down
```

### Running Tests

To run the tests:

```sh
make tests
```

### Running Database Migrations

To run the database migrations:

```sh
make migrate
```

To create a new migration file:

```sh
make migration
```

You will be prompted to enter a migration name. This will create new `.up.sql` and `.down.sql` files in the [`internal/migrations`](internal/migrations) directory.

### Viewing Logs

To view the logs:

```sh
make logs
```

## Project Structure

- [`cmd`](cmd): Contains the main entry point for the application.
- [`internal`](internal): Contains the core application code.
  - [`app/`](internal/app): Wire up use cases.
  - [`config/`](internal/config): Configuration handling.
  - [`container/`](internal/container): Dependency injection container.
  - [`http/`](internal/http): HTTP server and routing.
  - [`log/`](internal/log): Logging setup.
  - [`migrations/`](internal/migrations): Database migration files.
  - [`todo/`](internal/todo): Example feature module.
- [`pkg`](pkg): Contains reusable packages.
  - [`middleware/`](pkg/middleware): HTTP middleware.
  - [`mysql/`](pkg/mysql): MySQL database utilities.
  - [`validator/`](pkg/validator): Request validation.
- [`docs`](docs): API documentation.
- [`.github/workflows`](.github/workflows): GitHub Actions workflows for CI/CD.

## Configuration

Configuration is managed using the [`config.yml`](config.yml) file. Environment variables can override the configuration values. The environment variables should be prefixed with `MYVAR_` and use `_` as the delimiter.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
