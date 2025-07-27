FROM golang:1.23

WORKDIR /app

COPY . .

EXPOSE 8000

CMD ["go", "run", "."]