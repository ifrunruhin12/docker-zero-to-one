## ⚡ Level 10: *The Secret Society of Compose*

**Theme:** *"One command to rule them all"*

Spinning up multiple containers with custom networks, ports, volumes, etc. by hand is fun for practice… but imagine doing that for a web app with **frontend + backend + database + cache**. Yikes.

That’s where **Docker Compose** comes in. One YAML file, one command → your whole app starts like magic.

---

### 🧪 Task 10

1. Make a new folder `docker-level10`.

2. Inside it, create a `docker-compose.yml` with two services:

   * `web`: use nginx (expose port 8080 → 80).
   * `db`: use postgres (set env vars: `POSTGRES_USER=admin`, `POSTGRES_PASSWORD=secret`).

   Example scaffold:

   ```yaml
   version: "3.9"
   services:
     web:
       image: nginx
       ports:
         - "8080:80"
     db:
       image: postgres
       environment:
         POSTGRES_USER: admin
         POSTGRES_PASSWORD: secret
   ```

3. Run it:

   ```bash
   docker compose up
   ```

   (yes, newer Docker uses `docker compose` not `docker-compose`)

4. Open `localhost:8080` → you should see nginx.

5. Run `docker ps` → notice both nginx *and* postgres are running together.

6. Stop everything cleanly with:

   ```bash
   docker compose down
   ```

---

### 🎯 What you’ll learn:

* How **Compose** orchestrates multiple containers in one go.
* How Compose automatically creates a network, so services (`web`, `db`) can talk using their names.
* Why Compose is the **gateway drug** to Kubernetes (spoiler: YAML everywhere).
