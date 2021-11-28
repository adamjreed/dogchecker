FROM golang:1 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build ./cmd/dogchecker

FROM public.ecr.aws/lambda/go:1
COPY --from=build /app/dogchecker ${LAMBDA_TASK_ROOT}
CMD [ "dogchecker" ]