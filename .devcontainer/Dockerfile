FROM eu.gcr.io/cts-public-images-1/cts-base

ENV ZSH_SUGGESTIONS_REPO=https://github.com/zsh-users/zsh-autosuggestions
ENV ZSH_HIGHLIGHTING_REPO=https://github.com/zsh-users/zsh-syntax-highlighting.git
ENV GO_VERSION=1.18.2

RUN apt update && export DEBIAN_FRONTEND=noninteractive && \
  apt upgrade -y && \
  apt -y install --no-install-recommends \
  zsh \
  && apt clean -y && rm -rf /var/lib/apt/lists/*

# Install zsh syntax highlighting and auto-completion
RUN git clone ${ZSH_SUGGESTIONS_REPO} /home/vscode/.oh-my-zsh/custom/plugins/zsh-autosuggestions && \
    git clone ${ZSH_HIGHLIGHTING_REPO} /home/vscode/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting

RUN sed -i "s/plugins=(git)/plugins=(git zsh-autosuggestions zsh-syntax-highlighting)/g" /home/vscode/.zshrc

# Install Go
RUN wget -c https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-arm64.tar.gz && \
    rm -f https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV SHELL=/usr/bin/zsh
