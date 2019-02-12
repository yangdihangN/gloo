
===================================================
Package: `ingress.solo.io`
===================================================

.. _ingress.solo.io.github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto:


**Types:**


- :ref:`ingress.solo.io.Ingress` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto <https://github.com/solo-io/gloo/blob/master/projects/ingress/api/v1/ingress.proto>`_





.. _ingress.solo.io.Ingress:

Ingress
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

A simple wrapper for a Kubernetes Ingress Object.


::


   "kube_ingress_spec": .google.protobuf.Any
   "kube_ingress_status": .google.protobuf.Any
   "metadata": .core.solo.io.Metadata

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `kube_ingress_spec` | `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ | a raw byte representation of the kubernetes ingress this resource wraps | 
   `kube_ingress_status` | `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ | a raw byte representation of the ingress status of the kubernetes ingress object | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
