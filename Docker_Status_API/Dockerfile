FROM python:3.10.6

ENV PYTHONUNBUFFERED 1

WORKDIR /code

COPY . /code/

EXPOSE 9999

# Start the HTTP server when the container starts.
CMD [ "python", "Server.py" ]