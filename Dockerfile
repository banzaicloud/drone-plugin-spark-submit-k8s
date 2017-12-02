FROM banzaicloud/spark-base:v2.2.0-k8s-1.0.179
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD spark-submit-k8s /bin/
ENTRYPOINT ["/bin/spark-submit-k8s"]