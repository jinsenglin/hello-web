FROM golang:1.17.11-alpine
WORKDIR /api
ADD . /api
RUN cd /api && go build
EXPOSE 80
ENTRYPOINT ./api