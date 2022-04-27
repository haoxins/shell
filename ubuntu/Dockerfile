FROM ubuntu:22.04

ENV SHELL /bin/bash

ARG KUBECTL_ARCH=amd64
ARG KUBECTL_VERSION=v1.23.5
ARG KIND_VERSION=v0.12.0

SHELL ["/bin/bash", "-c"]

RUN export DEBIAN_FRONTEND=noninteractive \
  && apt-get -yq update \
  && apt-get -yq upgrade \
  && apt-get -yq install \
    bash \
    curl \
    sudo \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

# Install - kind
RUN curl -Lo ./kind https://kind.sigs.k8s.io/dl/${KIND_VERSION}/kind-linux-amd64 \
  && chmod +x ./kind \
  && mv ./kind /usr/local/bin/kind

# Install - kubectl
RUN curl -sL "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/${KUBECTL_ARCH}/kubectl" -o /usr/local/bin/kubectl \
  && curl -sL "https://dl.k8s.io/${KUBECTL_VERSION}/bin/linux/${KUBECTL_ARCH}/kubectl.sha256" -o /tmp/kubectl.sha256 \
  && echo "$(cat /tmp/kubectl.sha256) /usr/local/bin/kubectl" | sha256sum --check \
  && rm /tmp/kubectl.sha256 \
  && chmod +x /usr/local/bin/kubectl
