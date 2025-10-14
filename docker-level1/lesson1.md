# Level 1: Container Playground 🎮

## Lesson 1 – Containers Are Like Disposable Sandboxes

Right now, your containers are like **Minecraft creative mode worlds**—you build inside, then you can throw them away without hurting your actual PC. Every time you `docker run ubuntu bash`, you’re starting a fresh sandbox.
But here’s the catch: if you install something inside that container and then exit… it’s gone. Poof.
Why? Because containers are **ephemeral** (they don’t keep state unless you tell them to).

This level is about messing around inside a container and learning how to:

- Enter containers (`-it` = interactive mode).
- Install things inside.
- Understand that once you exit, those changes vanish unless you commit them.