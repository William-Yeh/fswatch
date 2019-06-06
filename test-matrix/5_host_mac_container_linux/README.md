# 5: Host OS (Mac) + Container (Linux)

| Host OS | App run as<br/>native app | App run as<br/>Linux Container |  App run as<br/>Windows Container | ‖<br/>‖ | App run in<br/>K8s (Linux Container) |
|---------|------------|-----------------|-------------------|----|-----------------------|
| Linux   | 1          | 4               | N/A               | ‖ | 8             |
| Mac     | 2          | **★ 5 ★**       | N/A               |
| Windows | 3          | 6               | 7                 |


Type `run.sh` to track the changes of this directory.

```
.
├── run.sh
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

Docker Desktop for Mac doesn't have much trouble here (AMAZING!) thanks to excellent implementation of **osxfs**.  See "[File system sharing (osxfs)](https://docs.docker.com/docker-for-mac/osxfs/)" and "[Performance tuning for volume mounts (shared filesystems)](https://docs.docker.com/docker-for-mac/osxfs-caching/)" articles for more information.

> **File system events**
>
> Most `inotify` events are supported in bind mounts, and likely `dnotify` and `fanotify` (though they have not been tested) are also supported. This means that file system events from macOS are sent into containers and trigger any listening processes there.

## Test from the container (inside)

Dig into the container, and change something under the `/mnt` directory:

```
% docker exec -it test5 sh
```