
===================================================
Package: `azure.plugins.gloo.solo.io`
===================================================

.. _azure.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/azure/azure.proto:


**Types:**


- :ref:`azure.plugins.gloo.solo.io.UpstreamSpec`
- :ref:`azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec`
- [AuthLevel](#AuthLevel)
- :ref:`azure.plugins.gloo.solo.io.DestinationSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/azure/azure.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/azure/azure.proto>`_





.. _azure.plugins.gloo.solo.io.UpstreamSpec:

UpstreamSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Upstream Spec for Azure Functions Upstreams
Azure Upstreams represent a collection of Azure Functions for a particular Azure Account
within a particular Function App


::


   "function_app_name": string
   "secret_ref": .core.solo.io.ResourceRef
   "functions": []azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `function_app_name` | `string` | The Name of the Azure Function App where the functions are grouped | 
   `secret_ref` | :ref:`core.solo.io.ResourceRef` | A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an [Azure Publish Profile JSON file](https://azure.microsoft.com/en-us/downloads/publishing-profile-overview/). {{ hide_not_implemented "Azure Secrets can be created with `glooctl secret create azure ...`" }} Note that this secret is not required unless Function Discovery is enabled | 
   `functions` | :ref:`[]azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec` |  | 



.. _azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec:

FunctionSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Function Spec for Functions on Azure Functions Upstreams
The Function Spec contains data necessary for Gloo to invoke Azure functions


::


   "function_name": string
   "auth_level": .azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `function_name` | `string` | The Name of the Azure Function as it appears in the Azure Functions Portal | 
   `auth_level` | :ref:`azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel` | Auth Level can bve either "anonymous" "function" or "admin" See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details | 



---
### <a name="AuthLevel">AuthLevel</a>



.. csv-table:: Enum Reference
   :header: "Name", "Description"
   :delim: |


   `Anonymous` | 

   `Function` | 

   `Admin` | 




.. _azure.plugins.gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "function_name": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `function_name` | `string` | The Function Name of the FunctionSpec to be invoked. | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
