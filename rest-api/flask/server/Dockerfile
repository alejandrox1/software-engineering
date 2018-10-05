FROM python:3.6.6-stretch

ENV PYTHONUNBUFFERED 1

RUN apt-get update -y && apt-get install -yq git bash-completion \
    vim tree \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /src/.*deb

ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bashrc /root/.bashrc
ADD https://raw.githubusercontent.com/alejandrox1/dev_env/master/local-setup/bash_prompt /root/.bash_prompt


WORKDIR /opt/server

ADD requirements.txt /opt/server
RUN pip install -r requirements.txt


CMD ["python", "api.py"]
