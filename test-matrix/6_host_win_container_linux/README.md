# 6: Host OS (Windows) + Container (Linux)

| Host OS | Native app | Linux Container | Windows Container |
|---------|------------|-----------------|-------------------|
| Linux   | 1          | 4               | N/A               |
| Mac     | 2          | 5               | N/A               |
| Windows | 3          | **★ 6 ★** (LCOW)| 7                 |


Type `run.bat` to track the changes of this directory.

```
.
├── run.bat
└── site
    ├── css
    │   └── main.css
    ├── index.html
    └── js
        └── main.js

3 directories, 4 files
```

## Test from the host (outside)

Change something under the current directory from the host.

NOTE: it should *not* work.


## Test from the container (inside)

Dig into the container, and change something under the `/mnt` directory:

```
% docker exec -it test6 sh
```