# file: .devcontainer/Dockerfile
# Development container for Subtitle Manager
# Optimized for development workflow with hot reloading and dev tools

FROM golang:1.24-bookworm

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure apt and install packages
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release \
    #
    # Install C++ tools for CGO
    && apt-get -y install build-essential cmake cppcheck valgrind clang lldb llvm gdb \
    #
    # Install SQLite for database operations
    && apt-get install -y sqlite3 libsqlite3-dev \
    #
    # Install FFmpeg for subtitle processing
    && apt-get install -y ffmpeg \
    #
    # Install additional development tools
    && apt-get install -y curl wget unzip zip jq \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*



# Note: Go tools will be installed automatically by VS Code Go extension when needed

# Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Create workspace directory before switching to non-root user
RUN mkdir -p /workspace

# Set the default user
USER $USERNAME

WORKDIR /workspace

# Set environment variables for development
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV SM_FFMPEG_PATH=/usr/bin/ffmpeg
ENV GOTOOLCHAIN=local

# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=dialog