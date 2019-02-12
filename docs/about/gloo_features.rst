Gloo Features
*************

**Supported Platforms**:

- Kubernetes

- HashiCorp Stack (Vault, Consul, Nomad)

- AWS Lambda

- Knative

- Microsoft Azure Functions

- Google Cloud Platform Functions

**Routing Features**:

- **Dynamic Load Balancing**: Load balance traffic across multiple upstream services.

- **Health Checks**: Active and passive monitoring of your upstream services.

- **OpenTracing**: Monitor requests using the well-supported OpenTracing standard

- **Monitoring**: Export HTTP metrics to Prometheus or Statsd

- **SSL**: Highly customizable options for adding SSL encryption to upstream services with full support for SNI.

- **Transformations**: Add, remove, or manipulate HTTP requests and responses.

- **Automated API Translation**: Automatically transform client requests to upstream API calls using Gloo’s Function Discovery

- **CLI**: Control your Gloo cluster from the command line.

- **Declarative API**: Gloo features a declarative YAML-based API; store your configuration as code and commit it with your projects.

- **Failure Recovery**: Gloo is completely stateless and will immediately return to the desired configuration at boot time.

- **Scalability**: Gloo acts as a control plane for Envoy, allowing Envoy instances and Gloo instances to be scaled independently. Both Gloo and Envoy are stateless.

- **Performance**: Gloo leverages Envoy for its high performance and low footprint.

- **Plugins**: Extendable architecture for adding functionality and integrations to Gloo.

- **Tooling**: Build and Deployment tool for customized builds and deployment options

- **Events**: Invoke APIs using CloudEvents.

- **Pub/Sub**: Publish HTTP requests to NATS

- **JSON-to-gRPC transcoding**: Connect JSON clients to gRPC services