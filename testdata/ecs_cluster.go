package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capacity_providers": {
				Type:  schema.TypeList,
				Elems: &schema.Schema{Type: schema.TypeString},
			},
			"cluster_name": {
				Type: schema.TypeString,
			},
			"configuration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"execute_command_configuration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kms_key_id": {
										Type: schema.TypeString,
									},
									"log_configuration": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloud_watch_encryption_enabled": {
													Type: schema.TypeBool,
												},
												"cloud_watch_log_group_name": {
													Type: schema.TypeString,
												},
												"s3_bucket_name": {
													Type: schema.TypeString,
												},
												"s3_encryption_enabled": {
													Type: schema.TypeBool,
												},
												"s3_key_prefix": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"logging": {
										Type: schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			"default_capacity_provider_strategy": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base": {
							Type: schema.TypeFloat,
						},
						"capacity_provider": {
							Type: schema.TypeString,
						},
						"weight": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"settings": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString,
						},
						"value": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}
