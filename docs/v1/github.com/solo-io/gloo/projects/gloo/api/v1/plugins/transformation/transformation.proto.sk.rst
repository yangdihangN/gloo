
===================================================
Package: `transformation.plugins.gloo.solo.io`
===================================================  
TODO: this was copied form the transformation filter.
TODO: instead of manually copying, we want to do it via script, similar to the java-control-plane
TODO: to solo-kit/api/envoy




.. _transformation.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/transformation.proto:


**Types:**


- :ref:`message.transformation.plugins.gloo.solo.io.RouteTransformations`
- :ref:`message.transformation.plugins.gloo.solo.io.Transformation`
- :ref:`message.transformation.plugins.gloo.solo.io.Extraction`
- :ref:`message.transformation.plugins.gloo.solo.io.TransformationTemplate`
- :ref:`message.transformation.plugins.gloo.solo.io.InjaTemplate`
- :ref:`message.transformation.plugins.gloo.solo.io.Passthrough`
- :ref:`message.transformation.plugins.gloo.solo.io.MergeExtractorsToBody`
- :ref:`message.transformation.plugins.gloo.solo.io.HeaderBodyTransform`
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/transformation.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/transformation/transformation.proto>`_




.. _message.transformation.plugins.gloo.solo.io.RouteTransformations:

RouteTransformations
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "request_transformation": .transformation.plugins.gloo.solo.io.Transformation
   "response_transformation": .transformation.plugins.gloo.solo.io.Transformation



.. _field.transformation.plugins.gloo.solo.io.RouteTransformations.request_transformation:

request_transformation
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.Transformation` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.RouteTransformations.response_transformation:

response_transformation
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.Transformation` 

Description:  






.. _message.transformation.plugins.gloo.solo.io.Transformation:

Transformation
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
[#proto-status: experimental]


::


   "transformation_template": .transformation.plugins.gloo.solo.io.TransformationTemplate
   "header_body_transform": .transformation.plugins.gloo.solo.io.HeaderBodyTransform



.. _field.transformation.plugins.gloo.solo.io.Transformation.transformation_template:

transformation_template
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.TransformationTemplate` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.Transformation.header_body_transform:

header_body_transform
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.HeaderBodyTransform` 

Description:  






.. _message.transformation.plugins.gloo.solo.io.Extraction:

Extraction
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "header": string
   "regex": string
   "subgroup": int



.. _field.transformation.plugins.gloo.solo.io.Extraction.header:

header
++++++++++++++++++++++++++

Type: `string` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.Extraction.regex:

regex
++++++++++++++++++++++++++

Type: `string` 

Description: what information to extract. if extraction fails the result is an empty value. 



.. _field.transformation.plugins.gloo.solo.io.Extraction.subgroup:

subgroup
++++++++++++++++++++++++++

Type: `int` 

Description:  






.. _message.transformation.plugins.gloo.solo.io.TransformationTemplate:

TransformationTemplate
~~~~~~~~~~~~~~~~~~~~~~~~~~



::


   "advanced_templates": bool
   "extractors": map<string, .transformation.plugins.gloo.solo.io.Extraction>
   "headers": map<string, string>
   "body": .transformation.plugins.gloo.solo.io.InjaTemplate
   "passthrough": .transformation.plugins.gloo.solo.io.Passthrough
   "merge_extractors_to_body": .transformation.plugins.gloo.solo.io.MergeExtractorsToBody



.. _field.transformation.plugins.gloo.solo.io.TransformationTemplate.advanced_templates:

advanced_templates
++++++++++++++++++++++++++

Type: `bool` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.TransformationTemplate.extractors:

extractors
++++++++++++++++++++++++++

Type: `map<string, .transformation.plugins.gloo.solo.io.Extraction>` 

Description: Extractors are in the origin request language domain 



.. _field.transformation.plugins.gloo.solo.io.TransformationTemplate.headers:

headers
++++++++++++++++++++++++++

Type: `map<string, string>` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.TransformationTemplate.body:

body
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.InjaTemplate` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.TransformationTemplate.passthrough:

passthrough
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.Passthrough` 

Description:  



.. _field.transformation.plugins.gloo.solo.io.TransformationTemplate.merge_extractors_to_body:

merge_extractors_to_body
++++++++++++++++++++++++++

Type: :ref:`message.transformation.plugins.gloo.solo.io.MergeExtractorsToBody` 

Description:  






.. _message.transformation.plugins.gloo.solo.io.InjaTemplate:

InjaTemplate
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
custom functions:
header_value(name) -> from the original headers
extracted_value(name, index) -> from the extracted values


::


   "text": string



.. _field.transformation.plugins.gloo.solo.io.InjaTemplate.text:

text
++++++++++++++++++++++++++

Type: `string` 

Description:  






.. _message.transformation.plugins.gloo.solo.io.Passthrough:

Passthrough
~~~~~~~~~~~~~~~~~~~~~~~~~~



::








.. _message.transformation.plugins.gloo.solo.io.MergeExtractorsToBody:

MergeExtractorsToBody
~~~~~~~~~~~~~~~~~~~~~~~~~~



::








.. _message.transformation.plugins.gloo.solo.io.HeaderBodyTransform:

HeaderBodyTransform
~~~~~~~~~~~~~~~~~~~~~~~~~~



::









.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
