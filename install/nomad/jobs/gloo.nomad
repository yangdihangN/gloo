job "gloo" {

  datacenters = ["[[.datacenter]]"]
  region      = "[[.region]]"
  type        = "service"

  update {
    max_parallel = 1
    min_healthy_time = "10s"
    healthy_deadline = "3m"
    auto_revert = false
    canary = 0
  }

  migrate {
    max_parallel = 1
    health_check = "checks"
    min_healthy_time = "10s"
    healthy_deadline = "5m"
  }


  group "gloo" {
    count = [[.gloo.replicas]]

    task "gloo" {
      driver = "docker"
      config {
        image = "[[.gloo.image.registry]]/[[.gloo.image.repository]]:[[.gloo.image.tag]]"
        port_map {
          xds = [[.gloo.xdsPort]]
        }
        args = [
          "--namespace=[[.global.namespace]]",
          "--dir=${NOMAD_TASK_DIR}/settings",
        ]
      }

      template {
        data = <<EOF
bindAddr: 0.0.0.0:[[.gloo.xdsPort]]
consul:
  address: [[.consul.address]]
  serviceDiscovery: {}
consulKvSource: {}
directoryArtifactSource:
  directory: /data
discoveryNamespace: [[.global.namespace]]
metadata:
  name: default
  namespace: [[.global.namespace]]
refreshRate: [[.global.refreshRate]]
vaultSecretSource:
  address: [[.vault.address]]
  token: [[.vault.token]]
EOF

        destination = "${NOMAD_TASK_DIR}/settings/[[.global.namespace]]/default.yaml"
      }

      resources {
        # cpu required in MHz
        cpu = [[.gloo.cpuLimit]]

        # memory required in MB
        memory = [[.gloo.memLimit]]

        network {
          # bandwidth required in MBits
          mbits = [[.gloo.bandwidthLimit]]
          port "xds" {}
        }
      }

      service {
        name = "gloo-xds"
        tags = ["gloo", "xds", "grpc"]
        port = "xds"
        check {
          name = "alive"
          type = "tcp"
          interval = "10s"
          timeout = "2s"
        }
      }

      vault {
        change_mode = "restart"
        policies = ["gloo"]
      }
    }

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }
  }


  group "discovery" {
    count = [[.discovery.replicas]]

    task "discovery" {
      driver = "docker"
      config {
        image = "[[.discovery.image.registry]]/[[.discovery.image.repository]]:[[.discovery.image.tag]]"
        args = [
          "--namespace=[[.global.namespace]]",
          "--dir=${NOMAD_TASK_DIR}/settings/",
        ]
      }

      template {
        data = <<EOF
bindAddr: 0.0.0.0:[[.gloo.xdsPort]]
consul:
  address: [[.consul.address]]
  serviceDiscovery: {}
consulKvSource: {}
directoryArtifactSource:
  directory: /data
discoveryNamespace: [[.global.namespace]]
metadata:
  name: default
  namespace: [[.global.namespace]]
refreshRate: [[.global.refreshRate]]
vaultSecretSource:
  address: [[.vault.address]]
  token: [[.vault.token]]
EOF

        destination = "${NOMAD_TASK_DIR}/settings/[[.global.namespace]]/default.yaml"
      }

      resources {
        # cpu required in MHz
        cpu = [[.discovery.cpuLimit]]

        # memory required in MB
        memory = [[.discovery.memLimit]]

        network {
          # bandwidth required in MBits
          mbits = [[.discovery.bandwidthLimit]]
        }
      }

      vault {
        change_mode = "restart"
        policies = ["gloo"]
      }
    }

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }
  }


  group "gateway" {
    count = [[.gateway.replicas]]

    task "gateway" {
      driver = "docker"
      config {
        image = "[[.gateway.image.registry]]/[[.gateway.image.repository]]:[[.gateway.image.tag]]"
        args = [
          "--namespace=[[.global.namespace]]",
          "--dir=${NOMAD_TASK_DIR}/settings/",
        ]
      }

      template {
        data = <<EOF
bindAddr: 0.0.0.0:[[.gloo.xdsPort]]
consul:
  address: [[.consul.address]]
  serviceDiscovery: {}
consulKvSource: {}
directoryArtifactSource:
  directory: /data
discoveryNamespace: [[.global.namespace]]
metadata:
  name: default
  namespace: [[.global.namespace]]
refreshRate: [[.global.refreshRate]]
vaultSecretSource:
  address: [[.vault.address]]
  token: [[.vault.token]]
EOF

        destination = "${NOMAD_TASK_DIR}/settings/[[.global.namespace]]/default.yaml"
      }

      resources {
        # cpu required in MHz
        cpu = [[.gateway.cpuLimit]]

        # memory required in MB
        memory = [[.gateway.memLimit]]

        network {
          # bandwidth required in MBits
          mbits = [[.gateway.bandwidthLimit]]
        }
      }

    }

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }
  }


  group "gateway-proxy" {
    count = [[.gatewayProxy.replicas]]

    task "gateway-proxy" {
      driver = "docker"
      config {
        image = "[[.gatewayProxy.image.registry]]/[[.gatewayProxy.image.repository]]:[[.gatewayProxy.image.tag]]"
        port_map {
          http = 8080
          https = 8443
          admin = 19000
        }
        entrypoint = ["envoy"]
        args = [
          "-c",
          "${NOMAD_TASK_DIR}/envoy.yaml",
          "--disable-hot-restart",
          "-l debug",
        ]
      }

      template {
        data = <<EOF
node:
  cluster: gateway
  id: gateway~{{ env "NOMAD_ALLOC_ID" }}
  metadata:
    # this line must match !
    role: "gloo-system~gateway-proxy"

static_resources:
  clusters:
  - name: xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: xds_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: {{ env "NOMAD_IP_gloo_xds" }}
                port_value: {{ env "NOMAD_PORT_gloo_xds" }}
    http2_protocol_options: {}
    type: STATIC

  - name: admin_port_cluster
    connect_timeout: 5.000s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: admin_port_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 19000

  listeners:
    - name: prometheus_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8081
      filter_chains:
        - filters:
            - name: envoy.http_connection_manager
              config:
                codec_type: auto
                stat_prefix: prometheus
                route_config:
                  name: prometheus_route
                  virtual_hosts:
                    - name: prometheus_host
                      domains:
                        - "*"
                      routes:
                        - match:
                            path: "/ready"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            path: "/server_info"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            prefix: "/metrics"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            prefix_rewrite: "/stats/prometheus"
                            cluster: admin_port_cluster
                http_filters:
                  - name: envoy.router
                    config: {}

dynamic_resources:
  ads_config:
    api_type: GRPC
    grpc_services:
    - envoy_grpc: {cluster_name: xds_cluster}
  cds_config:
    ads: {}
  lds_config:
    ads: {}

admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 0.0.0.0
      port_value: 19000
EOF

        destination = "${NOMAD_TASK_DIR}/envoy.yaml"
      }

      resources {
        # cpu required in MHz
        cpu = [[.gatewayProxy.cpuLimit]]

        # memory required in MB
        memory = [[.gatewayProxy.memLimit]]

        network {
          # bandwidth required in MBits
          mbits = [[.gatewayProxy.bandwidthLimit]]

          port "http" {}
          port "https" {}
          port "admin" {}
          port "stats" {}

        }
      }

      service {
        name = "gateway-proxy"
        tags = [
          "gloo",
          "http",
        ]
        port = "http"
        check {
          name = "alive"
          type = "tcp"
          interval = "10s"
          timeout = "2s"
        }
      }

      service {
        name = "gateway-proxy"
        tags = [
          "gloo",
          "https",
        ]
        port = "https"
        check {
          name = "alive"
          type = "tcp"
          interval = "10s"
          timeout = "2s"
        }
      }

      service {
        name = "gateway-proxy"
        tags = [
          "gloo",
          "admin",
        ]
        port = "admin"
        check {
          name = "alive"
          type = "tcp"
          interval = "10s"
          timeout = "2s"
        }
      }
    }

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

  }

 }
