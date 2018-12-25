FROM golang:1.11.4 as builder
WORKDIR /go/src/github.com/dechiad1/danny-ac/
COPY main.go .
RUN go get "k8s.io/api/admission/v1beta1" && go get "k8s.io/api/core/v1" && go get "k8s.io/apimachinery/pkg/apis/meta/v1"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/dechiad1/danny-ac/app .
CMD ["./app"]
