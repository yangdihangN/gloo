
===================================================
Package: `static.plugins.gloo.solo.io`
===================================================

.. _static.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/static/static.proto:


**Types:**


- :ref:`static.plugins.gloo.solo.io.UpstreamSpec`
- :ref:`static.plugins.gloo.solo.io.Host`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/static/static.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/static/static.proto>`_





.. _static.plugins.gloo.solo.io.UpstreamSpec:

UpstreamSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Static upstreams are used to route request to services listening at fixed IP/Addresses.
Static upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
for additional service-specific configuration.
Unlike upstreams created by service discovery, Static Upstreams must be created manually by users


::


   "hosts": []static.plugins.gloo.solo.io.Host
   "use_tls": bool
   "service_spec": .plugins.gloo.solo.io.ServiceSpec
   "use_http2": bool

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `hosts` | :ref:`[]static.plugins.gloo.solo.io.Host` | A list of addresses and ports at least one must be specified | 
   `use_tls` | `bool` | Attempt to use outbound TLS Gloo will automatically set this to true for port 443 | 
   `service_spec` | :ref:`plugins.gloo.solo.io.ServiceSpec` | An optional Service Spec describing the service listening at this address | 
   `use_http2` | `bool` | Use http2 when communicating with this upstream | 



.. _static.plugins.gloo.solo.io.Host:

Host
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Represents a single instance of an upstream


::


   "addr": string
   "port": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `addr` | `string` | Address (hostname or IP) | 
   `port` | `int` | Port the instance is listening on | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
