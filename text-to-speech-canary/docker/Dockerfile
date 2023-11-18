FROM amd64/python:latest

WORKDIR /app

RUN mkdir -p /gc-service-account/creds
RUN chmod 771 /gc-service-account/creds

COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE 1323
ENTRYPOINT ["python3", "main.py"]