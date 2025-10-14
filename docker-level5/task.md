# **Task 5 – Hello from Inside 🔊**

Goal: Run a container with a web server and make it visible on your host.

Steps:

1. Pull the `nginx` image from Docker Hub (`docker pull nginx`).
2. Run it, mapping port `8080` on your machine to `80` inside (`docker run -p 8080:80 nginx`).
3. Open your browser and go to `http://localhost:8080`.
4. You should see the **nginx welcome page**.

💡 Hints:

* If port `8080` is busy, pick another (like `3000:80`).
* Use `docker ps` to confirm it’s running.
* Stop it with `docker stop <container_id>`.


This task teaches you how to **expose apps to the outside world** — which is the foundation of deploying services with Docker. 🌐