#FROM alpine:3.11

FROM ubuntu:18.04

# base ubuntu install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates

ENV DONATION_SERVICE_CFG_PATH=/etc/csiabb DONATION_SERVICE_LOG_PATH=/var/log/csiabb DONATION_SERVICE_DATA=/opt/csiabb/data
RUN mkdir -p ${DONATION_SERVICE_CFG_PATH} ${DONATION_SERVICE_LOG_PATH} ${DONATION_SERVICE_DATA}

COPY build/bin/donation-service /usr/local/bin
COPY sampleconfig/donation-service.yaml ${DONATION_SERVICE_CFG_PATH}/

EXPOSE 8888

CMD ["donation-service"]
