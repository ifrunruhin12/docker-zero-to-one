# 🎯 Level 12 – Secrets & Environment Variables 🔐

Now let's talk about something that'll make you go "oh crap, I've been doing this wrong the whole time."

## 🎓 Lesson 12 – Never Hardcode Secrets

Look at your `docker-level10/docker-compose.yml`:

```yaml
environment:
  POSTGRES_USER: admin
  POSTGRES_PASSWORD: secret
  DATABASE_URL: postgres://admin:secret@db:5432/testdb?sslmode=disable
```

**DANGER ZONE.** 🚨

Your password is sitting there in plain text. If you push this to GitHub, it's like leaving your house key under the doormat and announcing it on Twitter.

In the real world:
- **Development**: You use `.env` files (which you `.gitignore`).
- **Production**: You use secrets management (AWS Secrets Manager, Vault, Kubernetes Secrets, etc.).

Docker has a few ways to handle this:

### 1. **`.env` File** (for local dev)

Create a `.env` file in your `docker-level10/` folder:

```env
POSTGRES_USER=admin
POSTGRES_PASSWORD=super_secret_password_123
DB_NAME=testdb
API_PORT=8080
```

Then in your `docker-compose.yml`, reference it:

```yaml
environment:
  POSTGRES_USER: ${POSTGRES_USER}
  POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
  POSTGRES_DB: ${DB_NAME}
```

Docker Compose **automatically reads** `.env` and substitutes the values. No hardcoding needed.

### 2. **`.env.example`** (for team sharing)

Commit this to Git so teammates know what variables are needed:

```env
POSTGRES_USER=admin
POSTGRES_PASSWORD=changeme
DB_NAME=testdb
API_PORT=8080
```

But **never commit the real `.env`**. Add it to `.gitignore`:

```
.env
.env.local
*.log
```

### 3. **`docker compose --env-file`** (override per environment)

```bash
docker compose --env-file .env.prod up
```

Different configs for dev, staging, prod. No code changes needed.

### Why this matters:

- If your repo goes public (oops!), attackers don't have your passwords.
- Different environments can have different secrets without changing code.
- This is **industry standard**. Companies get hacked because devs hardcode secrets. Don't be that dev. 🚫
