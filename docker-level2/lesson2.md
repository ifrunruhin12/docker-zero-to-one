# Level 2: The Container Photographer 📸

## Lesson 2 – Capturing Your Work

Docker lets you save the state of a container into a new image.

- Base image = like Minecraft’s default flat world.
- Modified container = you added a castle, a farm, and 7 cats.
- Commit = “yo, snapshot this world as a new template.”

Commands to learn:

- `docker ps -a` → see running + stopped containers.
- `docker commit <container_id> my-ubuntu-with-neofetch` → snapshot your masterpiece.
- `docker images` → check your new blueprint exists.
- `docker run -it my-ubuntu-with-neofetch` → spawn a fresh container with all your upgrades.