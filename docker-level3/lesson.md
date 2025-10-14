# **Level 3: The Recipe Book 📖 (Dockerfiles)**

### **Lesson 3 – From Manual Cooking to Recipes**

Right now, you’re like a chef who keeps making dishes by hand, then taking a photo when you like the result. That works, but… imagine you had to recreate it on another machine? Painful.

Instead, Docker has **Dockerfiles** → a text file that defines step-by-step how to build an image.
It’s like writing a **recipe card**:

Example:

```dockerfile
FROM ubuntu:latest
RUN apt-get update && apt-get install -y neofetch figlet cmatrix curl
CMD ["neofetch"]
```

* `FROM` = start with a base image.
* `RUN` = commands to run when building the image.
* `CMD` = the default command when the container runs.

You then build it like:

```bash
docker build -t my-uwubuntu-dockerfile .
```

…and boom → reproducible images.