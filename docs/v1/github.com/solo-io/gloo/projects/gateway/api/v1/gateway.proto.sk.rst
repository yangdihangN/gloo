
===================================================
Package: `gateway.solo.io`
===================================================

.. _gateway.solo.io.github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto:


**Types:**


- :ref:`message.gateway.solo.io.Gateway` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto <https://github.com/solo-io/gloo/blob/master/projects/gateway/api/v1/gateway.proto>`_




.. _message.gateway.solo.io.Gateway:

Gateway
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.


::


   "virtual_services": []core.solo.io.ResourceRef
   "bind_address": string
   "bind_port": int
   "plugins": .gloo.solo.io.ListenerPlugins
   "status": .core.solo.io.Status
   "metadata": .core.solo.io.Metadata



.. _field.gateway.solo.io.Gateway.virtual_services:

virtual_services
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.ResourceRef` 

Description: names of the the virtual services, which contain the actual routes for the gateway if the list is empty, the gateway will apply all virtual services to this gateway 



.. _field.gateway.solo.io.Gateway.bind_address:

bind_address
++++++++++++++++++++++++++

Type: `string` 

Description: the bind address the gateway should serve traffic on 



.. _field.gateway.solo.io.Gateway.bind_port:

bind_port
++++++++++++++++++++++++++

Type: `int` 

Description: bind ports must not conflict across gateways in a namespace 



.. _field.gateway.solo.io.Gateway.plugins:

plugins
++++++++++++++++++++++++++

Type: :ref:`message.gloo.solo.io.ListenerPlugins` 

Description: top level plugin configuration for all routes on the gateway 



.. _field.gateway.solo.io.Gateway.status:

status
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Status` 

Description: Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation 



.. _field.gateway.solo.io.Gateway.metadata:

metadata
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.Metadata` 

Description: Metadata contains the object metadata for this resource 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
