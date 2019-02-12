
===================================================
Package: `consul.plugins.gloo.solo.io`
===================================================

.. _consul.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/consul/consul.proto:


**Types:**


- :ref:`consul.plugins.gloo.solo.io.UpstreamSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/consul/consul.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/consul/consul.proto>`_





.. _consul.plugins.gloo.solo.io.UpstreamSpec:

UpstreamSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Upstream Spec for Consul Upstreams
consul Upstreams represent a set of one or more addressable pods for a consul Service
the Gloo consul Upstream maps to a single service port. Because consul Services support multiple ports,
Gloo requires that a different upstream be created for each port
consul Upstreams are typically generated automatically by Gloo from the consul API


::


   "service_name": string
   "service_tags": []string
   "service_spec": .plugins.gloo.solo.io.ServiceSpec
   "connect_enabled": bool

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `service_name` | `string` | The name of the Consul Service | 
   `service_tags` | `[]string` | The list of service tags Gloo should search for on a service instance before deciding whether or not to include the instance as part of this upstream | 
   `service_spec` | :ref:`plugins.gloo.solo.io.ServiceSpec` | An optional Service Spec describing the service listening at this address | 
   `connect_enabled` | `bool` | is this consul service connect enabled. | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
