FROM ubuntu:latest
LABEL authors="mj"

ENTRYPOINT ["top", "-b"]