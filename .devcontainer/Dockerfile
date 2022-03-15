FROM fedora:latest

RUN dnf install -y dnf-plugins-core && \
    dnf config-manager --add-repo https://rpm.releases.hashicorp.com/fedora/hashicorp.repo && \
    dnf install -y \
      git \
      golang \
      packer


