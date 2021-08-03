FROM golang:1.16-alpine AS builder

ARG BITBUCKET_USERNAME
ARG BITBUCKET_APP_PASSWORD
ARG COMPANY_NAME

RUN apk update && apk add git mercurial ca-certificates && rm -rf /var/cache/apk/*
RUN adduser -D -g '' appuser

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

RUN mkdir -p /app
RUN grep appuser /etc/passwd > /passwd1
COPY . /app
WORKDIR /app

RUN git config --global url."https://${BITBUCKET_USERNAME}:${BITBUCKET_APP_PASSWORD}@bitbucket.org/${COMPANY_NAME}/".insteadOf "https://bitbucket.org/${COMPANY_NAME}/"
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY="direct"
RUN go env -w GOPRIVATE="bitbucket.org/${COMPANY_NAME}"
RUN go mod tidy

WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /skeleton-app
RUN chmod +x /skeleton-app

FROM scratch
COPY --from=builder /passwd1 /etc/passwd
COPY --from=builder /skeleton-app /skeleton-app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER appuser
ENTRYPOINT ["/skeleton-app", "server-run"]