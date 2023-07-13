FROM golang:1.19

ARG DD_API_KEY
RUN test -n "$DD_API_KEY" || (echo "❗DD_API_KEY  not set" && false)
ARG DD_ENV=dev
ARG DD_SERVICE='sample-api'
ARG DD_VERSION='0.1.0'
ARG DD_HOSTNAME=local
ENV DD_API_KEY=$DD_API_KEY \
    DD_SERVICE=$DD_SERVICE \
    DD_VERSION=$DD_VERSION \
    DD_ENV=$DD_ENV \
    DD_LOGS_ENABLED=true
RUN DD_INSTALL_ONLY=true \
    DD_API_KEY=$DD_API_KEY \
    DD_SITE="datadoghq.com" \
    DD_HOSTNAME=$DD_HOSTNAME \
    bash -c "$(curl -L https://s3.amazonaws.com/dd-agent/scripts/install_script_agent7.sh)"
RUN sed -i "s/^# env:.*/env: $DD_ENV/g" /etc/datadog-agent/datadog.yaml
RUN sed -i 's/ *# logs_enabled: false/logs_enabled: true/g' /etc/datadog-agent/datadog.yaml
COPY --chown=dd-agent:dd-agent ./go.d /etc/datadog-agent/conf.d/go.d

WORKDIR /app
COPY . .
RUN go build -o main .
ENTRYPOINT ["scripts/entrypoint.sh"]
