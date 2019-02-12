
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


- :ref:`message.google.protobuf.DoubleValue`
- :ref:`message.google.protobuf.FloatValue`
- :ref:`message.google.protobuf.Int64Value`
- :ref:`message.google.protobuf.UInt64Value`
- :ref:`message.google.protobuf.Int32Value`
- :ref:`message.google.protobuf.UInt32Value`
- :ref:`message.google.protobuf.BoolValue`
- :ref:`message.google.protobuf.StringValue`
- :ref:`message.google.protobuf.BytesValue`
  



**Source File:** `google/protobuf/wrappers.proto`




.. _message.google.protobuf.DoubleValue:

DoubleValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `double`.

The JSON representation for `DoubleValue` is JSON number.


::


   "value": float



.. _field.google.protobuf.DoubleValue.value:

value
++++++++++++++++++++++++++

Type: `float` 

Description: The double value. 






.. _message.google.protobuf.FloatValue:

FloatValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `float`.

The JSON representation for `FloatValue` is JSON number.


::


   "value": float



.. _field.google.protobuf.FloatValue.value:

value
++++++++++++++++++++++++++

Type: `float` 

Description: The float value. 






.. _message.google.protobuf.Int64Value:

Int64Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `int64`.

The JSON representation for `Int64Value` is JSON string.


::


   "value": int



.. _field.google.protobuf.Int64Value.value:

value
++++++++++++++++++++++++++

Type: `int` 

Description: The int64 value. 






.. _message.google.protobuf.UInt64Value:

UInt64Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `uint64`.

The JSON representation for `UInt64Value` is JSON string.


::


   "value": int



.. _field.google.protobuf.UInt64Value.value:

value
++++++++++++++++++++++++++

Type: `int` 

Description: The uint64 value. 






.. _message.google.protobuf.Int32Value:

Int32Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `int32`.

The JSON representation for `Int32Value` is JSON number.


::


   "value": int



.. _field.google.protobuf.Int32Value.value:

value
++++++++++++++++++++++++++

Type: `int` 

Description: The int32 value. 






.. _message.google.protobuf.UInt32Value:

UInt32Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `uint32`.

The JSON representation for `UInt32Value` is JSON number.


::


   "value": int



.. _field.google.protobuf.UInt32Value.value:

value
++++++++++++++++++++++++++

Type: `int` 

Description: The uint32 value. 






.. _message.google.protobuf.BoolValue:

BoolValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `bool`.

The JSON representation for `BoolValue` is JSON `true` and `false`.


::


   "value": bool



.. _field.google.protobuf.BoolValue.value:

value
++++++++++++++++++++++++++

Type: `bool` 

Description: The bool value. 






.. _message.google.protobuf.StringValue:

StringValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `string`.

The JSON representation for `StringValue` is JSON string.


::


   "value": string



.. _field.google.protobuf.StringValue.value:

value
++++++++++++++++++++++++++

Type: `string` 

Description: The string value. 






.. _message.google.protobuf.BytesValue:

BytesValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
Wrapper message for `bytes`.

The JSON representation for `BytesValue` is JSON string.


::


   "value": bytes



.. _field.google.protobuf.BytesValue.value:

value
++++++++++++++++++++++++++

Type: `bytes` 

Description: The bytes value. 







.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
