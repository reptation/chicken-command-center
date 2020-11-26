from time import sleep

from io import BytesIO
from picamera import PiCamera

from django.http import HttpResponse
# from django.shortcuts import render

def index(request):
    return HttpResponse("Hello, world. You're at the watcher index.")

def stream(request):
    # Explicitly open a new file called my_image.jpg
    stream = BytesIO()
    camera = PiCamera()
    camera.resolution = (640, 480)
    camera.start_recording(stream, format='h264', quality=23)
    camera.wait_recording(15)

    sleep(2)
    camera.capture(my_file)
    # At this point my_file.flush() has been called, but the file has
    # not yet been closed
    my_file.close()
