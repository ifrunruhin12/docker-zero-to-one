# Task 2 – The Time Capsule ⏳

Goal: Create your own custom Ubuntu image that remembers your hard work.

Steps:

1. Run Ubuntu again and install at least 2–3 tools (neofetch + anything else).
2. Exit the container.
3. Use docker commit to save this modified container as a new image (give it a fun name, e.g., ubuntu-gamer-edition).
4. Run a fresh container from your new image.
5. Test if your tools are still there.

💡 Hint:

- `docker ps` -a shows the container IDs of stopped containers you can commit.
- Make sure you commit the right one (the one where you installed stuff).