
===================================================
Package: `plugins.gloo.solo.io`
===================================================

.. _plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/service_spec.proto:


**Types:**


- :ref:`message.plugins.gloo.solo.io.ServiceSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/service_spec.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/service_spec.proto>`_




.. _message.plugins.gloo.solo.io.ServiceSpec:

ServiceSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Describes APIs and application-level information for services
Gloo routes to. ServiceSpec is contained within the UpstreamSpec for certain types
of upstreams, including Kubernetes, Consul, and Static.
ServiceSpec configuration is opaque to Gloo and handled by Service Plugins.


::


   "rest": .rest.plugins.gloo.solo.io.ServiceSpec
   "grpc": .grpc.plugins.gloo.solo.io.ServiceSpec



.. _field.plugins.gloo.solo.io.ServiceSpec.rest:

rest
++++++++++++++++++++++++++

Type: :ref:`message.rest.plugins.gloo.solo.io.ServiceSpec` 

Description:  



.. _field.plugins.gloo.solo.io.ServiceSpec.grpc:

grpc
++++++++++++++++++++++++++

Type: :ref:`message.grpc.plugins.gloo.solo.io.ServiceSpec` 

Description:  







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
