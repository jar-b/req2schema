# req2schema

Generate a terraform provider `Resource` schema from an AWS request schema.

This tool is intended to supplement the Terraform AWS provider's
[skaff](https://github.com/hashicorp/terraform-provider-aws/tree/main/skaff) tool,
focusing specifically on a resources schema. `skaff` does not use any external
data to generate its source files, and therefore populates a placeholder `Schema`.
`req2schema` uses AWS HTTP request schemas to provide a stubbed out terraform resource
schema with the appropriate key names and data types (including nested objects). This 
can be helpful for large resources which are time consuming to write by hand.

## Installation

```console
$ go install github.com/jar-b/req2schema@latest
```

## Usage

```
Usage: req2schema [flags] [filename]

Flags:
  -out string
        output file (optional, prints to stdout if omitted)
```

### Examples

The following is an example of the AWS request schema syntax:

```
{
   "thing": "string",
   "second": number,
   "another": boolean,
   "nested": [
      {
         "value": "string"
      }
   ]
}
```

`req2schema` will infer data types from the field values and generate go source code with a 
terraform resource schema definition:

```console
$ req2schema testdata/simple.json
```

```go
package someservice

import (
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
        return &schema.Resource{
                Schema: map[string]*schema.Schema{
                        "another": {
                                Type: schema.TypeBool,
                        },
                        "nested": {
                                Type: schema.TypeList,
                                Elems: &schema.Resource{
                                        Schema: map[string]*schema.Schema{
                                                "value": {
                                                        Type: schema.TypeString,
                                                },
                                        },
                                },
                        },
                        "second": {
                                Type: schema.TypeFloat,
                        },
                        "thing": {
                                Type: schema.TypeString,
                        },
                        "tags":     tftags.TagsSchema(),
                        "tags_all": tftags.TagsSchemaComputed(),
                },
        }
}

```

The [`testdata`](testdata/) subfolder contains additional requests directly from the 
AWS documentation.

## Additional Considerations

- The resulting source code will be framed out, but incomplete.
  - All items will need one of `Required`, `Optional`, or `Computed` set (this cannot be
  inferred from the HTTP request schema alone).
  - `enum` types will need validation added.
  - If unable to determine the type from the incoming request, a `TODO` comment is added
  after the type definition as a reminder to determine the appropriate type.
- Not all AWS services include JSON request schemas in the documentation. Some services
use XML examples, which this tool does not support.
