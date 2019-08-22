# Installing on Nomad

Installation on Nomad requires the following:
- [Levant](https://github.com/jrasell/levant) installed on your local machine
- [Docker](https://github.com/jrasell/levant), [Consul](https://www.consul.io), [Vault](https://www.vaultproject.io), and [Nomad](https://www.nomadproject.io/) installed on the target host machine (which can also be your local machine)

If running on macOS, you will need to install the [Weave Net Docker Plugin](https://www.weave.works/docs/net/latest/install/plugin/plugin-v2/):

```bash
docker swarm init # if your docker host is not currently a swarm manager
docker plugin install weaveworks/net-plugin:latest_release --grant-all-permissions
docker plugin enable weaveworks/net-plugin:latest_release
docker network create --driver=weaveworks/net-plugin:latest_release --attachable weave

```


Note: This nomad job is experimental and designed to be used with a
specific vault+consul+nomad setup.

See `launch-consul-vault-nomad-dev.sh` to see how we run Nomad in a way
that supports `install.nomad`

See `get-gateway-url.sh` to see how we get the URL of the Gateway container
that's running inside of docker (through nomad).

To install:

`nomad run install.nomad`

## Creating a VM on GCE

Create an instance with all requirements installed:

```bash
export INSTANCE_NAME=nomad-gloo-e2e
export INSTANCE_ZONE=us-central1-a

gcloud compute instances create ${INSTANCE_NAME}  \
    --project solo-public \
    --zone ${INSTANCE_NAME} \
    --image-family ubuntu-1904 \
    --image-project ubuntu-os-cloud \
    --metadata-from-file startup-script=prepare-gce-cluster.sh \
    --tags ${INSTANCE_NAME}

gcloud compute firewall-rules create ${INSTANCE_NAME} \
    --project solo-public \
    --target-tags ${INSTANCE_NAME} \
    --allow tcp:4646,tcp:8500,tcp:8200,tcp:8080,tcp:8443

```
