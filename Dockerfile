   
FROM golang:alpine as build 

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go_services_info


FROM scratch as production

COPY --from=build /app/go_services_info .
# COPY --from=build /app/config ./config

# RUN apk add tzdata

ENV TZ=Asia/Bangkok

CMD ["/go_services_info"]