# Minikube in Ubuntu 22.04

## Install Minikube

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
```

## Install podman

```bash
OS=xUbuntu_22.04
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/unstable/$OS/ /" | tee /etc/apt/sources.list.d/devel:kubic:libcontainers:unstable.list

curl -fsSL https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/unstable/$OS/Release.key | gpg --dearmor | tee /etc/apt/trusted.gpg.d/devel_kubic_libcontainers_unstable.gpg > /dev/null

apt update
apt install podman -y
```

## Install cri-o

```bash
VERSION=1.24
OS=xUbuntu_22.04
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/$OS/ /" > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
echo "deb http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/$VERSION/$OS/ /" > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable:cri-o:$VERSION.list

curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/$VERSION/$OS/Release.key | apt-key add -
curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/$OS/Release.key | apt-key add -

apt-get update
apt-get install cri-o cri-o-runc
```

## Start minikube

```bash
minikube start --driver=podman --container-runtime=cri-o
```
