package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zones": {
				Type:  schema.TypeList,
				Elems: &schema.Schema{Type: schema.TypeString},
			},
			"multiplex_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maximum_video_buffer_delay_milliseconds": {
							Type: schema.TypeFloat,
						},
						"transport_stream_bitrate": {
							Type: schema.TypeFloat,
						},
						"transport_stream_id": {
							Type: schema.TypeFloat,
						},
						"transport_stream_reserved_bitrate": {
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
