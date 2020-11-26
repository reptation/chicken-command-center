from time import sleep

from io import BytesIO
from picamera import PiCamera

from django.http import HttpResponse, StreamingHttpResponse
# from django.shortcuts import render

def index(request):
    return HttpResponse("Hello, world. You're at the watcher index.")

def get_stream():
    my_stream = BytesIO()
    camera = PiCamera()
    camera.start_preview()
    # Camera warm-up time
    sleep(2)
    camera.capture(my_stream, 'jpeg')

def stream(request):
    return StreamingHttpResponse(get_stream())
