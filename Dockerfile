# Build from golang base image
FROM golang:1.10

LABEL authors="Calvin Feng"

COPY . /go/src/go-academy/tensor_go
WORKDIR /go/src/go-academy/tensor_go

EXPOSE 3000

ENV TF_TYPE "cpu"
ENV TARGET_DIRECTORY "/usr/local"

# Install TensorFlow C library
RUN curl -L \
   https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-linux-x86_64-1.6.0.tar.gz | \
   tar -C ${TARGET_DIRECTORY} -xz
RUN ldconfig

# Hide some warnings
ENV TF_CPP_MIN_LOG_LEVEL 2

RUN go get -u github.com/golang/dep/cmd/dep
RUN go install

# Install Node
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash - && apt-get install -y nodejs
RUN npm install
RUN npm run build

CMD tensor_go server