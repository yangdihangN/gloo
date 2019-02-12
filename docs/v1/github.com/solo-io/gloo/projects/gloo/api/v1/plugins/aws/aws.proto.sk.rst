
===================================================
Package: `aws.plugins.gloo.solo.io`
===================================================

.. _aws.plugins.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws.proto:


**Types:**


- :ref:`aws.plugins.gloo.solo.io.UpstreamSpec`
- :ref:`aws.plugins.gloo.solo.io.LambdaFunctionSpec`
- :ref:`aws.plugins.gloo.solo.io.DestinationSpec`
- [InvocationStyle](#InvocationStyle)
  



**Source File:** `github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws.proto <https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/aws/aws.proto>`_





.. _aws.plugins.gloo.solo.io.UpstreamSpec:

UpstreamSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Upstream Spec for AWS Lambda Upstreams
AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
in a particular region


::


   "region": string
   "secret_ref": .core.solo.io.ResourceRef
   "lambda_functions": []aws.plugins.gloo.solo.io.LambdaFunctionSpec

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `region` | `string` | The AWS Region where the desired Lambda Functions exxist | 
   `secret_ref` | :ref:`core.solo.io.ResourceRef` | A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret AWS Secrets can be created with `glooctl secret create aws ...` If the secret is created manually, it must conform to the following structure: ``` access_key: <aws access key> secret_key: <aws secret key> ``` | 
   `lambda_functions` | :ref:`[]aws.plugins.gloo.solo.io.LambdaFunctionSpec` | The list of Lambda Functions contained within this region. This list will be automatically populated by Gloo if discovery is enabled for AWS Lambda Functions | 



.. _aws.plugins.gloo.solo.io.LambdaFunctionSpec:

LambdaFunctionSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions:
- name of the function
- qualifier for the function


::


   "logical_name": string
   "lambda_function_name": string
   "qualifier": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `logical_name` | `string` | the logical name gloo should associate with this function. if left empty, it will default to lambda_function_name+qualifier | 
   `lambda_function_name` | `string` | The Name of the Lambda Function as it appears in the AWS Lambda Portal | 
   `qualifier` | `string` | The Qualifier for the Lambda Function. Qualifiers act as a kind of version for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info. | 



.. _aws.plugins.gloo.solo.io.DestinationSpec:

DestinationSpec
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions


::


   "logical_name": string
   "invocation_style": .aws.plugins.gloo.solo.io.DestinationSpec.InvocationStyle
   "response_transformation": bool

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `logical_name` | `string` | The Logical Name of the LambdaFunctionSpec to be invoked. | 
   `invocation_style` | :ref:`aws.plugins.gloo.solo.io.DestinationSpec.InvocationStyle` | Can be either Sync or Async. | 
   `response_transformation` | `bool` | de-jsonify response bodies returned from aws lambda | 



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
