uninstall_caddy() {
  sudo rm /etc/apt/sources.list.d/caddy-stable.list
  sudo apt remove caddy
}

uninstall_debian_daemon() {
  sudo systemctl disable --now caddy
  sudo rm -f /etc/systemd/system/caddy-api.service
  sudo systemctl daemon-reload
}

uninstall_roller() {
  rm -rf $HOME/.roller
}

uninstall_roller;