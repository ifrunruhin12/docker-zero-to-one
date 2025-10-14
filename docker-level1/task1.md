# Task 1 – The Chef’s Kitchen 🍳

Goal: Pretend you’re inside a brand new kitchen (Ubuntu container). Cook something (install + run a program). Then leave and see what happens.

Steps (fun edition):

1. Run an interactive Ubuntu container.
2. Inside it, install a package (like curl or figlet).
3. Run it and enjoy the output.
4. Exit the container.
5. Run a new Ubuntu container… see if your “cooking” is still there.

💡 Hints:

- Use apt-get update before installing stuff in Ubuntu.
- Try a fun package: figlet makes ASCII art banners.
- Notice: “Oh damn, my installed stuff is gone when I restart the container!”