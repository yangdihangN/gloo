
===================================================
Package: `gateway.solo.io`
===================================================

.. _gateway.solo.io.github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto:


**Types:**


- :ref:`gateway.solo.io.Gateway` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto <https://github.com/solo-io/gloo/blob/master/projects/gateway/api/v1/gateway.proto>`_





.. _gateway.solo.io.Gateway:

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

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `virtual_services` | :ref:`[]core.solo.io.ResourceRef` | names of the the virtual services, which contain the actual routes for the gateway if the list is empty, the gateway will apply all virtual services to this gateway | 
   `bind_address` | `string` | the bind address the gateway should serve traffic on | 
   `bind_port` | `int` | bind ports must not conflict across gateways in a namespace | 
   `plugins` | :ref:`gloo.solo.io.ListenerPlugins` | top level plugin configuration for all routes on the gateway | 
   `status` | :ref:`core.solo.io.Status` | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
