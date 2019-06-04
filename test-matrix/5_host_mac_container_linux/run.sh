#!/bin/bash

exec  docker run --name test5  -v $(pwd):/mnt  williamyeh/fswatch  fswatch /mnt/site /mnt
