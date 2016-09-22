# Caicloud executor dockerfile.
FROM index.caicloud.io/caicloud/debian:jessie

# Set the default timezone to Shanghai
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

RUN mkdir /app
WORKDIR /app
Add . /app
CMD ./k8s-service-stress-test
