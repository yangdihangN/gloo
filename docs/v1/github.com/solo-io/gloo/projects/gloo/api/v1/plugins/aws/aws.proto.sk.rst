
===================================================
Package: `aws.plugins.gloo.solo.io`
===================================================

.. _aws.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws.proto:


**Types:**


- :ref:`message.aws.plugins.gloo.solo.io.UpstreamSpec`
- :ref:`message.aws.plugins.gloo.solo.io.LambdaFunctionSpec`
- :ref:`message.aws.plugins.gloo.solo.io.DestinationSpec`
- [InvocationStyle](#InvocationStyle)
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/aws/aws.proto>`_




.. _message.aws.plugins.gloo.solo.io.UpstreamSpec:

UpstreamSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Upstream Spec for AWS Lambda Upstreams
AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
in a particular region


::


   "region": string
   "secret_ref": .core.solo.io.ResourceRef
   "lambda_functions": []aws.plugins.gloo.solo.io.LambdaFunctionSpec



.. _field.aws.plugins.gloo.solo.io.UpstreamSpec.region:

region
++++++++++++++++++++++++++

Type: `string` 

Description: The AWS Region where the desired Lambda Functions exxist 



.. _field.aws.plugins.gloo.solo.io.UpstreamSpec.secret_ref:

secret_ref
++++++++++++++++++++++++++

Type: :ref:`message.core.solo.io.ResourceRef` 

Description: A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret AWS Secrets can be created with `glooctl secret create aws ...` If the secret is created manually, it must conform to the following structure: ``` access_key: <aws access key> secret_key: <aws secret key> ``` 



.. _field.aws.plugins.gloo.solo.io.UpstreamSpec.lambda_functions:

lambda_functions
++++++++++++++++++++++++++

Type: :ref:`message.aws.plugins.gloo.solo.io.LambdaFunctionSpec` 

Description: The list of Lambda Functions contained within this region. This list will be automatically populated by Gloo if discovery is enabled for AWS Lambda Functions 






.. _message.aws.plugins.gloo.solo.io.LambdaFunctionSpec:

LambdaFunctionSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions:
- name of the function
- qualifier for the function


::


   "logical_name": string
   "lambda_function_name": string
   "qualifier": string



.. _field.aws.plugins.gloo.solo.io.LambdaFunctionSpec.logical_name:

logical_name
++++++++++++++++++++++++++

Type: `string` 

Description: the logical name gloo should associate with this function. if left empty, it will default to lambda_function_name+qualifier 



.. _field.aws.plugins.gloo.solo.io.LambdaFunctionSpec.lambda_function_name:

lambda_function_name
++++++++++++++++++++++++++

Type: `string` 

Description: The Name of the Lambda Function as it appears in the AWS Lambda Portal 



.. _field.aws.plugins.gloo.solo.io.LambdaFunctionSpec.qualifier:

qualifier
++++++++++++++++++++++++++

Type: `string` 

Description: The Qualifier for the Lambda Function. Qualifiers act as a kind of version for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info. 






.. _message.aws.plugins.gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions


::


   "logical_name": string
   "invocation_style": .aws.plugins.gloo.solo.io.DestinationSpec.InvocationStyle
   "response_transformation": bool



.. _field.aws.plugins.gloo.solo.io.DestinationSpec.logical_name:

logical_name
++++++++++++++++++++++++++

Type: `string` 

Description: The Logical Name of the LambdaFunctionSpec to be invoked. 



.. _field.aws.plugins.gloo.solo.io.DestinationSpec.invocation_style:

invocation_style
++++++++++++++++++++++++++

Type: :ref:`message.aws.plugins.gloo.solo.io.DestinationSpec.InvocationStyle` 

Description: Can be either Sync or Async. 



.. _field.aws.plugins.gloo.solo.io.DestinationSpec.response_transformation:

response_transformation
++++++++++++++++++++++++++

Type: `bool` 

Description: de-jsonify response bodies returned from aws lambda 






---
### <a name="InvocationStyle">InvocationStyle</a>



.. csv-table:: Enum Reference
   :header: "Name", "Description"
   :delim: |


   `SYNC` | 

   `ASYNC` | 





.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
