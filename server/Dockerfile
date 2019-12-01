FROM golang:1.13.3-alpine3.10

# Set ARGS
ARG SERVER_DIR

# Create app directory
RUN mkdir -p ${SERVER_DIR}

# Set working directory in the image
WORKDIR ${SERVER_DIR}

# install git, bash and dep
RUN apk add --update --no-cache git && \
    apk add --no-cache bash && \
    go get -u github.com/golang/dep/cmd/dep

# COPY files from host to the present location (.) in your image
COPY . .

# Run dep ensure to install all dependencies specified in Gopkg
RUN dep ensure

# RUN cd ${SERVER_DIR} && \
#     go build && \
#     chmod u+x server

# ENTRYPOINT ["tail", "-f", "/dev/null"]