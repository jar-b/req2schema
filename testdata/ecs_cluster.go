package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capacityProviders": {
				Type:  schema.TypeList,
				Elems: &schema.Schema{Type: schema.TypeString},
			},
			"clusterName": {
				Type: schema.TypeString,
			},
			"configuration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"executeCommandConfiguration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kmsKeyId": {
										Type: schema.TypeString,
									},
									"logConfiguration": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloudWatchEncryptionEnabled": {
													Type: schema.TypeBool,
												},
												"cloudWatchLogGroupName": {
													Type: schema.TypeString,
												},
												"s3BucketName": {
													Type: schema.TypeString,
												},
												"s3EncryptionEnabled": {
													Type: schema.TypeBool,
												},
												"s3KeyPrefix": {
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
			"defaultCapacityProviderStrategy": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base": {
							Type: schema.TypeFloat,
						},
						"capacityProvider": {
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
