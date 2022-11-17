FROM debian:stable-slim
LABEL kind="roller-test"
LABEL flavor="debian"

# Install sudo on debian to avoid unknown sudo error
RUN apt update -y
RUN apt install sudo -y
RUN apt update && apt install sudo -y
RUN apt install curl -y

# Create test user for install scenario
RUN sudo groupadd test
RUN sudo useradd --gid test --create-home tuser
RUN usermod -aG sudo tuser
RUN echo "tuser  ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/tuser
USER tuser


COPY install.sh .
COPY uninstall.sh .
RUN sudo chmod +x install.sh
RUN sudo chmod +x uninstall.sh

RUN ./install.sh
# RUN curl https://raw.githubusercontent.com/lucasGras/roller/main/install.sh | sh