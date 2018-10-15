FROM golang:1.11.1-stretch

RUN apt-get update -y && apt-get install -yq git bash-completion \
    vim tree \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /src/.*deb \
    && go get -d -v github.com/BurntSushi/toml \
    && go get -d -v github.com/gorilla/mux \
    && go get -d -v gopkg.in/mgo.v2 \
    && go get -d -v gopkg.in/mgo.v2/bson

ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bashrc /root/.bashrc
ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt /root/.bash_prompt


WORKDIR /go/src/server



CMD ["./server"]
