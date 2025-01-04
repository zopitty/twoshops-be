FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD twoshops-be /usr/bin/twoshops-be

CMD ["twoshops-be"]