FROM ubuntu:latest

RUN apt update \
    && apt upgrade \
    && apt install -y ca-certificates \
    && apt install -y golang make gcc git sudo \
    && echo "export PATH=${HOME}/go/bin:\${PATH}" >> /etc/profile

WORKDIR /build

CMD . /etc/profile && echo "PATH: ${PATH}" && make