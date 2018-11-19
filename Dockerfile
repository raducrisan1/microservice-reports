FROM alpine:3.8
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY ./microservice-reports /app/
WORKDIR /app
EXPOSE 3070
CMD ["./microservice-reports"]