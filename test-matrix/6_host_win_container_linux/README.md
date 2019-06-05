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

When there's a mismatch between host OS and container, fswatch may not work well.  Take [Docker Desktop for Windows](https://docs.docker.com/docker-for-windows/troubleshoot/#inotify-on-shared-drives-does-not-work) for example:

> **Inotify on shared drives does not work**
>
> Currently, `inotify` does not work on Docker Desktop for Windows. This becomes evident, for example, when an application needs to read/write to a container across a mounted drive. Instead of relying on filesystem `inotify`, we recommend using polling features for your framework or programming language.


## Test from the container (inside)

Dig into the container, and change something under the `/mnt` directory:

```
% docker exec -it test6 sh
```