package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"findings": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type: schema.TypeString,
						},
						"arn": {
							Type: schema.TypeString,
						},
						"confidence": {
							Type: schema.TypeFloat,
						},
						"created_at": {
							Type: schema.TypeString,
						},
						"description": {
							Type: schema.TypeString,
						},
						"id": {
							Type: schema.TypeString,
						},
						"partition": {
							Type: schema.TypeString,
						},
						"region": {
							Type: schema.TypeString,
						},
						"resource": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_key_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"access_key_id": {
													Type: schema.TypeString,
												},
												"principal_id": {
													Type: schema.TypeString,
												},
												"user_name": {
													Type: schema.TypeString,
												},
												"user_type": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"container_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"container_runtime": {
													Type: schema.TypeString,
												},
												"id": {
													Type: schema.TypeString,
												},
												"image": {
													Type: schema.TypeString,
												},
												"image_prefix": {
													Type: schema.TypeString,
												},
												"name": {
													Type: schema.TypeString,
												},
												"security_context": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"privileged": {
																Type: schema.TypeBool,
															},
														},
													},
												},
												"volume_mounts": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mount_path": {
																Type: schema.TypeString,
															},
															"name": {
																Type: schema.TypeString,
															},
														},
													},
												},
											},
										},
									},
									"ebs_volume_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"scanned_volume_details": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"device_name": {
																Type: schema.TypeString,
															},
															"encryption_type": {
																Type: schema.TypeString,
															},
															"kms_key_arn": {
																Type: schema.TypeString,
															},
															"snapshot_arn": {
																Type: schema.TypeString,
															},
															"volume_arn": {
																Type: schema.TypeString,
															},
															"volume_size_in_gb": {
																Type: schema.TypeFloat,
															},
															"volume_type": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"skipped_volume_details": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"device_name": {
																Type: schema.TypeString,
															},
															"encryption_type": {
																Type: schema.TypeString,
															},
															"kms_key_arn": {
																Type: schema.TypeString,
															},
															"snapshot_arn": {
																Type: schema.TypeString,
															},
															"volume_arn": {
																Type: schema.TypeString,
															},
															"volume_size_in_gb": {
																Type: schema.TypeFloat,
															},
															"volume_type": {
																Type: schema.TypeString,
															},
														},
													},
												},
											},
										},
									},
									"ecs_cluster_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"active_services_count": {
													Type: schema.TypeFloat,
												},
												"arn": {
													Type: schema.TypeString,
												},
												"name": {
													Type: schema.TypeString,
												},
												"registered_container_instances_count": {
													Type: schema.TypeFloat,
												},
												"running_tasks_count": {
													Type: schema.TypeFloat,
												},
												"status": {
													Type: schema.TypeString,
												},
												"tags": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type: schema.TypeString,
															},
															"value": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"task_details": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"arn": {
																Type: schema.TypeString,
															},
															"containers": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"container_runtime": {
																			Type: schema.TypeString,
																		},
																		"id": {
																			Type: schema.TypeString,
																		},
																		"image": {
																			Type: schema.TypeString,
																		},
																		"image_prefix": {
																			Type: schema.TypeString,
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																		"security_context": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"privileged": {
																						Type: schema.TypeBool,
																					},
																				},
																			},
																		},
																		"volume_mounts": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mount_path": {
																						Type: schema.TypeString,
																					},
																					"name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"created_at": {
																Type: schema.TypeFloat,
															},
															"definition_arn": {
																Type: schema.TypeString,
															},
															"group": {
																Type: schema.TypeString,
															},
															"started_at": {
																Type: schema.TypeFloat,
															},
															"started_by": {
																Type: schema.TypeString,
															},
															"tags": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": {
																			Type: schema.TypeString,
																		},
																		"value": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"version": {
																Type: schema.TypeString,
															},
															"volumes": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"host_path": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"path": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"eks_cluster_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"arn": {
													Type: schema.TypeString,
												},
												"created_at": {
													Type: schema.TypeFloat,
												},
												"name": {
													Type: schema.TypeString,
												},
												"status": {
													Type: schema.TypeString,
												},
												"tags": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type: schema.TypeString,
															},
															"value": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"vpc_id": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"instance_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"availability_zone": {
													Type: schema.TypeString,
												},
												"iam_instance_profile": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"arn": {
																Type: schema.TypeString,
															},
															"id": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"image_description": {
													Type: schema.TypeString,
												},
												"image_id": {
													Type: schema.TypeString,
												},
												"instance_id": {
													Type: schema.TypeString,
												},
												"instance_state": {
													Type: schema.TypeString,
												},
												"instance_type": {
													Type: schema.TypeString,
												},
												"launch_time": {
													Type: schema.TypeString,
												},
												"network_interfaces": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_addresses": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"network_interface_id": {
																Type: schema.TypeString,
															},
															"private_dns_name": {
																Type: schema.TypeString,
															},
															"private_ip_address": {
																Type: schema.TypeString,
															},
															"private_ip_addresses": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"private_dns_name": {
																			Type: schema.TypeString,
																		},
																		"private_ip_address": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"public_dns_name": {
																Type: schema.TypeString,
															},
															"public_ip": {
																Type: schema.TypeString,
															},
															"security_groups": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"group_id": {
																			Type: schema.TypeString,
																		},
																		"group_name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"subnet_id": {
																Type: schema.TypeString,
															},
															"vpc_id": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"outpost_arn": {
													Type: schema.TypeString,
												},
												"platform": {
													Type: schema.TypeString,
												},
												"product_codes": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"product_code_id": {
																Type: schema.TypeString,
															},
															"product_code_type": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"tags": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type: schema.TypeString,
															},
															"value": {
																Type: schema.TypeString,
															},
														},
													},
												},
											},
										},
									},
									"kubernetes_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kubernetes_user_details": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"groups": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"uid": {
																Type: schema.TypeString,
															},
															"username": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"kubernetes_workload_details": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"containers": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"container_runtime": {
																			Type: schema.TypeString,
																		},
																		"id": {
																			Type: schema.TypeString,
																		},
																		"image": {
																			Type: schema.TypeString,
																		},
																		"image_prefix": {
																			Type: schema.TypeString,
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																		"security_context": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"privileged": {
																						Type: schema.TypeBool,
																					},
																				},
																			},
																		},
																		"volume_mounts": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mount_path": {
																						Type: schema.TypeString,
																					},
																					"name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"host_network": {
																Type: schema.TypeBool,
															},
															"name": {
																Type: schema.TypeString,
															},
															"namespace": {
																Type: schema.TypeString,
															},
															"type": {
																Type: schema.TypeString,
															},
															"uid": {
																Type: schema.TypeString,
															},
															"volumes": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"host_path": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"path": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"lambda_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"description": {
													Type: schema.TypeString,
												},
												"function_arn": {
													Type: schema.TypeString,
												},
												"function_name": {
													Type: schema.TypeString,
												},
												"function_version": {
													Type: schema.TypeString,
												},
												"last_modified_at": {
													Type: schema.TypeFloat,
												},
												"revision_id": {
													Type: schema.TypeString,
												},
												"role": {
													Type: schema.TypeString,
												},
												"tags": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type: schema.TypeString,
															},
															"value": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"vpc_config": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"security_groups": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"group_id": {
																			Type: schema.TypeString,
																		},
																		"group_name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"subnet_ids": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"vpc_id": {
																Type: schema.TypeString,
															},
														},
													},
												},
											},
										},
									},
									"rds_db_instance_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"db_cluster_identifier": {
													Type: schema.TypeString,
												},
												"db_instance_arn": {
													Type: schema.TypeString,
												},
												"db_instance_identifier": {
													Type: schema.TypeString,
												},
												"engine": {
													Type: schema.TypeString,
												},
												"engine_version": {
													Type: schema.TypeString,
												},
												"tags": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type: schema.TypeString,
															},
															"value": {
																Type: schema.TypeString,
															},
														},
													},
												},
											},
										},
									},
									"rds_db_user_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"application": {
													Type: schema.TypeString,
												},
												"auth_method": {
													Type: schema.TypeString,
												},
												"database": {
													Type: schema.TypeString,
												},
												"ssl": {
													Type: schema.TypeString,
												},
												"user": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"resource_type": {
										Type: schema.TypeString,
									},
									"s3_bucket_details": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"arn": {
													Type: schema.TypeString,
												},
												"created_at": {
													Type: schema.TypeFloat,
												},
												"default_server_side_encryption": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"encryption_type": {
																Type: schema.TypeString,
															},
															"kms_master_key_arn": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"name": {
													Type: schema.TypeString,
												},
												"owner": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"public_access": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effective_permission": {
																Type: schema.TypeString,
															},
															"permission_configuration": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"account_level_permissions": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"block_public_access": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"block_public_acls": {
																									Type: schema.TypeBool,
																								},
																								"block_public_policy": {
																									Type: schema.TypeBool,
																								},
																								"ignore_public_acls": {
																									Type: schema.TypeBool,
																								},
																								"restrict_public_buckets": {
																									Type: schema.TypeBool,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"bucket_level_permissions": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_control_list": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"allows_public_read_access": {
																									Type: schema.TypeBool,
																								},
																								"allows_public_write_access": {
																									Type: schema.TypeBool,
																								},
																							},
																						},
																					},
																					"block_public_access": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"block_public_acls": {
																									Type: schema.TypeBool,
																								},
																								"block_public_policy": {
																									Type: schema.TypeBool,
																								},
																								"ignore_public_acls": {
																									Type: schema.TypeBool,
																								},
																								"restrict_public_buckets": {
																									Type: schema.TypeBool,
																								},
																							},
																						},
																					},
																					"bucket_policy": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"allows_public_read_access": {
																									Type: schema.TypeBool,
																								},
																								"allows_public_write_access": {
																									Type: schema.TypeBool,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												"tags": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type: schema.TypeString,
															},
															"value": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"type": {
													Type: schema.TypeString,
												},
											},
										},
									},
								},
							},
						},
						"schema_version": {
							Type: schema.TypeString,
						},
						"service": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action_type": {
													Type: schema.TypeString,
												},
												"aws_api_call_action": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"affected_resources": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"string": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"api": {
																Type: schema.TypeString,
															},
															"caller_type": {
																Type: schema.TypeString,
															},
															"domain_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"domain": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"error_code": {
																Type: schema.TypeString,
															},
															"remote_account_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"account_id": {
																			Type: schema.TypeString,
																		},
																		"affiliated": {
																			Type: schema.TypeBool,
																		},
																	},
																},
															},
															"remote_ip_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"city": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"city_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"country": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"country_code": {
																						Type: schema.TypeString,
																					},
																					"country_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"geo_location": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"lat": {
																						Type: schema.TypeFloat,
																					},
																					"lon": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"ip_address_v4": {
																			Type: schema.TypeString,
																		},
																		"organization": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"asn": {
																						Type: schema.TypeString,
																					},
																					"asn_org": {
																						Type: schema.TypeString,
																					},
																					"isp": {
																						Type: schema.TypeString,
																					},
																					"org": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"service_name": {
																Type: schema.TypeString,
															},
															"user_agent": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"dns_request_action": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"blocked": {
																Type: schema.TypeBool,
															},
															"domain": {
																Type: schema.TypeString,
															},
															"protocol": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"kubernetes_api_call_action": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameters": {
																Type: schema.TypeString,
															},
															"remote_ip_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"city": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"city_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"country": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"country_code": {
																						Type: schema.TypeString,
																					},
																					"country_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"geo_location": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"lat": {
																						Type: schema.TypeFloat,
																					},
																					"lon": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"ip_address_v4": {
																			Type: schema.TypeString,
																		},
																		"organization": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"asn": {
																						Type: schema.TypeString,
																					},
																					"asn_org": {
																						Type: schema.TypeString,
																					},
																					"isp": {
																						Type: schema.TypeString,
																					},
																					"org": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"request_uri": {
																Type: schema.TypeString,
															},
															"source_ips": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"status_code": {
																Type: schema.TypeFloat,
															},
															"user_agent": {
																Type: schema.TypeString,
															},
															"verb": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"network_connection_action": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"blocked": {
																Type: schema.TypeBool,
															},
															"connection_direction": {
																Type: schema.TypeString,
															},
															"local_ip_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ip_address_v4": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"local_port_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"port": {
																			Type: schema.TypeFloat,
																		},
																		"port_name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"protocol": {
																Type: schema.TypeString,
															},
															"remote_ip_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"city": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"city_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"country": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"country_code": {
																						Type: schema.TypeString,
																					},
																					"country_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"geo_location": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"lat": {
																						Type: schema.TypeFloat,
																					},
																					"lon": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"ip_address_v4": {
																			Type: schema.TypeString,
																		},
																		"organization": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"asn": {
																						Type: schema.TypeString,
																					},
																					"asn_org": {
																						Type: schema.TypeString,
																					},
																					"isp": {
																						Type: schema.TypeString,
																					},
																					"org": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"remote_port_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"port": {
																			Type: schema.TypeFloat,
																		},
																		"port_name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
												"port_probe_action": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"blocked": {
																Type: schema.TypeBool,
															},
															"port_probe_details": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"local_ip_details": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"ip_address_v4": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"local_port_details": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"port": {
																						Type: schema.TypeFloat,
																					},
																					"port_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"remote_ip_details": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"city": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"city_name": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																					"country": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"country_code": {
																									Type: schema.TypeString,
																								},
																								"country_name": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																					"geo_location": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"lat": {
																									Type: schema.TypeFloat,
																								},
																								"lon": {
																									Type: schema.TypeFloat,
																								},
																							},
																						},
																					},
																					"ip_address_v4": {
																						Type: schema.TypeString,
																					},
																					"organization": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"asn": {
																									Type: schema.TypeString,
																								},
																								"asn_org": {
																									Type: schema.TypeString,
																								},
																								"isp": {
																									Type: schema.TypeString,
																								},
																								"org": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												"rds_login_attempt_action": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"login_attributes": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"application": {
																			Type: schema.TypeString,
																		},
																		"failed_login_attempts": {
																			Type: schema.TypeFloat,
																		},
																		"successful_login_attempts": {
																			Type: schema.TypeFloat,
																		},
																		"user": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"remote_ip_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"city": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"city_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"country": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"country_code": {
																						Type: schema.TypeString,
																					},
																					"country_name": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"geo_location": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"lat": {
																						Type: schema.TypeFloat,
																					},
																					"lon": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"ip_address_v4": {
																			Type: schema.TypeString,
																		},
																		"organization": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"asn": {
																						Type: schema.TypeString,
																					},
																					"asn_org": {
																						Type: schema.TypeString,
																					},
																					"isp": {
																						Type: schema.TypeString,
																					},
																					"org": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"additional_info": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString,
												},
												"value": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"archived": {
										Type: schema.TypeBool,
									},
									"count": {
										Type: schema.TypeFloat,
									},
									"detector_id": {
										Type: schema.TypeString,
									},
									"ebs_volume_scan_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"scan_completed_at": {
													Type: schema.TypeFloat,
												},
												"scan_detections": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"highest_severity_threat_details": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"count": {
																			Type: schema.TypeFloat,
																		},
																		"severity": {
																			Type: schema.TypeString,
																		},
																		"threat_name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"scanned_item_count": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"files": {
																			Type: schema.TypeFloat,
																		},
																		"total_gb": {
																			Type: schema.TypeFloat,
																		},
																		"volumes": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"threat_detected_by_name": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"item_count": {
																			Type: schema.TypeFloat,
																		},
																		"shortened": {
																			Type: schema.TypeBool,
																		},
																		"threat_names": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"file_paths": {
																						Type: schema.TypeList,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"file_name": {
																									Type: schema.TypeString,
																								},
																								"file_path": {
																									Type: schema.TypeString,
																								},
																								"hash": {
																									Type: schema.TypeString,
																								},
																								"volume_arn": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																					"item_count": {
																						Type: schema.TypeFloat,
																					},
																					"name": {
																						Type: schema.TypeString,
																					},
																					"severity": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"unique_threat_name_count": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"threats_detected_item_count": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"files": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
														},
													},
												},
												"scan_id": {
													Type: schema.TypeString,
												},
												"scan_started_at": {
													Type: schema.TypeFloat,
												},
												"scan_type": {
													Type: schema.TypeString,
												},
												"sources": {
													Type:  schema.TypeList,
													Elems: &schema.Schema{Type: schema.TypeString},
												},
												"trigger_finding_id": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"event_first_seen": {
										Type: schema.TypeString,
									},
									"event_last_seen": {
										Type: schema.TypeString,
									},
									"evidence": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threat_intelligence_details": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"threat_list_name": {
																Type: schema.TypeString,
															},
															"threat_names": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
														},
													},
												},
											},
										},
									},
									"feature_name": {
										Type: schema.TypeString,
									},
									"resource_role": {
										Type: schema.TypeString,
									},
									"runtime_details": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"context": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address_family": {
																Type: schema.TypeString,
															},
															"file_system_type": {
																Type: schema.TypeString,
															},
															"flags": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"iana_protocol_number": {
																Type: schema.TypeFloat,
															},
															"ld_preload_value": {
																Type: schema.TypeString,
															},
															"library_path": {
																Type: schema.TypeString,
															},
															"memory_regions": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"modified_at": {
																Type: schema.TypeFloat,
															},
															"modifying_process": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"euid": {
																			Type: schema.TypeFloat,
																		},
																		"executable_path": {
																			Type: schema.TypeString,
																		},
																		"executable_sha256": {
																			Type: schema.TypeString,
																		},
																		"lineage": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"euid": {
																						Type: schema.TypeFloat,
																					},
																					"executable_path": {
																						Type: schema.TypeString,
																					},
																					"name": {
																						Type: schema.TypeString,
																					},
																					"namespace_pid": {
																						Type: schema.TypeFloat,
																					},
																					"parent_uuid": {
																						Type: schema.TypeString,
																					},
																					"pid": {
																						Type: schema.TypeFloat,
																					},
																					"start_time": {
																						Type: schema.TypeFloat,
																					},
																					"user_id": {
																						Type: schema.TypeFloat,
																					},
																					"uuid": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																		"namespace_pid": {
																			Type: schema.TypeFloat,
																		},
																		"parent_uuid": {
																			Type: schema.TypeString,
																		},
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																		"pwd": {
																			Type: schema.TypeString,
																		},
																		"start_time": {
																			Type: schema.TypeFloat,
																		},
																		"user": {
																			Type: schema.TypeString,
																		},
																		"user_id": {
																			Type: schema.TypeFloat,
																		},
																		"uuid": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"module_file_path": {
																Type: schema.TypeString,
															},
															"module_name": {
																Type: schema.TypeString,
															},
															"module_sha256": {
																Type: schema.TypeString,
															},
															"mount_source": {
																Type: schema.TypeString,
															},
															"mount_target": {
																Type: schema.TypeString,
															},
															"release_agent_path": {
																Type: schema.TypeString,
															},
															"runc_binary_path": {
																Type: schema.TypeString,
															},
															"script_path": {
																Type: schema.TypeString,
															},
															"shell_history_file_path": {
																Type: schema.TypeString,
															},
															"socket_path": {
																Type: schema.TypeString,
															},
															"target_process": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"euid": {
																			Type: schema.TypeFloat,
																		},
																		"executable_path": {
																			Type: schema.TypeString,
																		},
																		"executable_sha256": {
																			Type: schema.TypeString,
																		},
																		"lineage": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"euid": {
																						Type: schema.TypeFloat,
																					},
																					"executable_path": {
																						Type: schema.TypeString,
																					},
																					"name": {
																						Type: schema.TypeString,
																					},
																					"namespace_pid": {
																						Type: schema.TypeFloat,
																					},
																					"parent_uuid": {
																						Type: schema.TypeString,
																					},
																					"pid": {
																						Type: schema.TypeFloat,
																					},
																					"start_time": {
																						Type: schema.TypeFloat,
																					},
																					"user_id": {
																						Type: schema.TypeFloat,
																					},
																					"uuid": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																		"namespace_pid": {
																			Type: schema.TypeFloat,
																		},
																		"parent_uuid": {
																			Type: schema.TypeString,
																		},
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																		"pwd": {
																			Type: schema.TypeString,
																		},
																		"start_time": {
																			Type: schema.TypeFloat,
																		},
																		"user": {
																			Type: schema.TypeString,
																		},
																		"user_id": {
																			Type: schema.TypeFloat,
																		},
																		"uuid": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
												"process": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"euid": {
																Type: schema.TypeFloat,
															},
															"executable_path": {
																Type: schema.TypeString,
															},
															"executable_sha256": {
																Type: schema.TypeString,
															},
															"lineage": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"euid": {
																			Type: schema.TypeFloat,
																		},
																		"executable_path": {
																			Type: schema.TypeString,
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																		"namespace_pid": {
																			Type: schema.TypeFloat,
																		},
																		"parent_uuid": {
																			Type: schema.TypeString,
																		},
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																		"start_time": {
																			Type: schema.TypeFloat,
																		},
																		"user_id": {
																			Type: schema.TypeFloat,
																		},
																		"uuid": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"name": {
																Type: schema.TypeString,
															},
															"namespace_pid": {
																Type: schema.TypeFloat,
															},
															"parent_uuid": {
																Type: schema.TypeString,
															},
															"pid": {
																Type: schema.TypeFloat,
															},
															"pwd": {
																Type: schema.TypeString,
															},
															"start_time": {
																Type: schema.TypeFloat,
															},
															"user": {
																Type: schema.TypeString,
															},
															"user_id": {
																Type: schema.TypeFloat,
															},
															"uuid": {
																Type: schema.TypeString,
															},
														},
													},
												},
											},
										},
									},
									"service_name": {
										Type: schema.TypeString,
									},
									"user_feedback": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"severity": {
							Type: schema.TypeFloat,
						},
						"title": {
							Type: schema.TypeString,
						},
						"type": {
							Type: schema.TypeString,
						},
						"updated_at": {
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
