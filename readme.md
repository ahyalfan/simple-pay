# Golang Base Project

This is a base project for implementasi simple e wallet using Golang with the Fiber framework, Gorm for database operations, Redis for cahce ,midtrans for handle real money, filebeat kibana for log core,and PostgreSQL as the database.godotenv Configuration is managed via an env file.

## Features

- **Transaction core**: Core of transaction e money
- **TopUp**: Top up transaction use real money from mitrans api
- **Notification Realtime**
- **Mobile PIN**: a mobile PIN a one transaction
- **Detect Location**: Detect location a user when login
- **Log core**: centralized log use kibana and filebeat
- **Queue Server**: implement queue server use redis connection
- **Fiber Framework**: Fast and minimalistic web framework for Go.
- **GORM**: SQL builder and query library for Go.
- **PostgreSQL**: Relational database for storing application data.
- **Env File**: Simple configuration management using environment variables.
- **Validator**: simple validator for your application
- **auth basic**: basic authentication token for your application

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Golang](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [Filebeat](https://www.elastic.co/beats/filebeat)

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/shellrean/golang-base-project-clean-directory.git
   cd golang-base-project-clean-directory
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Create and configure `.env` file:**

   Create a `.env` file in the root directory and add your configuration variables.

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASS=yourpassword
   DB_NAME=yourdbname

   SERVER_HOST=localhost
   SERVER_PORT=8700

   MAIL_HOST=smtp.gmail.com
    MAIL_PORT=587
    MAIL_USERNAME=email
    MAIL_PASSWORD=<password>

    REDIS_ADDR=
    REDIS_PASSWORD=
    REDIS_DB=

    MIDTRANS_KEY=server-key
    # production atau sandbox
    MIDTRANS_ENV=

    QUEUE_REDIS_ADDR=localhost:6379
    QUEUE_REDIS_PASSWORD=
    QUEUE_REDIS_DB=0
   ```

4. **Set up PostgreSQL database:**

   Make sure your PostgreSQL server is running and create a database matching your `.env` configuration.

   ```sh
   psql -U yourusername -c "CREATE DATABASE yourdbname;"
   ```

5. **Set up Docker compose:**
   Make sure your Docker server is running and create a Docker compose file

   ```sh
   docker compose up
   ```

6. **clone server redis queue from my github repository**

   ```sh
   git clone https://github.com/ahyalfan/redis-queue-server-go.git
   ```

### Running the Application

Start the application with the following command:

```sh
go run main.go
```
