
# 📝 **Level 7 Review – The Diary Container**

You:

* Created a new level folder ✅
* Ran an Ubuntu container with a **volume mount** (`-v $(pwd)/mydata:/data`) ✅
* Wrote a note inside `/data` (`note1.txt`) ✅
* Exited and verified the file **persisted on your host** inside `mydata/note1.txt` ✅

## What you learned:

* **`cat` vs `echo` confusion**: `cat "something" > file.txt` doesn’t work because `cat` expects a file to read, not a string. You fixed it by using `echo "..." > file.txt`. That’s a classic shell mistake — you caught it perfectly. 💯
* The **magic of volumes**: files inside `/data` survive across container runs, and you can even open/edit them from outside the container. This is *the key trick* for databases, logs, configs — basically any real-world app that doesn’t want amnesia.
* You officially **broke through Docker’s “goldfish memory” problem** 🐠 → volumes are the antidote.

## Level 7 status: ✅ **Mission Complete!**

You now understand how to give containers “long-term memory.” That’s huge — without this, Docker is just glorified temporary VMs. Now we can build apps that persist.