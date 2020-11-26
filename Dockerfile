FROM python:3.9

COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Mounts the application code to the image
COPY . app
WORKDIR /app

EXPOSE 8000

# runs the production server
ENTRYPOINT ["python", "chickencommandcenter/manage.py"]
CMD ["runserver", "0.0.0.0:8000"]

