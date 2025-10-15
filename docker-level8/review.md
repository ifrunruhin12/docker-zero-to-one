# 📝 Review of Level 8 (so far)

## 🔑 What you nailed

1. **Realizing you need a Dockerfile first** → yes, Docker can’t run Go source directly, it needs a container image you build. ✅
2. **Dockerfile placement** → yup, it has to be where your Go source (and `go.mod`) lives, unless you copy from elsewhere. ✅
3. **Understanding `-p HOST:CONTAINER`** → you figured out that your Go app listens on `:8080` inside the container, so you had to map host port → container port correctly. ✅

   * `docker run -p 9090:8080` → host’s `9090` maps to container’s `8080`.
4. **Port conflicts** → you discovered host ports can’t be reused (`nginx` on `8080` vs Go app on `8080` conflict). ✅
5. **Multiple nginx containers** → you ran one on `8080:80`, another on `8081:80`, both worked fine. ✅

🔥 That’s literally the main networking lesson I wanted you to get from this task. You’re ahead of schedule.

---

## ⚠️ The “Why tho?” moments

### 1. Why didn’t `/events` SSE show updates?

This is not Docker’s fault — this is how **Postman** handles **Server-Sent Events (SSE)**.

* SSE works by **keeping the HTTP connection open** and streaming text chunks.
* Postman… doesn’t show live-streaming output well. It buffers, and you see nothing.
* Try in your browser (`http://localhost:9090/events`) or `curl`:

  ```sh
  curl http://localhost:9090/events
  ```

  You’ll see the stream of `data: <timestamp>` messages every second. 🌀

---

### 2. Why did `docker run -p 8082:8080 nginx` fail?

Here’s the ghost 👻:

* Nginx **inside the container** listens on **port 80** by default.
* When you mapped `8082:8080`, Docker tried forwarding host `8082` → container `8080`.
* But nothing was listening on `8080` inside the container → connection reset.

💡 Fix: always map to the port the app listens on inside.
For nginx → it’s **80**, so you should do `-p 8082:80`.

---

### 3. Why can multiple containers use port `80` inside?

Because inside their isolated network namespaces, they each get their own `80`.

* Think of each container as its own mini-computer → they don’t conflict with each other.
* On the host, though, ports are unique. You can’t bind `8080` twice.

So yes:

* Container A: listens on `80` internally, mapped to host `8080`.
* Container B: also listens on `80` internally, mapped to host `8081`.
* Both coexist just fine.

---

## 🪦 Level 8 Status

✅ You’ve mastered host-to-container port mapping.
✅ You’ve seen how services can run on the same container port but different host ports.
✅ You ran into a real-world dev gotcha (SSE testing in Postman).

You basically **exorcised the haunted house**.
