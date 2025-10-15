# **Task 7 – The Diary Container 📔**

Goal: Learn how to persist data outside a container.

Steps:

1. Run an `ubuntu` container and create a file inside (`echo "secret" > /data/diary.txt`).
2. Exit the container and start a new one → file is gone (amnesia).
3. Now run with a volume:

   ```bash
   docker run -it -v $(pwd)/mydata:/data ubuntu
   ```

   * This mounts your host folder `mydata` into `/data` inside the container.
4. Create a file inside `/data` again.
5. Exit, restart container → file still there, because it lives on your host.

💡 Hints:

* `-v host_path:container_path` = volume binding.
* `$(pwd)` just means “current directory” in Linux.