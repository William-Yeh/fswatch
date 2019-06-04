@ECHO OFF

set IMAGE=williamyeh/fswatch:1.0-windowsnanoserver-1809
docker run --name test7  -v %cd%:C:\mnt  %IMAGE%  fswatch \mnt\site \mnt