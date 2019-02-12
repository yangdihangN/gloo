
===================================================
Package: `transformation.plugins.gloo.solo.io`
===================================================  
TODO: this was copied form the transformation filter.
TODO: instead of manually copying, we want to do it via script, similar to the java-control-plane
TODO: to solo-kit/api/envoy




.. _transformation.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/transformation.proto:


**Types:**


- :ref:`transformation.plugins.gloo.solo.io.RouteTransformations`
- :ref:`transformation.plugins.gloo.solo.io.Transformation`
- :ref:`transformation.plugins.gloo.solo.io.Extraction`
- :ref:`transformation.plugins.gloo.solo.io.TransformationTemplate`
- :ref:`transformation.plugins.gloo.solo.io.InjaTemplate`
- :ref:`transformation.plugins.gloo.solo.io.Passthrough`
- :ref:`transformation.plugins.gloo.solo.io.MergeExtractorsToBody`
- :ref:`transformation.plugins.gloo.solo.io.HeaderBodyTransform`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/transformation.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/transformation/transformation.proto>`_





.. _transformation.plugins.gloo.solo.io.RouteTransformations:

RouteTransformations
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "request_transformation": .transformation.plugins.gloo.solo.io.Transformation
   "response_transformation": .transformation.plugins.gloo.solo.io.Transformation

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `request_transformation` | :ref:`transformation.plugins.gloo.solo.io.Transformation` |  | 
   `response_transformation` | :ref:`transformation.plugins.gloo.solo.io.Transformation` |  | 



.. _transformation.plugins.gloo.solo.io.Transformation:

Transformation
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
[#proto-status: experimental]


::


   "transformation_template": .transformation.plugins.gloo.solo.io.TransformationTemplate
   "header_body_transform": .transformation.plugins.gloo.solo.io.HeaderBodyTransform

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `transformation_template` | :ref:`transformation.plugins.gloo.solo.io.TransformationTemplate` |  | 
   `header_body_transform` | :ref:`transformation.plugins.gloo.solo.io.HeaderBodyTransform` |  | 



.. _transformation.plugins.gloo.solo.io.Extraction:

Extraction
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "header": string
   "regex": string
   "subgroup": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `header` | `string` |  | 
   `regex` | `string` | what information to extract. if extraction fails the result is an empty value. | 
   `subgroup` | `int` |  | 



.. _transformation.plugins.gloo.solo.io.TransformationTemplate:

TransformationTemplate
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "advanced_templates": bool
   "extractors": map<string, .transformation.plugins.gloo.solo.io.Extraction>
   "headers": map<string, string>
   "body": .transformation.plugins.gloo.solo.io.InjaTemplate
   "passthrough": .transformation.plugins.gloo.solo.io.Passthrough
   "merge_extractors_to_body": .transformation.plugins.gloo.solo.io.MergeExtractorsToBody

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `advanced_templates` | `bool` |  | 
   `extractors` | `map<string, .transformation.plugins.gloo.solo.io.Extraction>` | Extractors are in the origin request language domain | 
   `headers` | `map<string, string>` |  | 
   `body` | :ref:`transformation.plugins.gloo.solo.io.InjaTemplate` |  | 
   `passthrough` | :ref:`transformation.plugins.gloo.solo.io.Passthrough` |  | 
   `merge_extractors_to_body` | :ref:`transformation.plugins.gloo.solo.io.MergeExtractorsToBody` |  | 



.. _transformation.plugins.gloo.solo.io.InjaTemplate:

InjaTemplate
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
custom functions:
header_value(name) -> from the original headers
extracted_value(name, index) -> from the extracted values


::


   "text": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `text` | `string` |  | 



.. _transformation.plugins.gloo.solo.io.Passthrough:

Passthrough
~~~~~~~~~~~~~~~~~~~~~~~~~~



::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _transformation.plugins.gloo.solo.io.MergeExtractorsToBody:

MergeExtractorsToBody
~~~~~~~~~~~~~~~~~~~~~~~~~~



::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |





.. _transformation.plugins.gloo.solo.io.HeaderBodyTransform:

HeaderBodyTransform
~~~~~~~~~~~~~~~~~~~~~~~~~~



::



.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |






.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
