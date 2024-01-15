# RQLite: Demo data initialization 

This container is designed to initialize a rqlite database with test data. It checks if the database is empty and, if so, loads predefined data.

## Prerequisites

- Docker
- Access to a rqlite database

## Environment Variables

- `RQLITE_HOST`: Host of the rqlite database.
- `RQLITE_PORT`: Port of the rqlite database.
- `RQLITE_USER`: Username for rqlite authentication.
- `RQLITE_PASSWORD`: Password for rqlite authentication.

## Run service
You can operate this container either as a standalone application or within a containerized environment. The database is required. It is designed to automatically generate the necessary tables should they not already exist.

Also: See the main [README.md](/README.md) to get information for a complete containerized setup. 