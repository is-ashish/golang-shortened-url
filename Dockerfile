FROM golang:1.17.2

WORKDIR /home
COPY ./ /home

RUN cd /home && go build -o shortened

CMD ["/home/shortened"]