# fswatch - Watch for changes in file system

## Purpose

Watch for any changes recursively in specific directories.

Docker image: [williamyeh/fswatch](https://hub.docker.com/r/williamyeh/fswatch/)


## Binary

Executable binaries for 3 platforms are provided in this image.  You can extract them under Linux or Mac OS X:

```
% docker run  -v $(pwd):/mnt  williamyeh/fswatch  ./copy-files.sh
```

or under Windows:

```
C:> docker run  -v %cd%:/mnt  williamyeh/fswatch  ./copy-files.sh
```


## Usage

```
fswatch - Watch for changes in file system.

Usage:
  fswatch <dir>...
```


## Reference

- [fsnotify](https://fsnotify.org/): Cross-platform file system notifications for Go.

- [How to detect file changes in Golang](https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f)


## LICENSE

Apache License 2.0.  See the [LICENSE](LICENSE) file.
