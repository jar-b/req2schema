package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availabilityZones": {
				Type:  schema.TypeList,
				Elems: &schema.Schema{Type: schema.TypeString},
			},
			"multiplexSettings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maximumVideoBufferDelayMilliseconds": {
							Type: schema.TypeFloat,
						},
						"transportStreamBitrate": {
							Type: schema.TypeFloat,
						},
						"transportStreamId": {
							Type: schema.TypeFloat,
						},
						"transportStreamReservedBitrate": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString,
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}
