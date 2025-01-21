#!/bin/bash

read -n 1 -p "Would you want to export the image file? (Y/n) " ans;
ans=${ans:-Y}

DOCKERHUB_REPO="giterhub/electerm-sync-server"

echo ""
echo "Repo: $DOCKERHUB_REPO"

docker build -t "electerm-base:latest" -f docker/Dockerfile.base .
docker build -t "$DOCKERHUB_REPO:latest" --build-arg BASE_IMAGE="electerm-base:latest" -f docker/Dockerfile.amd64 .

docker images | grep "$DOCKERHUB_REPO"

case $ans in
  y|Y)
    echo "export the image file..."
    docker image save $DOCKERHUB_REPO:latest | gzip > electerm-sync-server.tar.gz
    echo "export done."
    ;;
esac
