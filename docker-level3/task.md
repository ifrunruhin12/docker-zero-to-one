# **Task 3 – Write Your First Recipe 🍜**

Goal: Create a Dockerfile that builds an Ubuntu image with at least 2 tools pre-installed, then run it.

Steps:

1. Create a folder `docker-level3` and inside it, make a file called `Dockerfile`.
2. Write a Dockerfile starting from `ubuntu:latest`.
3. Use `RUN` to install at least 2 tools you like.
4. Set a `CMD` (like `figlet "Welcome to uwubuntu!"`).
5. Build it with `docker build -t uwubuntu-v2 .`
6. Run it with `docker run -it uwubuntu-v2`.

💡 Hints:

* Remember: you need `apt-get update` before installs.
* Use `-y` so you don’t get stuck on “Do you want to continue? [Y/n]”.
* If CMD is a fun command (like `neofetch`), you’ll see it run instantly when the container starts.