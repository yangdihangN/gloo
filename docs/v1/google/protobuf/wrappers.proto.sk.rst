
===================================================
Package: `google.protobuf`
===================================================  
Protocol Buffers - Google's data interchange format
Copyright 2008 Google Inc.  All rights reserved.
https://developers.google.com/protocol-buffers/

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.  
Wrappers for primitive (non-message) types. These types are useful
for embedding primitives in the `google.protobuf.Any` type and for places
where we need to distinguish between the absence of a primitive
typed field and its default value.




.. _google.protobuf.google/protobuf/wrappers.proto:


**Types:**


- :ref:`google.protobuf.DoubleValue`
- :ref:`google.protobuf.FloatValue`
- :ref:`google.protobuf.Int64Value`
- :ref:`google.protobuf.UInt64Value`
- :ref:`google.protobuf.Int32Value`
- :ref:`google.protobuf.UInt32Value`
- :ref:`google.protobuf.BoolValue`
- :ref:`google.protobuf.StringValue`
- :ref:`google.protobuf.BytesValue`
  



**Source File:** `google/protobuf/wrappers.proto`





.. _google.protobuf.DoubleValue:

DoubleValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `double`.

The JSON representation for `DoubleValue` is JSON number.


::


   "value": float

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `float` | The double value. | 



.. _google.protobuf.FloatValue:

FloatValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `float`.

The JSON representation for `FloatValue` is JSON number.


::


   "value": float

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `float` | The float value. | 



.. _google.protobuf.Int64Value:

Int64Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `int64`.

The JSON representation for `Int64Value` is JSON string.


::


   "value": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `int` | The int64 value. | 



.. _google.protobuf.UInt64Value:

UInt64Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `uint64`.

The JSON representation for `UInt64Value` is JSON string.


::


   "value": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `int` | The uint64 value. | 



.. _google.protobuf.Int32Value:

Int32Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `int32`.

The JSON representation for `Int32Value` is JSON number.


::


   "value": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `int` | The int32 value. | 



.. _google.protobuf.UInt32Value:

UInt32Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `uint32`.

The JSON representation for `UInt32Value` is JSON number.


::


   "value": int

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `int` | The uint32 value. | 



.. _google.protobuf.BoolValue:

BoolValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `bool`.

The JSON representation for `BoolValue` is JSON `true` and `false`.


::


   "value": bool

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `bool` | The bool value. | 



.. _google.protobuf.StringValue:

StringValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `string`.

The JSON representation for `StringValue` is JSON string.


::


   "value": string

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `string` | The string value. | 



.. _google.protobuf.BytesValue:

BytesValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `bytes`.

The JSON representation for `BytesValue` is JSON string.


::


   "value": bytes

.. csv-table:: Fields Reference
   :header: "Field" , "Type", "Description", "Default"
   :delim: |


   `value` | `bytes` | The bytes value. | 




.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
