ARG NGINX_VERSION=1.25.0

## 构建geoip2的模块
FROM nginx:${NGINX_VERSION} AS builder

ENV NGINX_VERSION=1.25.0 \
    GEOIP2_NAME=ngx_http_geoip2_module-3.4 \
    NGINX_PREFIX=/etc/nginx


# 安装GeoIP2相关的库
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    build-essential libpcre3-dev zlib1g-dev libperl-dev \
    # gpg gpg-agent \
    libgeoip-dev \
    libmaxminddb-dev \
    && rm -rf /var/lib/apt/lists/*


COPY ./ngx_http_geoip2_module-3.4.tar.gz /tmp/
# 下载GeoIP2相关的源码并编译安装ngx_http_geoip2_module模块

RUN cd /tmp \
    && tar xvf ngx_http_geoip2_module-3.4.tar.gz 

## Download NGINX, verify integrity and extract
RUN cd /tmp && \
    curl -O http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz && \
    # curl -O http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz.asc && \
    # gpg --verify nginx-${NGINX_VERSION}.tar.gz.asc nginx-${NGINX_VERSION}.tar.gz && \
    tar xzf nginx-${NGINX_VERSION}.tar.gz


## Compile NGINX with desired module
RUN cd /tmp/nginx-${NGINX_VERSION} && \
    ./configure --with-compat --add-dynamic-module=/tmp/ngx_http_geoip2_module-3.4 --with-stream && \
    make modules


# starting    
FROM nginx:${NGINX_VERSION}

ENV NGINX_VERSION=1.25.0 \
    NGINX_PREFIX=/etc/nginx


RUN apt-get update && apt-get install -y \
    wget \
    curl \
    libgeoip-dev \
    libmaxminddb-dev \
    && rm -rf /var/lib/apt/lists/*


COPY --from=builder /tmp/nginx-${NGINX_VERSION}/objs/ngx_http_geoip2_module.so ${NGINX_PREFIX}/modules/ngx_http_geoip2_module.so
COPY --from=builder /tmp/nginx-${NGINX_VERSION}/objs/ngx_stream_geoip2_module.so ${NGINX_PREFIX}/modules/ngx_stream_geoip2_module.so



COPY ./GeoLite2-City_20240416.tar.gz ${NGINX_PREFIX}/geoip2/GeoLite2-City_20240416.tar.gz
COPY ./GeoLite2-Country_20240416.tar.gz ${NGINX_PREFIX}/geoip2/GeoLite2-Country_20240416.tar.gz


RUN cd ${NGINX_PREFIX}/geoip2 && \
    tar xvf GeoLite2-City_20240416.tar.gz && \
    tar xvf GeoLite2-Country_20240416.tar.gz

