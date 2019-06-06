# 4: Host OS (Linux) + Container (Linux)

| Host OS | App run as<br/>native app | App run as<br/>Linux Container |  App run as<br/>Windows Container | ‖<br/>‖ | App run in<br/>K8s (Linux Container) |
|---------|------------|-----------------|-------------------|----|-----------------------|
| Linux   | 1          | **★ 4 ★**       | N/A               | ‖ | 8             |
| Mac     | 2          | 5               | N/A               |
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

A `Vagrantfile` is also provided in case that you're familiar with Vagrant.


## Test from the host (outside)

Change something under the current directory from the host (maybe Vagrant).


## Test from the container (inside)

Dig into the container, and change something under the `/mnt` directory:

```
% docker exec -it test4 sh
```