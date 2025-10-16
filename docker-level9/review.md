## 🔹 1. What does `docker network create mynet` actually do?

When you run:

```bash
docker network create mynet
```

Docker sets up a **new virtual network** inside your system using its own network driver (by default `bridge`).

Here’s what happens under the hood:

* It creates a **Linux bridge interface** (basically a virtual switch) on your host, named something like `br-<random_id>`.
* Any container you attach to `mynet` gets a **virtual ethernet interface** connected to that bridge.
* Docker also sets up **internal DNS** so containers on that network can talk to each other by their container names (`web1`, `web2`, etc.), instead of IP addresses.

So `mynet` is like a private LAN that only containers you attach to it can use. Containers outside that network can’t access it unless you connect them explicitly.

---

## 🔹 2. What happened when you did `curl http://web1`?

Let’s trace it:

1. You typed `curl http://web1`.

   * `curl` is a command-line tool that makes an HTTP request.
   * `http://web1` means: "send an HTTP request to a server named `web1` on port 80 (default)."

2. DNS inside Docker kicks in:

   * Since you’re inside the same Docker network (`mynet`), Docker’s embedded DNS resolves `web1` → to the container’s internal IP (like `172.18.0.2`).

3. Your curl request reaches the Nginx container (`web1`).

   * Nginx listens on port 80 inside the container.
   * It receives the HTTP GET request from curl.

4. Nginx responds with its **default welcome page** (HTML code).

5. Curl just prints the **raw HTML response** to your terminal.

   * Terminals don’t render HTML (like a browser does), they just display text.
   * So you saw the welcome page source (HTML tags and all), instead of a styled webpage.

If you opened the same URL in a browser (e.g., `http://localhost:8080` when port-mapped), the browser would render that HTML properly with colors, headings, etc.

---

✅ So in short:

* `mynet` made containers talk via DNS instead of IP.
* `curl` sent an HTTP request, got raw HTML back, and dumped it in your terminal.
* That’s why it looked like “HTML code” instead of a styled web page.

Ahh you caught me slippin’ 😅 — you’re totally right. Let’s rewind and get our bearings.


## 🧾 Level 9 Review

You crushed container networking:

* Created a custom network with `docker network create mynet`.
* Learned how container names become DNS inside the network (`curl http://web1`).
* Discovered why container IDs are long but 12 chars are enough.
* Understood that Docker basically ships its own DNS for containers in a network.

**Big insight:** This is the foundation for multi-service apps (like web + DB). Without this, Docker Compose would just be duct tape. You now speak “container-to-container.”
