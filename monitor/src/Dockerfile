FROM 1.9.3-alpine3.7 as builder

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# setting working directory
WORKDIR /go/src/app

# installing dependencies
RUN go get github.com/leonelquinteros/gorand
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/sirupsen/logrus
#RUN go get gopkg.in/yaml.v2
RUN go get github.com/ghodss/yaml

COPY / /go/src/app/
RUN go build -o app

FROM alpine:3.7

RUN apk update && apk upgrade && \
    apk add --no-cache bash openssh curl

WORKDIR /go/src/app
COPY --from=builder /go/src/app/app /go/src/app/app

EXPOSE 10010

CMD ["./app"]