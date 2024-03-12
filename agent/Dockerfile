FROM gocd/gocd-agent-ubuntu-18.04:v20.4.0
ENV GO_SERVER_URL=http://gocdserver:8153/go
USER root
RUN apt-get update && \
    apt-get install -y golang && \ 
    apt-get clean
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
WORKDIR $GOPATH/src