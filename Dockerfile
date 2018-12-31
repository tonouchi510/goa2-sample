# swagger-uiサービスを起動するためのイメージ
FROM nginx:1.15-alpine as swagger-ui

ENV SWAGGER_URL "http://localhost:8000/v1/swagger.json"
ENV PORT 3000

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN git clone https://github.com/swagger-api/swagger-ui.git /swagger-ui
RUN cp -rf /swagger-ui/dist/* /usr/share/nginx/html/

COPY ./docker/nginx/nginx.conf /etc/nginx/
COPY ./docker/nginx/run.sh /usr/share/nginx/

RUN chmod +x /usr/share/nginx/run.sh

EXPOSE 3000

