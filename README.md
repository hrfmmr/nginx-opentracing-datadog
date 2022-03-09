# NGINX Opentracing and APM demo

- App is borrowed from Docker tutorial
  - https://docs.docker.com/get-started/02_our_app/
  - https://docs.docker.com/get-started/07_multi_container/

- Opentracing settings are borrowed from Datadog opentracing plugin
  - https://github.com/DataDog/dd-opentracing-cpp/tree/master/examples/nginx-tracing

- How to run
  ```
  export DD_API_KEY="YourDatadogApiKey"
  docker-compose build
  docker-compose up
  ```
  - Open http://localhost:8888