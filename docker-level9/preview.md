# 🚀 Level 9 – “The Secret Tunnels Between Containers” 🕳️

Up until now, you’ve been talking **host → container**. But what if containers need to talk **container → container** without exposing ports to the host?

## 🎓 Lesson Preview

* Docker creates a default `bridge` network for containers.
* By default, containers on this network can reach **the outside internet**, but not each other by name.
* If you create a **user-defined network**, containers can talk to each other using their **names** as hostnames.