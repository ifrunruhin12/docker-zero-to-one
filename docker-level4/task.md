# **Task 4 – Make It Stay 🏠**

Goal: Build a Dockerfile that keeps the container alive instead of running once and dying.

Steps:

1. Create a new Dockerfile (start with Ubuntu again).
2. Install something fun if you like (figlet/neofetch).
3. Set the **CMD** to run `bash` instead of a one-shot command. (Hint: `CMD ["bash"]`)
4. Build it (`docker build -t uwubuntu-stay .`).
5. Run it (`docker run -it uwubuntu-stay`).

What should happen:

* You’ll land inside a shell **inside the container**.
* Now your container is “alive” until you exit it.

💡 Hint:

* This is the difference between **containers for tasks** (run → exit) and **containers for environments** (stay alive until you exit).
* Try running a one-shot version (`CMD ["figlet", "yo"]`) vs. the `bash` version to feel the difference.