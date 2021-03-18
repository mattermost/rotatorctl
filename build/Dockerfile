# Build the mattermost eks rotator
ARG DOCKER_BUILD_IMAGE=golang:1.15.8
ARG DOCKER_BASE_IMAGE=alpine:3.13

FROM ${DOCKER_BUILD_IMAGE} AS build
WORKDIR /rotatorctl/
COPY . /rotatorctl/
RUN apt-get update -yq && apt-get install -yq unzip
RUN make build


# Final Image
FROM ${DOCKER_BASE_IMAGE}
LABEL name="Mattermost Rotatorctl" \
  maintainer="cloud-team@mattermost.com" \
  vendor="Mattermost" \
  distribution-scope="public" \
  architecture="x86_64" \
  url="https://mattermost.com" \
  io.k8s.description="Mattermost Rotatorctl accelerates k8s cluster node rotations" \
  io.k8s.display-name="Mattermost Rotatorctl"

ENV ROTATORCTL=/rotatorctl/rotatorctl \
    USER_UID=10001 \
    USER_NAME=rotatorctl

RUN  apk update && apk add libc6-compat && apk add ca-certificates
COPY --from=build /rotatorctl/build/_output/bin/rotatorctl /rotatorctl/rotatorctl
COPY --from=build /rotatorctl/build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup
WORKDIR /rotatorctl/

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
