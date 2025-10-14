# **Level 4: Keeping the Lights On 💡**

### **Lesson 4 – Why Your Container Yeets Itself**

Think of a Docker container as a **robot programmed to do one task**:

* If you tell it to play one note on the piano, it does it and walks off stage. 🎹🚶
* If you want it to stick around, you need to run something that **doesn’t stop** (like `bash` or a server).

That’s why in the real world, containers often run:

* **Web servers** (nginx, apache, etc.)
* **Background daemons** (databases, message queues)
* Or an **interactive shell** (so you can poke around).