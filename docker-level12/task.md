## 🧪 Task 12 – Externalize Your Secrets

**Goal**: Move all hardcoded passwords from your `docker-compose.yml` to a `.env` file.

### Steps:

1. Navigate to `docker-level10/`.
2. Create a `.env` file with all your sensitive variables:
   ```env
   POSTGRES_USER=admin
   POSTGRES_PASSWORD=secret
   POSTGRES_DB=testdb
   DATABASE_URL=postgres://admin:secret@db:5432/testdb?sslmode=disable
   ```
3. Create a `.env.example` file (same structure, but with placeholder values):
   ```env
   POSTGRES_USER=your_username
   POSTGRES_PASSWORD=your_password
   POSTGRES_DB=your_database
   DATABASE_URL=your_connection_string
   ```
4. Update your `docker-compose.yml` to use `${VARIABLE_NAME}` instead of hardcoded values.
5. Add `.env` to your `.gitignore` (so it doesn't get committed).
6. Test it: `docker compose up` → should work exactly the same.
7. Verify `.env.example` is in Git, but `.env` is not.

### 💡 Hints (if stuck):

- The syntax in `docker-compose.yml` is `${VARIABLE_NAME}`.
- Compose reads `.env` automatically from your current directory (no extra command needed).
- If a variable isn't in `.env`, it falls back to the system environment (or stays empty).
- Check that your services still connect properly (`curl http://localhost:8080/api/users`).
- If you mess up, just check `docker logs <container_id>` to see what went wrong.

### 🎯 Why this matters:

This is **non-negotiable** in real work. Every company, every DevOps team, every deployment pipeline uses this pattern. You're learning the right way from the start. 💪
