FROM golang:1.17 as build
RUN mkdir /hoge
COPY main.go /hoge
CMD ["go", "run", "/hoge/main.go"]