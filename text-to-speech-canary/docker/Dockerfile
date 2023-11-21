FROM bitnami/python:3.10

RUN mkdir -p /gcsa
RUN chmod 771 /gcsa
WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE 1323
CMD ["python3", "main.py"]