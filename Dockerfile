FROM golang:1.24

EXPOSE 3000

ARG PORT \
    REPOSITORIES_DIR \
    DB_URI \
    AUTHORIZATION_KEY

ENV LIBGIT2_VERSION=1.5.0 \
    GIT2GO_VERSION=34 \
    GIN_MODE=release \
    PORT=$PORT \
    REPOSITORIES_DIR=$REPOSITORIES_DIR \
    DB_URI=$DB_URI \
    AUTHORIZATION_KEY=$AUTHORIZATION_KEY

WORKDIR /app

# Install dependencies for libgit2
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    pkg-config \
    libssl-dev \
    libssh-dev \
    libcurl4-openssl-dev \
    libz-dev \
    libhttp-parser-dev \
    git \
    && rm -rf /var/lib/apt/lists/*

# Install libgit2
RUN git clone --branch v${LIBGIT2_VERSION} --depth 1 https://github.com/libgit2/libgit2.git /tmp/libgit2 \
    && cd /tmp/libgit2 \
    && mkdir build && cd build \
    && cmake .. -DBUILD_SHARED_LIBS=ON -DBUILD_CLAR=OFF \
    && cmake --build . --target install \
    && ldconfig \
    && rm -rf /tmp/libgit2

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install github.com/libgit2/git2go/v${GIT2GO_VERSION}

COPY . .

RUN mkdir projects/

RUN chown -R root /app/projects/

RUN git config --global --add safe.directory '*'

RUN go build -o main .

CMD ["./main"]