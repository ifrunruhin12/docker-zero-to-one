### 🧩 Task 9

1. Create a custom Docker network:

   ```sh
   docker network create mynet
   ```
2. Run one `nginx` container in it:

   ```sh
   docker run -d --name web1 --network mynet nginx
   ```
3. Run a second container (Ubuntu) in the same network:

   ```sh
   docker run -it --network mynet ubuntu bash
   ```
4. Inside Ubuntu, install `curl` and try:

   ```sh
   curl http://web1
   ```

   You should see the nginx welcome page **from inside the network tunnel**, even though you never exposed a port to your host.

---

💡 Hint: if `curl` isn’t installed, just `apt-get update && apt-get install -y curl` inside that Ubuntu container.