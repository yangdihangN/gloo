
===================================================
Package: `rest.plugins.gloo.solo.io`
===================================================

.. _rest.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest.proto:


**Types:**


- :ref:`rest.plugins.gloo.solo.io.ServiceSpec`
- :ref:`rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo`
- :ref:`rest.plugins.gloo.solo.io.DestinationSpec`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/rest/rest.proto>`_





.. _rest.plugins.gloo.solo.io.ServiceSpec:

ServiceSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "transformations": map<string, .transformation.plugins.gloo.solo.io.TransformationTemplate>
   "swagger_info": .rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `transformations` | `map<string, .transformation.plugins.gloo.solo.io.TransformationTemplate>` |  | 
   `swagger_info` | :ref:`rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo` |  | 



.. _rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo:

SwaggerInfo
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "url": string
   "inline": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `url` | `string` |  | 
   `inline` | `string` |  | 



.. _rest.plugins.gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
This is only for upstream with REST service spec


::


   "function_name": string
   "parameters": .transformation.plugins.gloo.solo.io.Parameters
   "response_transformation": .transformation.plugins.gloo.solo.io.TransformationTemplate

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `function_name` | `string` |  | 
   `parameters` | :ref:`transformation.plugins.gloo.solo.io.Parameters` |  | 
   `response_transformation` | :ref:`transformation.plugins.gloo.solo.io.TransformationTemplate` |  | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
