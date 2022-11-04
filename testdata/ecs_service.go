package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capacityProviderStrategy": {
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
			"clientToken": {
				Type: schema.TypeString,
			},
			"cluster": {
				Type: schema.TypeString,
			},
			"deploymentConfiguration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deploymentCircuitBreaker": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeBool,
									},
									"rollback": {
										Type: schema.TypeBool,
									},
								},
							},
						},
						"maximumPercent": {
							Type: schema.TypeFloat,
						},
						"minimumHealthyPercent": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"deploymentController": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"desiredCount": {
				Type: schema.TypeFloat,
			},
			"enableECSManagedTags": {
				Type: schema.TypeBool,
			},
			"enableExecuteCommand": {
				Type: schema.TypeBool,
			},
			"healthCheckGracePeriodSeconds": {
				Type: schema.TypeFloat,
			},
			"launchType": {
				Type: schema.TypeString,
			},
			"loadBalancers": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"containerName": {
							Type: schema.TypeString,
						},
						"containerPort": {
							Type: schema.TypeFloat,
						},
						"loadBalancerName": {
							Type: schema.TypeString,
						},
						"targetGroupArn": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"networkConfiguration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awsvpcConfiguration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"assignPublicIp": {
										Type: schema.TypeString,
									},
									"securityGroups": {
										Type:  schema.TypeList,
										Elems: &schema.Schema{Type: schema.TypeString},
									},
									"subnets": {
										Type:  schema.TypeList,
										Elems: &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"placementConstraints": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type: schema.TypeString,
						},
						"type": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"placementStrategy": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type: schema.TypeString,
						},
						"type": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"platformVersion": {
				Type: schema.TypeString,
			},
			"propagateTags": {
				Type: schema.TypeString,
			},
			"role": {
				Type: schema.TypeString,
			},
			"schedulingStrategy": {
				Type: schema.TypeString,
			},
			"serviceName": {
				Type: schema.TypeString,
			},
			"serviceRegistries": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"containerName": {
							Type: schema.TypeString,
						},
						"containerPort": {
							Type: schema.TypeFloat,
						},
						"port": {
							Type: schema.TypeFloat,
						},
						"registryArn": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"taskDefinition": {
				Type: schema.TypeString,
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}
