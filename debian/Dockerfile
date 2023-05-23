FROM debian

ENV SHELL /bin/bash

ARG KUBECTL_ARCH=amd64
ARG KUBECTL_VERSION=v1.27.1
ARG KIND_VERSION=v0.18.0

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

# Using Rust nightly
RUN curl -sSf https://sh.rustup.rs | sh -s -- -y \
  && ~/.cargo/bin/rustup install nightly \
  && ~/.cargo/bin/rustup default nightly

ENV PATH="$HOME/.cargo/bin:$PATH"

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

# Install - kustomize
RUN curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash \
  && chmod +x ./kustomize \
  && mv ./kustomize /usr/local/bin/kustomize
