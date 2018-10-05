FROM golang:1.0.4-stretch

RUN apt-get update -y && apt-get install -yq git bash-completion \
    vim tree \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /src/.*deb

ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bashrc /root/.bashrc
ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt /root/.bash_prompt


WORKDIR /go/src/server



CMD ["python", "api.py"]
