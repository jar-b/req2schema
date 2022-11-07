package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capacity_provider_strategy": {
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
			"client_token": {
				Type: schema.TypeString,
			},
			"cluster": {
				Type: schema.TypeString,
			},
			"deployment_configuration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_circuit_breaker": {
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
						"maximum_percent": {
							Type: schema.TypeFloat,
						},
						"minimum_healthy_percent": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"deployment_controller": {
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
			"desired_count": {
				Type: schema.TypeFloat,
			},
			"enable_ecs_managed_tags": {
				Type: schema.TypeBool,
			},
			"enable_execute_command": {
				Type: schema.TypeBool,
			},
			"health_check_grace_period_seconds": {
				Type: schema.TypeFloat,
			},
			"launch_type": {
				Type: schema.TypeString,
			},
			"load_balancers": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"container_name": {
							Type: schema.TypeString,
						},
						"container_port": {
							Type: schema.TypeFloat,
						},
						"load_balancer_name": {
							Type: schema.TypeString,
						},
						"target_group_arn": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"network_configuration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awsvpc_configuration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"assign_public_ip": {
										Type: schema.TypeString,
									},
									"security_groups": {
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
			"placement_constraints": {
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
			"placement_strategy": {
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
			"platform_version": {
				Type: schema.TypeString,
			},
			"propagate_tags": {
				Type: schema.TypeString,
			},
			"role": {
				Type: schema.TypeString,
			},
			"scheduling_strategy": {
				Type: schema.TypeString,
			},
			"service_name": {
				Type: schema.TypeString,
			},
			"service_registries": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"container_name": {
							Type: schema.TypeString,
						},
						"container_port": {
							Type: schema.TypeFloat,
						},
						"port": {
							Type: schema.TypeFloat,
						},
						"registry_arn": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"task_definition": {
				Type: schema.TypeString,
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}
