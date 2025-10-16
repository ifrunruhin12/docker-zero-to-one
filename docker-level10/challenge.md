## 🔥 Level 10 — Bonus Secret Challenge

Make nginx talk to Postgres and show DB data in your browser.

1. Keep your `docker-compose.yml` with `web` and `db`.
2. Add a third service: `api` (a tiny Go or Python app).

   * The `api` connects to Postgres (`host=db`, port 5432, user `admin`, pass `secret`).
   * It has an endpoint `/users` that returns something simple like “users: [Alice, Bob]” fetched from Postgres.
3. Configure nginx as a reverse proxy in front of `api`.

   * Requests to `/api` go to your `api` container.
   * Requests to `/` show the nginx welcome page.
4. Visit `http://localhost:8080/api/users` → you should see data from Postgres, traveling through `api`, into `nginx`, out to your browser.

This chain proves:
**browser → nginx → api container → postgres container → back out.**

That’s *container networking in action* under Compose.
