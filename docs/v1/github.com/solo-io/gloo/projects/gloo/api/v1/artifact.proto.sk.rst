
===================================================
Package: `gloo.solo.io`
===================================================

.. _gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/artifact.proto:


**Types:**


- :ref:`gloo.solo.io.Artifact` **Top-Level Resource**
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/artifact.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/artifact.proto>`_





.. _gloo.solo.io.Artifact:

Artifact
~~~~~~~~~~~~~~~~~~~~~~~~~~

 

Gloo Artifacts are used by Gloo to store small bits of binary or file data.

Certain plugins such as the gRPC plugin read and write artifacts to one of Gloo's configured
storage layer.

Artifacts can be backed by files on disk, Kubernetes ConfigMaps, and Consul Key/Value pairs.

Supported artifact backends can be selected in Gloo's boostrap options.


::


   "data": string
   "metadata": .core.solo.io.Metadata

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `data` | `string` | Raw data data being stored | 
   `metadata` | :ref:`core.solo.io.Metadata` | Metadata contains the object metadata for this resource | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
