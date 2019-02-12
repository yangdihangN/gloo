
===================================================
Package: `core.solo.io`
===================================================

.. _core.solo.io.github.com/solo-io/solo-kit/api/v1/metadata.proto:


**Types:**


- :ref:`message.core.solo.io.Metadata`
  



**Source File:** `github.com/solo-io/solo-kit/api/v1/metadata.proto <https://github.com/solo-io/solo-kit/blob/master/api/v1/metadata.proto>`_




.. _message.core.solo.io.Metadata:

Metadata
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
*
Metadata contains general properties of resources for purposes of versioning, annotating, and namespacing.


::


   "name": string
   "namespace": string
   "resource_version": string
   "labels": map<string, string>
   "annotations": map<string, string>



.. _field.core.solo.io.Metadata.name:

name
++++++++++++++++++++++++++

Type: `string` 

Description: Name of the resource. Names must be unique and follow the following syntax rules: One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters. 



.. _field.core.solo.io.Metadata.namespace:

namespace
++++++++++++++++++++++++++

Type: `string` 

Description: Namespace is used for the namespacing of resources. 



.. _field.core.solo.io.Metadata.resource_version:

resource_version
++++++++++++++++++++++++++

Type: `string` 

Description: An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. 



.. _field.core.solo.io.Metadata.labels:

labels
++++++++++++++++++++++++++

Type: `map<string, string>` 

Description: Map of string keys and values that can be used to organize and categorize (scope and select) objects. Some resources contain `selectors` which can be linked with other resources by their labels 



.. _field.core.solo.io.Metadata.annotations:

annotations
++++++++++++++++++++++++++

Type: `map<string, string>` 

Description: Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
