
===================================================
Package: `clusteringress.gloo.solo.io`
===================================================

.. _clusteringress.gloo.solo.io.github.com/solo-io/gloo/projects/clusteringress/api/v1/cluster_ingress.proto:


**Types:**


- :ref:`clusteringress.gloo.solo.io.ClusterIngress` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/clusteringress/api/v1/cluster_ingress.proto <https://github.com/solo-io/gloo/blob/master/projects/clusteringress/api/v1/cluster_ingress.proto>`_





.. _clusteringress.gloo.solo.io.ClusterIngress:

ClusterIngress
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
A simple wrapper for a kNative ClusterIngress Object.


::


   "metadata": .core.solo.io.Metadata
   "status": .core.solo.io.Status
   "cluster_ingress_spec": .google.protobuf.Any
   "cluster_ingress_status": .google.protobuf.Any

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `metadata` | :ref:`core.solo.io.Metadata` |  | 
   `status` | :ref:`core.solo.io.Status` |  | 
   `cluster_ingress_spec` | `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ | a raw byte representation of the cluster ingress this resource wraps | 
   `cluster_ingress_status` | `.google.protobuf.Any<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any>`_ | a raw byte representation of the ingress status of the cluster ingress object | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
