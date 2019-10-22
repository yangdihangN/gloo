---
menuTitle: Config Validation
title: Config Reporting & Validation in Gloo
weight: 45
description: (Kubernetes Only) Gloo can be configured to validate configuration before it is applied to the cluster. With validation enabled, any attempt to apply invalid configuration to the cluster will be rejected.
---

# Motivation

When configuring an API gateway or edge proxy, invalid configuration can quickly lead to bugs, service outages, and 
security vulnerabilities. 

This document explains features in Gloo designed to prevent invalid configuration from propagating to the 
data plane (the Gateway Proxies).

# How Gloo Validates Configuration

Gloo's configuration takes the form of **Virtual Services** written by users.
Users may also  write **Gateways** (to configure [listeners](https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/listeners) 
and **Route Tables** (to decentralize routing configurations from Virtual Services).

Whenever Gloo configuration objects are updated, Gloo validates and processes the new configuration.

Validation in Gloo is comprised of a four step process:

1. First, resources are admitted (or rejected) via a [Kubernetes Validating Admission Webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/). Configuration options for the webhook live
in the `settings.gloo.solo.io` custom resource.

2. Once the resource is admitted, it is joined together with the other configuration objects to produce a finalized configuration snapshot, in the form of the internally-facing [**Proxy** config object]({{< protobuf name="gloo.solo.io.Proxy">}}). 

    * If any admitted virtual service has invalid configuration, it will be omitted from the **Proxy**.
    
    * If an admitted virtual service becomes invalid due to an update, the last valid configuration of that virtual service will be persisted.

3. Gloo will report a `status` on the admitted resources indicating whether they were accepted as part of the snapshot and stored on the **Proxy**.

4. Gloo processes the **Proxy** along with service discovery data to produce the final 
[Envoy xDS Snapshot](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol). 
If configuration errors are encountered at this point, Gloo will report them to the Proxy as
 well as the user-facing config objects which produced the Proxy. At this point, Envoy
  configuration will not be updated until the errors are resolved.

Each *Proxy* gets its own configuration; if config for an individual proxy is invalid, it does not affect the other proxies.
The proxy that *Gateways* and their *Virtual Services* will be applied to can be configured via the `proxyNames` option on 
  the [`Gateway` resource]({{< protobuf name="gateway.solo.io.v2.Gateway">}}).

{{% notice note %}}

- You can run `glooctl check` locally to easily spot any configuration errors on resources that have been admitted to your cluster.

{{% /notice %}}

# Warnings and Errors

Gloo processes an admitted config resource, it can report one of three status types on the resource:

- *Accepted*: The resource has been accepted and applied to the system.
- *Rejected*: The resource has invalid configuration and has not been applied to the system.
- *Warning*: The resource has valid config but points to a missing/misconfigured resource.

When a resource is in *Rejected* or *Warning* state, its configuration is not propagated to the proxy.

# Using the Validating Webhook

Admission Validation provides a safeguard to ensure invalid config does not make its way 
into the Gloo snapshot, which would otherwise block updates from reaching the proxy. If a 
resource would be written or modified in such a way to cause Gloo to report an error, it is 
instead rejected by the Kubernetes API Server before it is written to persistent storage.

See [the Admission Controller Guide]({{< ref "/gloo_routing/validation/admission_control">}}) 
to learn how to configure and use Gloo's admission control feature.

# Sanitizing Config

Gloo can be configured to pass partially config to Envoy by admitting it through an internal process referred to as *sanitizing*.

Rather than refuse to update Envoy with invalid config, Gloo can replace the invalid pieces of configuration with preconfigured 
defaults.

See [the Route Replacement Guide]({{< ref "/gloo_routing/validation/invalid_route_replacement">}}) 
to learn how to configure and use Gloo's sanitization feature.

# Further Reading

{{% children description="true" %}}

We appreciate questions and feedback on Gloo validation or any other feature on [the solo.io slack channel](https://slack.solo.io/) as well as our [GitHub issues page](https://github.com/solo-io/gloo).


