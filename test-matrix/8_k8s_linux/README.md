# 8: Kubernetes (Linux)

| Host OS | App run as<br/>native app | App run as<br/>Linux Container |  App run as<br/>Windows Container | ‖ | App run in<br/>K8s (Linux Container) |
|---------|------------|-----------------|-------------------|----|-----------------------|
| Linux   | 1          | 4               | N/A               | ‖ | **★ 8 ★**             |
| Mac     | 2          | 5               | N/A               |
| Windows | 3          | 6               | 7                 |



Use [Skaffold](https://skaffold.dev/) to track the changes of ConfigMap stored in `fswatch-config.yml`.

```
.
├── fswatch-config.yml
├── fswatch-deployment.yml
└── skaffold.yaml

0 directories, 3 files
```


## Usage

1. Create a `fswatch` namespace for this test:

   ```
   % kubectl create ns fswatch
   ```

2. Use [Skaffold](https://skaffold.dev/) to load configmap content (`fswatch-config`) and bring up fswatch app (`deployment/fswatch`) in the `fswatch` namespace:

   ```
   % skaffold dev  -n fswatch
   ```

3. Watch for fswatch logs:

   ```
   % kubectl logs -f deployment/fswatch  -n fswatch
   ```

4. Edit the ConfigMap content, eithor by:

   ```
   % kubectl edit configmap fswatch-config  -n fswatch
   ```

   or by:

   ```
   % vi fswatch-config.yml
   ```


## Demo

You can see the demo at:

[![asciicast](https://asciinema.org/a/250736.svg)](https://asciinema.org/a/250736)


### Caution about symbolic links

However, you should watch for *directories* instead of merely for *files*.  It is because Kubernetes may use symbolic links to point to versioned ConfigMap volumes, and `inotify` doesn't work well with such symbolic links.

Let's retry the demo, but this time we'll focus on the directory layout from the pod's point of view.

[![asciicast](https://asciinema.org/a/251141.svg)](https://asciinema.org/a/251141)

The demo from [0:38](https://asciinema.org/a/251141?t=0:38) to 1:55 shows the directory layout from the pod's point of view:

```
/mnt/site # ls -al
total 12
drwxrwxrwx    3 root  root   4096 Jun 12 06:19 .
drwxr-xr-x    1 root  root   4096 Jun 12 06:19 ..
drwxr-xr-x    2 root  root   4096 Jun 12 06:19 ..2019_06_12_06_19_15.187277003
lrwxrwxrwx    1 root  root     31 Jun 12 06:19 ..data -> ..2019_06_12_06_19_15.187277003
lrwxrwxrwx    1 root  root     15 Jun 12 06:19 main.css -> ..data/main.css
```


Simply put, things go well if you're inside the pod and watch for `/mnt/site` directory, but may not go well if you try to watch for a specific file `/mnt/site/main.css` since it is internally a symbolic link managed by Kubernetes.
