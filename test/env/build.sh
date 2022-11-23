PODNAME='roller-testenv-pod'

# Debian variables
DEBIAN_IMAGE_NAME='roller-testenv-debian'
DEBIAN_CONTAINER_NAME='rtdebian'

setup_pod() {
    podman pod create --name $PODNAME
}

teardown_pod() {
    podman pod rm -fi $PODNAME
}

install_scenario_debian() {
  printf "Update scripts..."
  mkdir -p debian/resources
  rm debian/resources/*
  cp ../../scripts/ debian/resources -r
  cp ../../scripts/ debian/resources -r
  printf "done.\n------"

  printf "Cleaning containers"
  podman container rm -f $DEBIAN_CONTAINER_NAME
  printf "done.\n------"

  printf "Creating fresh containers"
  podman build debian --force-rm -t $DEBIAN_IMAGE_NAME
  # use volumes to copy install / uninstall scripts
  podman run -v $(pwd)/debian/resources:/home/tuser -dt --pod $PODNAME --name $DEBIAN_CONTAINER_NAME roller/debian
  # Run pre script and install script in container
  podman exec -it --user tuser rtdebian bash -c "./home/tuser/scripts/pre/testenv-pre-install.sh"
  podman exec -it --user tuser rtdebian bash -c "sudo ./home/tuser/scripts/install.sh"

  if [ "$?" != 0 ]; then
    printf "Debian installation scenario failed"
    exit 1
  fi
  printf "Debian installation scenario done.\n------"
}

# Init pod
setup_pod;
# Run installation scenarios
install_scenario_debian;
# Teardown pod
teardown_pod;
exit 0