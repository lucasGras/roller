#!/usr/bin/env bash

install_caddy() {
    which caddy > /dev/null
    if [ "$?" = 0 ]; then
        echo "Caddy already installed, skip."
        return 0
    fi
    # Webi community-maintained installation method
    # Handle linux and macos and avoid tricky os gestion
    curl -sS https://webi.sh/caddy | sh
    if [ "$?" = 0 ]; then
        echo "Caddy installation failed."
        exit 1
    fi

    # Will see how to handle sudo
    sudo mv caddy /usr/bin/
    caddy version
    if [ "$?" = 0 ]; then
        echo "Caddy binary not detected."
        exit 1
    fi
    sudo groupadd --system caddy
    sudo useradd --system \
        --gid caddy \
        --create-home \
        --home-dir /var/lib/caddy \
        --shell /usr/sbin/nologin \
        --comment "Caddy web server" \
        caddy
}

install_debian_daemon() {
    # Use the caddy one or my own service ? Want to avoid hard maintaining stuff..
    wget https://github.com/caddyserver/dist/blob/master/init/caddy-api.service > caddy-api.service
    if [ "$?" = 0 ]; then
        echo "Caddy systemd service not reachable."
        exit 1
    fi
    sudo mv caddy-api.service /etc/systemd/system/
    sudo systemctl daemon-reload
    sudo systemctl enable --now caddy
    systemctl status caddy
    if [ "$?" = 0 ]; then
        echo "Caddy systemd service failed to start."
        exit 1
    fi
}

install_roller() {
    mdkir $HOME/.roller
    touch $HOME/.roller/rollerFile
    echo '{"projects": []}' > $HOME/.roller/rollerFile
    # curl from github
}

if [ "$(uname)" == "Darwin" ]; then
    # Do something under Mac OS X platform
    echo "Platform Darwin supported but not advised !";
    install_roller;
    install_caddy;
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    # Do something under GNU/Linux platform
    echo "Platform Linux supported";
    install_roller;
    install_caddy;

    # Check distrib for linux compatibility
    install_debian_daemon;
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
    # Do something under 32 bits Windows NT platform
    echo "Platform 32 bits Windows not supported" && exit 1;
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
    # Do something under 64 bits Windows NT platform
    echo "Platform 64 bits Windows not supported" && exit 1;
fi

# For linux: systemctl
# For mac: launchctl