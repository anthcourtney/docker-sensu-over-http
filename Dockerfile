FROM centos:7

RUN yum -y install ntp && \
    yum -y install ruby golang rubygems && \
    yum -y install python-setuptools git && \
    easy_install supervisor
RUN gem install sensu-plugins-ntp --no-rdoc --no-ri

COPY etc/supervisord.conf /etc/supervisord.conf

ADD go/ /go
RUN export GOPATH=/go && \
    go get github.com/gorilla/mux && \
    go install sensu-over-http

CMD ["/usr/bin/supervisord"]
