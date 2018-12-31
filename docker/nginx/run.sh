#! /bin/sh

set -e
NGINX_ROOT=/usr/share/nginx/html
INDEX_FILE=$NGINX_ROOT/index.html

if [[ -n "${SWAGGER_URL}" ]]; then
  sed -i "s|https://petstore.swagger.io/v2/swagger.json|${SWAGGER_URL}|g" $INDEX_FILE
  sed -i "s|http://example.com/api|${SWAGGER_URL}|g" $INDEX_FILE
fi

# replace the PORT that nginx listens on if PORT is supplied
if [[ -n "${PORT}" ]]; then
    sed -i "s|8080|${PORT}|g" /etc/nginx/nginx.conf
fi

exec nginx -g 'daemon off;'