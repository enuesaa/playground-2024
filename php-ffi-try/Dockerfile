FROM --platform=linux/arm64 php:8.3

# ffi
# php で ffi は ext として提供されている
RUN apt update && \
    apt install libffi-dev
RUN docker-php-ext-configure ffi --with-ffi \
    && docker-php-ext-install ffi

# install go
RUN curl -LO https://go.dev/dl/go1.23.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
ENV PATH="$PATH:/usr/local/go/bin"

WORKDIR /workspace
