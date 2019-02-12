
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




.. _google.protobuf.google/protobuf/struct.proto:


**Types:**


- :ref:`message.google.protobuf.Struct`
- :ref:`message.google.protobuf.Value`
- :ref:`message.google.protobuf.ListValue`
  

 

**Enums:**


	- [NullValue](#NullValue)



**Source File:** `google/protobuf/struct.proto`




.. _message.google.protobuf.Struct:

Struct
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
`Struct` represents a structured data value, consisting of fields
which map to dynamically typed values. In some languages, `Struct`
might be supported by a native representation. For example, in
scripting languages like JS a struct is represented as an
object. The details of that representation are described together
with the proto support for the language.

The JSON representation for `Struct` is JSON object.


::


   "fields": map<string, .google.protobuf.Value>



.. _field.google.protobuf.Struct.fields:

fields
++++++++++++++++++++++++++

Type: `map<string, .google.protobuf.Value>` 

Description: Unordered map of dynamically typed values. 






.. _message.google.protobuf.Value:

Value
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
`Value` represents a dynamically typed value which can be either
null, a number, a string, a boolean, a recursive struct value, or a
list of values. A producer of value is expected to set one of that
variants, absence of any variant indicates an error.

The JSON representation for `Value` is JSON value.


::


   "null_value": .google.protobuf.NullValue
   "number_value": float
   "string_value": string
   "bool_value": bool
   "struct_value": .google.protobuf.Struct
   "list_value": .google.protobuf.ListValue



.. _field.google.protobuf.Value.null_value:

null_value
++++++++++++++++++++++++++

Type: `.google.protobuf.NullValue<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/null-value>`_ 

Description: Represents a null value. 



.. _field.google.protobuf.Value.number_value:

number_value
++++++++++++++++++++++++++

Type: `float` 

Description: Represents a double value. 



.. _field.google.protobuf.Value.string_value:

string_value
++++++++++++++++++++++++++

Type: `string` 

Description: Represents a string value. 



.. _field.google.protobuf.Value.bool_value:

bool_value
++++++++++++++++++++++++++

Type: `bool` 

Description: Represents a boolean value. 



.. _field.google.protobuf.Value.struct_value:

struct_value
++++++++++++++++++++++++++

Type: `.google.protobuf.Struct<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/struct>`_ 

Description: Represents a structured value. 



.. _field.google.protobuf.Value.list_value:

list_value
++++++++++++++++++++++++++

Type: `.google.protobuf.ListValue<https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/list-value>`_ 

Description: Represents a repeated `Value`. 






.. _message.google.protobuf.ListValue:

ListValue
~~~~~~~~~~~~~~~~~~~~~~~~~~

 
`ListValue` is a wrapper around a repeated field of values.

The JSON representation for `ListValue` is JSON array.


::


   "values": []google.protobuf.Value



.. _field.google.protobuf.ListValue.values:

values
++++++++++++++++++++++++++

Type: :ref:`message.google.protobuf.Value` 

Description: Repeated field of dynamically typed values. 






### <a name="NullValue">NullValue</a>

Description: `NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

 The JSON representation for `NullValue` is JSON `null`.

.. csv-table:: Fields Reference
   :header: "Name", "Description"
   :delim: |


   NULL_VALUE | Null value.


.. raw:: html
   <!-- Start of HubSpot Embed Code -->
   <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
   <!-- End of HubSpot Embed Code -->
