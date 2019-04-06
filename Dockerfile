# viron用イメージ
FROM node:9 as viron

# Setup project
RUN git clone https://github.com/cam-inc/viron.git /viron
RUN sed -i "s|ssl: true,|ssl: false,|g" /viron/rollup.local.config.js

RUN chown -R node:node /viron
ENV HOME /viron
USER node
WORKDIR $HOME

RUN npm install

EXPOSE 8080
USER root
CMD npm start