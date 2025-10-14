# 📝 **Level 3 Review: Recipes Gone Wild 🍜**

1. **The cmatrix fail** – Yup, some packages (like `cmatrix` or `tzdata`) are **interactive** during install (asking you for keyboard layout, fonts, etc.).

   * In Docker builds, there’s **no human to press buttons**, so the build freezes or fails.
   * The pro fix? You pass environment variables or flags to make installs **non-interactive**. (We’ll deal with this later — that’s an “advanced recipe hack.”)

2. **The container exits instantly** – This one is key.

   * A container runs **only as long as its main process runs**.
   * Your Dockerfile had `CMD ["neofetch"]` → so Docker starts the container, runs `neofetch`, then when that’s done… *poof!* Container exits.
   * Same with `figlet`. It prints → exits → container dies.

👉 Translation: **Docker containers are not virtual machines.** They don’t sit around waiting unless you tell them to run something that keeps running.
