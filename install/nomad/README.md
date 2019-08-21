

If running on macOS, you will need to install the [Weave Net Docker Plugin](https://www.weave.works/docs/net/latest/install/plugin/plugin-v2/):

```bash
docker swarm init # if your docker host is not currently a swarm manager
docker plugin install weaveworks/net-plugin:latest_release
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