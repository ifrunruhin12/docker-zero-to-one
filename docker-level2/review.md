# 📝 **Level 2 Review: The Time Capsule ⏳**

* You **committed** your container → turned it into a new custom image (`my-uwubuntu`).
* You verified it with `docker images`.
* You ran it → your tools (neofetch, figlet, cmatrix, curl) lived on. 🎉
* You discovered the “catch”: any changes you make after the commit (like `python3` + `hello.py`) are **not remembered** unless you commit again.

👉 Translation: **commits freeze a moment in time**. Anything after that moment is just like doodles on a whiteboard—you’ll lose them unless you take another snapshot.

Now you understand **manual snapshots**, but Docker pros don’t live on `docker commit`—they use **Dockerfiles** to script image creation. That’s where we’re heading next.

