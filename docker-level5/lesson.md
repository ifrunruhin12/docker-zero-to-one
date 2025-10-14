# **Level 5: The Shipping Container 🚢**

### **Lesson 5 – Ports & Networking**

Right now, your containers are like secret labs — stuff happens inside, but the outside world (your host machine) can’t access it.

If you run a **web server inside a container**, you won’t see it on your browser unless you **map ports**. That’s like cutting a door in your shipping container so people can visit.

For example:

```bash
docker run -p 8080:80 nginx
```

* `-p 8080:80` → maps **port 8080 on your machine** to **port 80 inside the container**.
* So if nginx listens on port 80 inside the container, you can visit `http://localhost:8080` outside.

⚡ Fun fact: without `-p`, the server is running but only accessible *inside* the container.
