FROM ubuntu:23.10

ENV SHELL /bin/bash

ARG KUBECTL_ARCH=amd64
ARG KUBECTL_VERSION=v1.28.2

SHELL ["/bin/bash", "-c"]

RUN export DEBIAN_FRONTEND=noninteractive \
  && apt-get -yq update \
  && apt-get -yq upgrade \
  && apt-get -yq install \
    bash \
    curl \
    golang \
    make \
    python3 \
    sudo \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

RUN curl -sSf https://sh.rustup.rs | sh -s -- -y \
  && ~/.cargo/bin/rustup install stable \
  && ~/.cargo/bin/rustup update stable

ENV PATH="$HOME/.cargo/bin:$PATH"

# Install - kubectl
RUN curl -sL "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/${KUBECTL_ARCH}/kubectl" -o /usr/local/bin/kubectl \
  && curl -sL "https://dl.k8s.io/${KUBECTL_VERSION}/bin/linux/${KUBECTL_ARCH}/kubectl.sha256" -o /tmp/kubectl.sha256 \
  && echo "$(cat /tmp/kubectl.sha256) /usr/local/bin/kubectl" | sha256sum --check \
  && rm /tmp/kubectl.sha256 \
  && chmod +x /usr/local/bin/kubectl

# Install - kustomize
RUN curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash \
  && chmod +x ./kustomize \
  && mv ./kustomize /usr/local/bin/kustomize
