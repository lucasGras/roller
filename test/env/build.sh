# Clean docker volumes, we want to test from clean envs
docker container rm roller-testenv-debian -fv

cp ../../scripts/install.sh debian/resources
cp ../../scripts/uninstall.sh debian/resources
docker build -t roller-testenv-debian debian

docker run -it --name roller-testenv-debian roller-testenv-debian