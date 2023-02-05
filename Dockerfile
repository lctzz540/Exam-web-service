FROM golang:latest
WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

RUN go build -buildvcs=false .

EXPOSE 8080

CMD [ "./Exam-web-service" ]
