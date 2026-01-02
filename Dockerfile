FROM golang:1.17-alpine

WORKDIR /app
COPY ./ ./
ENV GOPROXY https://goproxy.cn
RUN go build -o /snow

EXPOSE 8882
CMD [ "/snow" ]