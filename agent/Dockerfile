FROM gocd/gocd-agent-ubuntu-18.04:v20.4.0
ENV GO_SERVER_URL=http://gocdserver:8153/go
USER root
RUN apt-get update && \
    apt-get install -y golang && \
    apt-get install -y python3.8 python3-pip && \ 
    apt-get clean 
COPY ./requirements.txt /requirements.txt
RUN python3.8 -m pip install --upgrade pip setuptools wheel
RUN python3.8 -m pip install -r requirements.txt
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
WORKDIR $GOPATH/src
EXPOSE 8001