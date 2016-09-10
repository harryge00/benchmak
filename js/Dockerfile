FROM index.caicloud.io/caicloud/node:0.12-slim

WORKDIR /app
ADD . /app

# Set the default timezone to Shanghai
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

CMD ["node", "es_query.js"]
