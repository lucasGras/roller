PODNAME='roller-testenv-pod'
DEBIAN_IMAGE_NAME='roller-testenv-debian'
DEBIAN_CONTAINER_NAME='rtdebian'

cp ../../scripts/install.sh debian/
cp ../../scripts/uninstall.sh debian/


podman container rm $DEBIAN_CONTAINER_NAME
podman pod rm -fi $PODNAME

podman pod create --name $PODNAME

podman build debian --no-cache -t $DEBIAN_IMAGE_NAME

podman run -dt --pod roller-testenv-pod --name $DEBIAN_CONTAINER_NAME roller/debian
podman attach $DEBIAN_CONTAINER_NAME