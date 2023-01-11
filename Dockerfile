FROM golang:1.17-alpine
WORKDIR app
COPY . .
CMD ["tail", "-f", "/dev/null"]