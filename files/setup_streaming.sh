mkdir -p /dev/shm/streaming
if [ ! -e /var/www/html/streaming ]; then
    ln -s  /dev/shm/streaming /var/www/html/streaming
fi 

ln -sf /home/pi/streaming/index.html /var/www/html/streaming/index.html
