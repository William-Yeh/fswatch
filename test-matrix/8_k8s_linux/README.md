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
