ARG GO_VERSION=1.25

FROM golang:${GO_VERSION}

WORKDIR /nyg-backend

COPY . /nyg-backend

COPY docker-runthis.sh /nyg-backend/docker-runthis.sh

RUN mkdir -p /nyg-backend/bin

RUN ls -l /nyg-backend

RUN for dir in collection credentials game graphql-chat-app msg-webhook patch post profile protos SQL-Manager updates; do \
    dirname=$(basename "$dirname"); \    
    if [ "$dirname" != "SQL-Manager" ] && [ -f "/nyg-backend/$dir/go.mod" ]; then \
    echo "Building $dir"; \
    cd "/nyg-backend/$dir" && go mod tidy && go build -o /nyg-backend/bin/$dir main.go && cd /nyg-backend; \
    else \
    echo "No go.mod in $dir, skipping build"; \
    fi \
    done

# CMD ["/bin/sh", "-c", "for binary in /nyg-backend/bin/*; do echo \"Running $binary:\"; $binary; done"]
RUN chmod +x /nyg-backend/docker-runthis.sh

CMD ["/nyg-backend/docker-runthis.sh"]
