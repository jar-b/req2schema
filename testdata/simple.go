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
