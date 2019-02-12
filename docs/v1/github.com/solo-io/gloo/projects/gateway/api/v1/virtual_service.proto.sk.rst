
===================================================
Package: `gateway.solo.io`
===================================================

.. _gateway.solo.io.github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto:


**Types:**


- :ref:`gateway.solo.io.VirtualService` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto <https://github.com/solo-io/gloo/blob/master/projects/gateway/api/v1/virtual_service.proto>`_





.. _gateway.solo.io.VirtualService:

VirtualService
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
A virtual service describes the set of routes to match for a set of domains.
Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).


::


   "virtual_host": .gloo.solo.io.VirtualHost
   "ssl_config": .gloo.solo.io.SslConfig
   "display_name": string
   "status": .core.solo.io.Status
   "metadata": .core.solo.io.Metadata

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `virtual_host` | :ref:`gloo.solo.io.VirtualHost` |  | 
   `ssl_config` | :ref:`gloo.solo.io.SslConfig` | If provided, the Gateway will serve TLS/SSL traffic for this set of routes | 
   `display_name` | `string` | Display only, optional descriptive name. Unlike metadata.name, DisplayName can be changed without deleting the resource. | 
   `status` | :ref:`core.solo.io.Status` | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
