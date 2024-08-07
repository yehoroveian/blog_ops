FROM arm64v8/golang:1.22 as build
WORKDIR /build

ARG CGO_ENABLED=0
ENV CGO_ENABLED=${CGO_ENABLED}

ARG GOOS=linux
ARG GOOS=${GOOS}

ARG GO111MODULE=on
ENV GO111MODULE=${GO111MODULE}

COPY go.mod go.sum ./
COPY . .

RUN go build -tags lambda.norpc -o deploy ./cmd/lambdas/ops/deploy/main.go

FROM public.ecr.aws/lambda/provided:al2023 AS deploy-lambda
COPY --from=build /build/deploy ./deploy
ENTRYPOINT [ "./deploy" ]