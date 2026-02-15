# Database setup

## Postico (or any client)

Postgres requires a **database name** to connect. If you don't have one yet:

1. **First connection** – use the default database:
   - **Host:** `localhost`
   - **Port:** `5432`
   - **User:** `admin` (or your user)
   - **Password:** (your password)
   - **Database:** `postgres`  ← use this; it always exists.

2. **Create the app database** – in the SQL query tab run:
   ```sql
   CREATE DATABASE metadb;
   ```

3. **Connect to `metadb`** – disconnect and create a new connection with **Database:** `metadb`, same host/user/password.

4. **Create tables** – open `db/migration/000001_init_schema.up.sql` and run its contents in the `metadb` connection (or use the script below).

## From the command line

With `psql` in PATH and same user/password:

```bash
# Create database (connects to default 'postgres' first)
psql "postgres://admin:admin1234@localhost:5432/postgres?sslmode=disable" -c "CREATE DATABASE metadb;"

# Run migrations
psql "postgres://admin:admin1234@localhost:5432/metadb?sslmode=disable" -f db/migration/000001_init_schema.up.sql
```

Your `.env` already points at `metadb`; once it exists and the migration has been run, tests and the API will work.
