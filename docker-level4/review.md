# 📝 **Level 4 Review: The Container That Stayed**

* First, you proved the “one-shot” behavior again (`figlet yo` → instant exit). ✅
* Then you swapped CMD to `bash` → boom, now the container **stays alive** until you decide to exit. ✅
* You installed your dev tools (`neofetch`, `figlet`, `python3`, `neovim`) → turning your container into a **mini dev workstation**. ✅
* You even wrote Python scripts and ran them inside → meaning you just created your **own dev environment in Docker**. ✅

👉 The key lesson you just unlocked: **A container lives as long as its main process lives.**
You now understand why most “real” containers run something long-living like `bash`, `nginx`, or `postgres`.

This is one of Docker’s deepest mental shifts → containers aren’t VMs, they’re **process sandboxes**.
