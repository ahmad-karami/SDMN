# HTTP Status API in Docker

This project is a small HTTP status API built on Python's `http.server` module and packaged as a Docker image. `Server.py` serves one endpoint, `/api/v1/status`, and keeps a status that is either `OK` or `not OK`. A GET request reports the current status as JSON, and a POST of `{"status": "not OK"}` flips it. The `Dockerfile` builds a `python:3.10.6` image that runs the server on port 9999.

## How it works
The `Dockerfile` copies the folder into `/code`, exposes port 9999, and runs `python Server.py` when the container starts. The server listens on `172.17.0.2:9999`, the address Docker usually gives the first container on its default bridge network. A global `toggle` holds the status. GET on `/api/v1/status` returns 200 with `OK` while `toggle` is false and `not OK` while it is true. A valid POST returns 201 with `not OK` and flips `toggle`, so a second POST sets the status back to `OK`. Requests to other paths and POSTs with any other body return 400.
