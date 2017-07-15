FROM alpine

RUN apk update && apk add flite
RUN mkdir /app
ADD . /app/
WORKDIR /app
CMD ["/app/main"]
