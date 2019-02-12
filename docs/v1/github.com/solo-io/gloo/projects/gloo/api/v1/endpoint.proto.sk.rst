
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/endpoint.proto:


**Types:**


- :ref:`gloo.solo.io.Endpoint` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/endpoint.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/endpoint.proto>`_





.. _gloo.solo.io.Endpoint:

Endpoint
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

Endpoints represent dynamically discovered address/ports where an upstream service is listening


::


   "upstreams": []core.solo.io.ResourceRef
   "address": string
   "port": int
   "metadata": .core.solo.io.Metadata

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `upstreams` | :ref:`[]core.solo.io.ResourceRef` | List of the upstreams the endpoint belongs to | 
   `address` | `string` | Address of the endpoint (ip or hostname) | 
   `port` | `int` | listening port for the endpoint | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
