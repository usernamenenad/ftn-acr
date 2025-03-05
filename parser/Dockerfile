FROM python:3.13-alpine3.20

WORKDIR /app/

COPY requirements.txt /app/

RUN pip install -r requirements.txt

COPY . /app/

CMD [ "python", "main.py" ]