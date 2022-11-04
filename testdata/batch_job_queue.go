package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"computeEnvironmentOrder": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"computeEnvironment": {
							Type: schema.TypeString,
						},
						"order": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"jobQueueName": {
				Type: schema.TypeString,
			},
			"priority": {
				Type: schema.TypeFloat,
			},
			"schedulingPolicyArn": {
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
