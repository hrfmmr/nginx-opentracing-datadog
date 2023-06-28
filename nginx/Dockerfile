FROM nginx:1.24.0

RUN apt-get update && \
  apt-get install -y \
    wget \
    jq

# https://docs.datadoghq.com/tracing/trace_collection/proxy_setup/?tab=nginx#module-installation
RUN get_latest_release() { \
      curl --silent "https://api.github.com/repos/$1/releases/latest" | jq --raw-output .tag_name; \
    } && \
    BASE_IMAGE=nginx:1.24.0 && \
    BASE_IMAGE_WITHOUT_COLONS=$(echo "$BASE_IMAGE" | tr ':' '_') && \
    RELEASE_TAG=$(get_latest_release DataDog/nginx-datadog) && \
    tarball="$BASE_IMAGE_WITHOUT_COLONS-ngx_http_datadog_module.so.tgz" && \
    wget "https://github.com/DataDog/nginx-datadog/releases/download/$RELEASE_TAG/$tarball" && \
    tar -xzf "$tarball" -C /usr/lib/nginx/modules && \
    rm "$tarball"
