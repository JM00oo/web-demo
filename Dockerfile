FROM golang:1.10.1-alpine3.7

USER root
RUN apk update && apk upgrade && apk add --no-cache bash git openssh && mkdir /root/.ssh
ADD id_rsa /root/.ssh/id_rsa
ADD config.go /tmp/config.go
RUN go get -u github.com/kardianos/govendor
RUN echo $GOPATH
RUN echo "github.com,192.30.253.113 ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==" >> /root/.ssh/known_hosts
RUN cd $GOPATH/src && mkdir -p github.com && cd github.com && git clone git@github.com:JM00oo/web-demo.git && cd web-demo && govendor sync
RUN cp /tmp/config.go /go/src/github.com/web-demo/config

CMD cd $GOPATH/src/github.com/web-demo/ && go run main.go
