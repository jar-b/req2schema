package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compute_environment_order": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compute_environment": {
							Type: schema.TypeString,
						},
						"order": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"job_queue_name": {
				Type: schema.TypeString,
			},
			"priority": {
				Type: schema.TypeFloat,
			},
			"scheduling_policy_arn": {
				Type: schema.TypeString,
			},
			"state": {
				Type: schema.TypeString,
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}
