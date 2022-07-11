FROM python:3-alpine
WORKDIR /api
ADD . /api
RUN cd /api && pip install -r requirements.txt
EXPOSE 80
ENTRYPOINT ["python3", "/api/main.py"]