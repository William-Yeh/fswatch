#!/bin/bash

exec  docker run --name test4  -v $(pwd):/mnt  williamyeh/fswatch  fswatch /mnt/site /mnt