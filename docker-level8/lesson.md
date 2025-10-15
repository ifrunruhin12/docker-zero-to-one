# 🚀 **Level 8 – The Haunted House of Ports 🏚️**

Now we’re going deeper into **networking magic**.
You’ve already seen `-p 8080:8080` in action, but we’ll crank it up a notch.

## 🎓 Lesson:

* Every container gets its own **private network namespace** → think of it like having its own little LAN.
* By default, containers can **talk to the internet** (outbound), but not vice versa (inbound).
* `-p HOST:CONTAINER` punches a hole through the isolation wall:

  * `docker run -p 5000:8080 my-app` → requests to host’s **localhost:5000** are forwarded to container’s **8080**.
* If you run multiple containers, each can bind to **different host ports** → but inside, they can all still run on **8080** safely.
* By default, containers don’t see each other unless you put them on the same custom network (that’s coming in future levels 🔮).


