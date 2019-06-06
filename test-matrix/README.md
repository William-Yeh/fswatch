# Test matrix for fswatch

Test whether the fswatch works in native app mode and in container mode.

| Host OS | App run as<br/>native app | App run as<br/>Linux Container |  App run as<br/>Windows Container | ‖<br/>‖ | App run in<br/>K8s (Linux Container) |
|---------|------------|-----------------|-------------------|----|-----------------------|
| Linux   | [1](1_native_linux)          | [4](4_host_linux_container_linux)               | N/A               | ‖ | [8](8_k8s_linux)             |
| Mac     | [2](2_native_mac)          | [5](5_host_mac_container_linux)               | N/A               |
| Windows | [3](3_native_win)          | [6](6_host_win_container_linux) (LCOW)        | [7](7_host_win_container_win) (WCOW)          |


When there's a mismatch between host OS and container (esp. [6](6_host_win_container_linux)), fswatch may not work well.  Take [Docker Desktop for Windows](https://docs.docker.com/docker-for-windows/troubleshoot/#inotify-on-shared-drives-does-not-work) for example:

> **Inotify on shared drives does not work**
>
> Currently, `inotify` does not work on Docker Desktop for Windows. This becomes evident, for example, when an application needs to read/write to a container across a mounted drive. Instead of relying on filesystem `inotify`, we recommend using polling features for your framework or programming language.


Docker Desktop for Mac doesn't have much trouble here ([5](5_host_mac_container_linux)) thanks to excellent implementation of **osxfs**.  See "[File system sharing (osxfs)](https://docs.docker.com/docker-for-mac/osxfs/)" and "[Performance tuning for volume mounts (shared filesystems)](https://docs.docker.com/docker-for-mac/osxfs-caching/)" articles for more information.

> **File system events**
>
> Most `inotify` events are supported in bind mounts, and likely `dnotify` and `fanotify` (though they have not been tested) are also supported. This means that file system events from macOS are sent into containers and trigger any listening processes there.
