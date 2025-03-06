# Internet Tester

This is a simple internet tester application written in Go. It performs the following tasks and stores the results in a PostgreSQL database:

-   **Speed Test:** Measures internet speed (download and upload) and latency (ping).
-   **Packet Loss Test:** Analyzes packet loss during the test.
-   **Request Logging:** Logs HTTP requests made to specified addresses, including response time and status code.

## Environment Variables

-   **DB_CONN_STR:** Connection string for the PostgreSQL database.
    -   Default: `"host=postgres user=user password=password dbname=pingdb port=5432 sslmode=disable"`
-   **TARGET_URL:** The URL to be used for request logging.
    -   Default: `"https://google.com.br"`
-   **TIMEOUT_IN_SECS:** Timeout for HTTP requests in seconds.
    -   Default: `"5s"`
-   **INTERVAL_BETWEEN_REQUESTS_IN_SECS:** Interval between internet tests in seconds.
    -   Default: `"10s"`

## Running the Application

### Using .env File

1.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

2.  **Create .env File:**
    Create a [.env](http://_vscodecontentref_/0) file in the root directory of the project and add the following environment variables:
    ```env
    DB_CONN_STR="host=localhost user=user password=password dbname=pingdb port=5432 sslmode=disable"
    TARGET_URL="https://google.com.br"
    TIMEOUT_IN_SECS="5"
    INTERVAL_BETWEEN_REQUESTS_IN_SECS="10"
    ```
    Modify the values according to your needs.

3.  **Run the Application:**
    ```bash
    go run main.go
    ```

### Using Docker Compose

Ensure you have Docker and Docker Compose installed on your system.

1. **Build and Run the Docker Container:**
    ```sh
    docker compose up -d --build
    ```

2. **Stop and Remove the Containers:**
    ```sh
    docker compose down
    ```

3. **Stop and Remove the Containers and the PostgreSQL Data:**
    ```sh
    docker compose down --volumes
    ```

## PostgreSQL Credentials

To connect to the PostgreSQL database, use the following credentials:

- **Host:** `localhost`
- **Port:** `5432`
- **Database:** `pingdb`
- **User:** `user`
- **Password:** `password`

## Viewing Inserted Records

You can copy the contents of the [query.sql](./query.sql) file to a PostgreSQL client to view the records inserted into the tables. For example:

1. Open your PostgreSQL client (e.g., `psql`).
2. Connect to the database using the credentials provided above.
3. Copy the contents of [query.sql](./query.sql) and execute it in the PostgreSQL client to see the inserted records.

## Exporting Data to CSV

You can export the data to a CSV file using database management tools like DBeaver. Follow these steps or this [tutorial](https://dbeaver.com/2024/09/19/how-to-export-data-in-dbeaver/):

1. Open DBeaver and connect to your PostgreSQL database using the credentials provided above.
2. Navigate to the table you want to export.
3. Right-click on the table and select "Export Data".
4. Choose "CSV" as the format and configure the export settings as needed.
5. Complete the export process to generate the CSV file.
