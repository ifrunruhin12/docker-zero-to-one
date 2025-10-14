# 📝 Level 1 Review: The Amnesiac Container 🧠💨

- You cooked up some tools (neofetch, curl, figlet).
- You flexed with ASCII art banners.
- You even peeked at curl’s spellbook 📜 (don’t worry, curl is confusing even for pros).
- Then _bam!_ — when you left and came back, your container was **clean slate mode** again.

That’s not a bug, that’s **the way of Docker**. Containers are **stateless and ephemeral** by design. Think of them like **Snapchat streaks**: gone if you don’t save them.

Why?

Because containers are built from images (blueprints). If you want your “neofetch-powered Ubuntu” to persist, you need to bake those changes into an image.