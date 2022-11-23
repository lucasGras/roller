FROM debian:stable-slim
LABEL kind="roller-test"
LABEL flavor="debian"

# Install sudo on debian to avoid unknown sudo error
RUN apt update -y
RUN apt install sudo -y

# Create test user for install scenario
RUN sudo groupadd test
RUN sudo useradd --gid test --create-home tuser
RUN usermod -aG sudo tuser
RUN echo "tuser  ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/tuser
USER tuser

# Requirements for installation
RUN sudo apt install curl -y
RUN sudo apt install systemd -y


WORKDIR /home/tuser

CMD ["/usr/sbin/init"]
# RUN ./install.sh
# RUN curl https://raw.githubusercontent.com/lucasGras/roller/main/install.sh | sh