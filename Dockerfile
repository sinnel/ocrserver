FROM golang:1.14
WORKDIR /go/src/ocrserver

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -u github.com/gin-gonic/gin

RUN sed -i 's/http:\/\/archive\.ubuntu\.com\/ubuntu\//http:\/\/mirrors\.163\.com\/ubuntu\//g' /etc/apt/sources.list
RUN apt update \
    && apt install -y \
      ca-certificates \
      libtesseract-dev \
      tesseract-ocr
RUN go get -t github.com/otiai10/gosseract/v2

ADD . /go/src/ocrserver
EXPOSE 8080

CMD go run main.go