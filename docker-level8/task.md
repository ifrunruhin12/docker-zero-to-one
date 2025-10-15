# 🧩 Task 8:

1. Run two **different web servers** in containers at the same time (e.g., `nginx` and your Go “hello” app).

   * Map `nginx` to `localhost:8080`.
   * Map your Go app to `localhost:9090`.

2. Open both in your browser → confirm you’re talking to two totally separate apps, both isolated but coexisting.

3. **Bonus spooky challenge 👻**:
   Run **two `nginx` containers** at once, but map one to `8081` and the other to `8082`. Can you figure out *why both can still run even though they’re both `nginx` listening on port 80 inside*?

---

💡 Hint if you get stuck: Remember, the **inside container port** (second number) can be the same, but the **host port** (first number) must be unique.
