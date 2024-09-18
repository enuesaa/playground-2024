FROM python:3.11

RUN pip install poetry

WORKDIR /workspace

COPY . .

RUN poetry install

# CMD poetry run start
