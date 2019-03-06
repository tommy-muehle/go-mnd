FROM golang:1.12-alpine

LABEL "com.github.actions.icon"="code"
LABEL "com.github.actions.color"="green-dark"
LABEL "com.github.actions.name"="Go Tools"
LABEL "com.github.actions.description"="This is an action to run different go commands."

ENV CGO_ENABLED=0

COPY ./entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

