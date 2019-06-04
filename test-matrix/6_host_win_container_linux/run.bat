@ECHO OFF

docker run --name test6  -v %cd%:/mnt  williamyeh/fswatch  fswatch /mnt/site /mnt