package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cdi_input_specification": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resolution": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
					},
				},
			},
			"channel_class": {
				Type: schema.TypeString, // TODO: enum, add validation
			},
			"destinations": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeString,
						},
						"media_package_settings": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"channel_id": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"multiplex_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multiplex_id": {
										Type: schema.TypeString,
									},
									"program_name": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"settings": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password_param": {
										Type: schema.TypeString,
									},
									"stream_name": {
										Type: schema.TypeString,
									},
									"url": {
										Type: schema.TypeString,
									},
									"username": {
										Type: schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			"encoder_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"audio_descriptions": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audio_normalization_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"algorithm_control": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"target_lkfs": {
													Type: schema.TypeFloat,
												},
											},
										},
									},
									"audio_selector_name": {
										Type: schema.TypeString,
									},
									"audio_type": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"audio_type_control": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"audio_watermarking_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nielsen_watermarks_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nielsen_cbet_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cbet_check_digit_string": {
																			Type: schema.TypeString,
																		},
																		"cbet_stepaside": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"csid": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"nielsen_distribution_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"nielsen_naes_ii_nw_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"check_digit_string": {
																			Type: schema.TypeString,
																		},
																		"sid": {
																			Type: schema.TypeFloat,
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
									"codec_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aac_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"coding_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"input_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"rate_control_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"raw_format": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sample_rate": {
																Type: schema.TypeFloat,
															},
															"spec": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"vbr_quality": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"ac3_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"bitstream_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"coding_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"dialnorm": {
																Type: schema.TypeInt,
															},
															"drc_profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lfe_filter": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"metadata_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"eac3_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"attenuation_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"bitstream_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"coding_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"dc_filter": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"dialnorm": {
																Type: schema.TypeInt,
															},
															"drc_line": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"drc_rf": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lfe_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lfe_filter": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lo_ro_center_mix_level": {
																Type: schema.TypeFloat,
															},
															"lo_ro_surround_mix_level": {
																Type: schema.TypeFloat,
															},
															"lt_rt_center_mix_level": {
																Type: schema.TypeFloat,
															},
															"lt_rt_surround_mix_level": {
																Type: schema.TypeFloat,
															},
															"metadata_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"passthrough_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"phase_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"stereo_downmix": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"surround_ex_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"surround_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"mp2_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"coding_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sample_rate": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"pass_through_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"wav_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bit_depth": {
																Type: schema.TypeFloat,
															},
															"coding_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sample_rate": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
											},
										},
									},
									"language_code": {
										Type: schema.TypeString,
									},
									"language_code_control": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"name": {
										Type: schema.TypeString,
									},
									"remix_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"channel_mappings": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"input_channel_levels": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"gain": {
																			Type: schema.TypeInt,
																		},
																		"input_channel": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"output_channel": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"channels_in": {
													Type: schema.TypeInt,
												},
												"channels_out": {
													Type: schema.TypeInt,
												},
											},
										},
									},
									"stream_name": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"avail_blanking": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"avail_blanking_image": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password_param": {
													Type: schema.TypeString,
												},
												"uri": {
													Type: schema.TypeString,
												},
												"username": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"state": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"avail_configuration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"avail_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"scte35_splice_insert": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ad_avail_offset": {
																Type: schema.TypeInt,
															},
															"no_regional_blackout_flag": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"web_delivery_allowed_flag": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"scte35_time_signal_apos": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ad_avail_offset": {
																Type: schema.TypeInt,
															},
															"no_regional_blackout_flag": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"web_delivery_allowed_flag": {
																Type: schema.TypeString, // TODO: enum, add validation
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
						"blackout_slate": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blackout_slate_image": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password_param": {
													Type: schema.TypeString,
												},
												"uri": {
													Type: schema.TypeString,
												},
												"username": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"network_end_blackout": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"network_end_blackout_image": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password_param": {
													Type: schema.TypeString,
												},
												"uri": {
													Type: schema.TypeString,
												},
												"username": {
													Type: schema.TypeString,
												},
											},
										},
									},
									"network_id": {
										Type: schema.TypeString,
									},
									"state": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"caption_descriptions": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"caption_selector_name": {
										Type: schema.TypeString,
									},
									"destination_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"arib_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"burn_in_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"background_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"background_opacity": {
																Type: schema.TypeInt,
															},
															"font": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"password_param": {
																			Type: schema.TypeString,
																		},
																		"uri": {
																			Type: schema.TypeString,
																		},
																		"username": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"font_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"font_opacity": {
																Type: schema.TypeInt,
															},
															"font_resolution": {
																Type: schema.TypeInt,
															},
															"font_size": {
																Type: schema.TypeString,
															},
															"outline_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"outline_size": {
																Type: schema.TypeInt,
															},
															"shadow_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"shadow_opacity": {
																Type: schema.TypeInt,
															},
															"shadow_x_offset": {
																Type: schema.TypeInt,
															},
															"shadow_y_offset": {
																Type: schema.TypeInt,
															},
															"teletext_grid_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"x_position": {
																Type: schema.TypeInt,
															},
															"y_position": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"dvb_sub_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"background_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"background_opacity": {
																Type: schema.TypeInt,
															},
															"font": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"password_param": {
																			Type: schema.TypeString,
																		},
																		"uri": {
																			Type: schema.TypeString,
																		},
																		"username": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"font_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"font_opacity": {
																Type: schema.TypeInt,
															},
															"font_resolution": {
																Type: schema.TypeInt,
															},
															"font_size": {
																Type: schema.TypeString,
															},
															"outline_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"outline_size": {
																Type: schema.TypeInt,
															},
															"shadow_color": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"shadow_opacity": {
																Type: schema.TypeInt,
															},
															"shadow_x_offset": {
																Type: schema.TypeInt,
															},
															"shadow_y_offset": {
																Type: schema.TypeInt,
															},
															"teletext_grid_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"x_position": {
																Type: schema.TypeInt,
															},
															"y_position": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"ebu_tt_d_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"copyright_holder": {
																Type: schema.TypeString,
															},
															"fill_line_gap": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"font_family": {
																Type: schema.TypeString,
															},
															"style_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"embedded_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"embedded_plus_scte20_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"rtmp_caption_info_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"scte20_plus_embedded_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"scte27_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"smpte_tt_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"teletext_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"ttml_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"style_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"webvtt_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"style_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
											},
										},
									},
									"language_code": {
										Type: schema.TypeString,
									},
									"language_description": {
										Type: schema.TypeString,
									},
									"name": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"feature_activations": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"input_prepare_schedule_actions": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"global_configuration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initial_audio_gain": {
										Type: schema.TypeInt,
									},
									"input_end_action": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"input_loss_behavior": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"black_frame_msec": {
													Type: schema.TypeInt,
												},
												"input_loss_image_color": {
													Type: schema.TypeString,
												},
												"input_loss_image_slate": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"password_param": {
																Type: schema.TypeString,
															},
															"uri": {
																Type: schema.TypeString,
															},
															"username": {
																Type: schema.TypeString,
															},
														},
													},
												},
												"input_loss_image_type": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"repeat_frame_msec": {
													Type: schema.TypeInt,
												},
											},
										},
									},
									"output_locking_mode": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"output_timing_source": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"support_low_framerate_inputs": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"motion_graphics_configuration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"motion_graphics_insertion": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"motion_graphics_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"html_motion_graphics_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
											},
										},
									},
								},
							},
						},
						"nielsen_configuration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"distributor_id": {
										Type: schema.TypeString,
									},
									"nielsen_pcm_to_id3_tagging": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"output_groups": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString,
									},
									"output_group_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"archive_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"archive_cdn_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"archive_s3_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"canned_acl": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"log_uploads": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																				},
																			},
																		},
																	},
																},
															},
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination_ref_id": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"rollover_interval": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"frame_capture_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination_ref_id": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"frame_capture_cdn_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frame_capture_s3_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"canned_acl": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"log_uploads": {
																						Type: schema.TypeString, // TODO: enum, add validation
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
												"hls_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ad_markers": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"base_url_content": {
																Type: schema.TypeString,
															},
															"base_url_content1": {
																Type: schema.TypeString,
															},
															"base_url_manifest": {
																Type: schema.TypeString,
															},
															"base_url_manifest1": {
																Type: schema.TypeString,
															},
															"caption_language_mappings": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"caption_channel": {
																			Type: schema.TypeInt,
																		},
																		"language_code": {
																			Type: schema.TypeString,
																		},
																		"language_description": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"caption_language_setting": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"client_cache": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"codec_specification": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"constant_iv": {
																Type: schema.TypeString,
															},
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination_ref_id": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"directory_structure": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"discontinuity_tags": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"encryption_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"hls_cdn_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hls_akamai_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connection_retry_interval": {
																						Type: schema.TypeInt,
																					},
																					"filecache_duration": {
																						Type: schema.TypeInt,
																					},
																					"http_transfer_mode": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"num_retries": {
																						Type: schema.TypeInt,
																					},
																					"restart_delay": {
																						Type: schema.TypeInt,
																					},
																					"salt": {
																						Type: schema.TypeString,
																					},
																					"token": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"hls_basic_put_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connection_retry_interval": {
																						Type: schema.TypeInt,
																					},
																					"filecache_duration": {
																						Type: schema.TypeInt,
																					},
																					"num_retries": {
																						Type: schema.TypeInt,
																					},
																					"restart_delay": {
																						Type: schema.TypeInt,
																					},
																				},
																			},
																		},
																		"hls_media_store_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connection_retry_interval": {
																						Type: schema.TypeInt,
																					},
																					"filecache_duration": {
																						Type: schema.TypeInt,
																					},
																					"media_store_storage_class": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"num_retries": {
																						Type: schema.TypeInt,
																					},
																					"restart_delay": {
																						Type: schema.TypeInt,
																					},
																				},
																			},
																		},
																		"hls_s3_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"canned_acl": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"log_uploads": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																				},
																			},
																		},
																		"hls_webdav_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connection_retry_interval": {
																						Type: schema.TypeInt,
																					},
																					"filecache_duration": {
																						Type: schema.TypeInt,
																					},
																					"http_transfer_mode": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"num_retries": {
																						Type: schema.TypeInt,
																					},
																					"restart_delay": {
																						Type: schema.TypeInt,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"hls_id3_segment_tagging": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"i_frame_only_playlists": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"incomplete_segment_behavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"index_n_segments": {
																Type: schema.TypeInt,
															},
															"input_loss_action": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"iv_in_manifest": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"iv_source": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"keep_segments": {
																Type: schema.TypeInt,
															},
															"key_format": {
																Type: schema.TypeString,
															},
															"key_format_versions": {
																Type: schema.TypeString,
															},
															"key_provider_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"static_key_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key_provider_server": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"password_param": {
																									Type: schema.TypeString,
																								},
																								"uri": {
																									Type: schema.TypeString,
																								},
																								"username": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																					"static_key_value": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"manifest_compression": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"manifest_duration_format": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"min_segment_length": {
																Type: schema.TypeInt,
															},
															"mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"output_selection": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"program_date_time": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"program_date_time_clock": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"program_date_time_period": {
																Type: schema.TypeInt,
															},
															"redundant_manifest": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"segment_length": {
																Type: schema.TypeInt,
															},
															"segmentation_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"segments_per_subdirectory": {
																Type: schema.TypeInt,
															},
															"stream_inf_resolution": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timed_metadata_id3_frame": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timed_metadata_id3_period": {
																Type: schema.TypeInt,
															},
															"timestamp_delta_milliseconds": {
																Type: schema.TypeInt,
															},
															"ts_file_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"media_package_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination_ref_id": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
												"ms_smooth_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"acquisition_point_id": {
																Type: schema.TypeString,
															},
															"audio_only_timecode_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"certificate_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"connection_retry_interval": {
																Type: schema.TypeInt,
															},
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination_ref_id": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"event_id": {
																Type: schema.TypeString,
															},
															"event_id_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"event_stop_behavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"filecache_duration": {
																Type: schema.TypeInt,
															},
															"fragment_length": {
																Type: schema.TypeInt,
															},
															"input_loss_action": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"num_retries": {
																Type: schema.TypeInt,
															},
															"restart_delay": {
																Type: schema.TypeInt,
															},
															"segmentation_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"send_delay_ms": {
																Type: schema.TypeInt,
															},
															"sparse_track_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"stream_manifest_behavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timestamp_offset": {
																Type: schema.TypeString,
															},
															"timestamp_offset_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"multiplex_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"rtmp_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ad_markers": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"authentication_scheme": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"cache_full_behavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"cache_length": {
																Type: schema.TypeInt,
															},
															"caption_data": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"input_loss_action": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"restart_delay": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"udp_group_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"input_loss_action": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timed_metadata_id3_frame": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timed_metadata_id3_period": {
																Type: schema.TypeInt,
															},
														},
													},
												},
											},
										},
									},
									"outputs": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_description_names": {
													Type:  schema.TypeList,
													Elems: &schema.Schema{Type: schema.TypeString},
												},
												"caption_description_names": {
													Type:  schema.TypeList,
													Elems: &schema.Schema{Type: schema.TypeString},
												},
												"output_name": {
													Type: schema.TypeString,
												},
												"output_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"archive_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"container_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"m2ts_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"absent_input_audio_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"arib": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"arib_captions_pid": {
																									Type: schema.TypeString,
																								},
																								"arib_captions_pid_control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audio_buffer_model": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audio_frames_per_pes": {
																									Type: schema.TypeInt,
																								},
																								"audio_pids": {
																									Type: schema.TypeString,
																								},
																								"audio_stream_type": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"bitrate": {
																									Type: schema.TypeInt,
																								},
																								"buffer_model": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"cc_descriptor": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"dvb_nit_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"network_id": {
																												Type: schema.TypeInt,
																											},
																											"network_name": {
																												Type: schema.TypeString,
																											},
																											"rep_interval": {
																												Type: schema.TypeInt,
																											},
																										},
																									},
																								},
																								"dvb_sdt_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"output_sdt": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"rep_interval": {
																												Type: schema.TypeInt,
																											},
																											"service_name": {
																												Type: schema.TypeString,
																											},
																											"service_provider_name": {
																												Type: schema.TypeString,
																											},
																										},
																									},
																								},
																								"dvb_sub_pids": {
																									Type: schema.TypeString,
																								},
																								"dvb_tdt_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"rep_interval": {
																												Type: schema.TypeInt,
																											},
																										},
																									},
																								},
																								"dvb_teletext_pid": {
																									Type: schema.TypeString,
																								},
																								"ebif": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebp_audio_interval": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebp_lookahead_ms": {
																									Type: schema.TypeInt,
																								},
																								"ebp_placement": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ecm_pid": {
																									Type: schema.TypeString,
																								},
																								"es_rate_in_pes": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"etv_platform_pid": {
																									Type: schema.TypeString,
																								},
																								"etv_signal_pid": {
																									Type: schema.TypeString,
																								},
																								"fragment_time": {
																									Type: schema.TypeFloat,
																								},
																								"klv": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"klv_data_pids": {
																									Type: schema.TypeString,
																								},
																								"nielsen_id3_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"null_packet_bitrate": {
																									Type: schema.TypeFloat,
																								},
																								"pat_interval": {
																									Type: schema.TypeInt,
																								},
																								"pcr_control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"pcr_period": {
																									Type: schema.TypeInt,
																								},
																								"pcr_pid": {
																									Type: schema.TypeString,
																								},
																								"pmt_interval": {
																									Type: schema.TypeInt,
																								},
																								"pmt_pid": {
																									Type: schema.TypeString,
																								},
																								"program_num": {
																									Type: schema.TypeInt,
																								},
																								"rate_mode": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte27_pids": {
																									Type: schema.TypeString,
																								},
																								"scte35_control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte35_pid": {
																									Type: schema.TypeString,
																								},
																								"segmentation_markers": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentation_style": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentation_time": {
																									Type: schema.TypeFloat,
																								},
																								"timed_metadata_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"timed_metadata_pid": {
																									Type: schema.TypeString,
																								},
																								"transport_stream_id": {
																									Type: schema.TypeInt,
																								},
																								"video_pid": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																					"raw_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																					},
																				},
																			},
																		},
																		"extension": {
																			Type: schema.TypeString,
																		},
																		"name_modifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"frame_capture_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name_modifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"hls_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"h265_packaging_type": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"hls_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"audio_only_hls_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"audio_group_id": {
																									Type: schema.TypeString,
																								},
																								"audio_only_image": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"password_param": {
																												Type: schema.TypeString,
																											},
																											"uri": {
																												Type: schema.TypeString,
																											},
																											"username": {
																												Type: schema.TypeString,
																											},
																										},
																									},
																								},
																								"audio_track_type": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segment_type": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																							},
																						},
																					},
																					"fmp4_hls_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"audio_rendition_sets": {
																									Type: schema.TypeString,
																								},
																								"nielsen_id3_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"timed_metadata_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																							},
																						},
																					},
																					"frame_capture_hls_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																					},
																					"standard_hls_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"audio_rendition_sets": {
																									Type: schema.TypeString,
																								},
																								"m3u8_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"audio_frames_per_pes": {
																												Type: schema.TypeInt,
																											},
																											"audio_pids": {
																												Type: schema.TypeString,
																											},
																											"ecm_pid": {
																												Type: schema.TypeString,
																											},
																											"nielsen_id3_behavior": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"pat_interval": {
																												Type: schema.TypeInt,
																											},
																											"pcr_control": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"pcr_period": {
																												Type: schema.TypeInt,
																											},
																											"pcr_pid": {
																												Type: schema.TypeString,
																											},
																											"pmt_interval": {
																												Type: schema.TypeInt,
																											},
																											"pmt_pid": {
																												Type: schema.TypeString,
																											},
																											"program_num": {
																												Type: schema.TypeInt,
																											},
																											"scte35_behavior": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"scte35_pid": {
																												Type: schema.TypeString,
																											},
																											"timed_metadata_behavior": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"timed_metadata_pid": {
																												Type: schema.TypeString,
																											},
																											"transport_stream_id": {
																												Type: schema.TypeInt,
																											},
																											"video_pid": {
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
																		"name_modifier": {
																			Type: schema.TypeString,
																		},
																		"segment_modifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"media_package_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
															},
															"ms_smooth_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"h265_packaging_type": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"name_modifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"multiplex_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"destination_ref_id": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"rtmp_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"certificate_mode": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"connection_retry_interval": {
																			Type: schema.TypeInt,
																		},
																		"destination": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"destination_ref_id": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"num_retries": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"udp_output_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"buffer_msec": {
																			Type: schema.TypeInt,
																		},
																		"container_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"m2ts_settings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"absent_input_audio_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"arib": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"arib_captions_pid": {
																									Type: schema.TypeString,
																								},
																								"arib_captions_pid_control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audio_buffer_model": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audio_frames_per_pes": {
																									Type: schema.TypeInt,
																								},
																								"audio_pids": {
																									Type: schema.TypeString,
																								},
																								"audio_stream_type": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"bitrate": {
																									Type: schema.TypeInt,
																								},
																								"buffer_model": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"cc_descriptor": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"dvb_nit_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"network_id": {
																												Type: schema.TypeInt,
																											},
																											"network_name": {
																												Type: schema.TypeString,
																											},
																											"rep_interval": {
																												Type: schema.TypeInt,
																											},
																										},
																									},
																								},
																								"dvb_sdt_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"output_sdt": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"rep_interval": {
																												Type: schema.TypeInt,
																											},
																											"service_name": {
																												Type: schema.TypeString,
																											},
																											"service_provider_name": {
																												Type: schema.TypeString,
																											},
																										},
																									},
																								},
																								"dvb_sub_pids": {
																									Type: schema.TypeString,
																								},
																								"dvb_tdt_settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"rep_interval": {
																												Type: schema.TypeInt,
																											},
																										},
																									},
																								},
																								"dvb_teletext_pid": {
																									Type: schema.TypeString,
																								},
																								"ebif": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebp_audio_interval": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebp_lookahead_ms": {
																									Type: schema.TypeInt,
																								},
																								"ebp_placement": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ecm_pid": {
																									Type: schema.TypeString,
																								},
																								"es_rate_in_pes": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"etv_platform_pid": {
																									Type: schema.TypeString,
																								},
																								"etv_signal_pid": {
																									Type: schema.TypeString,
																								},
																								"fragment_time": {
																									Type: schema.TypeFloat,
																								},
																								"klv": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"klv_data_pids": {
																									Type: schema.TypeString,
																								},
																								"nielsen_id3_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"null_packet_bitrate": {
																									Type: schema.TypeFloat,
																								},
																								"pat_interval": {
																									Type: schema.TypeInt,
																								},
																								"pcr_control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"pcr_period": {
																									Type: schema.TypeInt,
																								},
																								"pcr_pid": {
																									Type: schema.TypeString,
																								},
																								"pmt_interval": {
																									Type: schema.TypeInt,
																								},
																								"pmt_pid": {
																									Type: schema.TypeString,
																								},
																								"program_num": {
																									Type: schema.TypeInt,
																								},
																								"rate_mode": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte27_pids": {
																									Type: schema.TypeString,
																								},
																								"scte35_control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte35_pid": {
																									Type: schema.TypeString,
																								},
																								"segmentation_markers": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentation_style": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentation_time": {
																									Type: schema.TypeFloat,
																								},
																								"timed_metadata_behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"timed_metadata_pid": {
																									Type: schema.TypeString,
																								},
																								"transport_stream_id": {
																									Type: schema.TypeInt,
																								},
																								"video_pid": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"destination": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"destination_ref_id": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"fec_output_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"column_depth": {
																						Type: schema.TypeInt,
																					},
																					"include_fec": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"row_length": {
																						Type: schema.TypeInt,
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
												"video_description_name": {
													Type: schema.TypeString,
												},
											},
										},
									},
								},
							},
						},
						"timecode_config": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"sync_threshold": {
										Type: schema.TypeInt,
									},
								},
							},
						},
						"video_descriptions": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"codec_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"frame_capture_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"capture_interval": {
																Type: schema.TypeInt,
															},
															"capture_interval_units": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"h264_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"afd_signaling": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"bitrate": {
																Type: schema.TypeInt,
															},
															"buf_fill_pct": {
																Type: schema.TypeInt,
															},
															"buf_size": {
																Type: schema.TypeInt,
															},
															"color_metadata": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"color_space_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"color_space_passthrough_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"rec601_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"rec709_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																	},
																},
															},
															"entropy_encoding": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"filter_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"temporal_filter_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"post_filter_sharpening": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"strength": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																				},
																			},
																		},
																	},
																},
															},
															"fixed_afd": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"flicker_aq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"force_field_pictures": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerate_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerate_denominator": {
																Type: schema.TypeInt,
															},
															"framerate_numerator": {
																Type: schema.TypeInt,
															},
															"gop_b_reference": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"gop_closed_cadence": {
																Type: schema.TypeInt,
															},
															"gop_num_b_frames": {
																Type: schema.TypeInt,
															},
															"gop_size": {
																Type: schema.TypeFloat,
															},
															"gop_size_units": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"level": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"look_ahead_rate_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"max_bitrate": {
																Type: schema.TypeInt,
															},
															"min_i_interval": {
																Type: schema.TypeInt,
															},
															"num_ref_frames": {
																Type: schema.TypeInt,
															},
															"par_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"par_denominator": {
																Type: schema.TypeInt,
															},
															"par_numerator": {
																Type: schema.TypeInt,
															},
															"profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"quality_level": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"qvbr_quality_level": {
																Type: schema.TypeInt,
															},
															"rate_control_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scan_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scene_change_detect": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"slices": {
																Type: schema.TypeInt,
															},
															"softness": {
																Type: schema.TypeInt,
															},
															"spatial_aq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"subgop_length": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"syntax": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"temporal_aq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timecode_insertion": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"h265_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"afd_signaling": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"alternative_transfer_function": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"bitrate": {
																Type: schema.TypeInt,
															},
															"buf_size": {
																Type: schema.TypeInt,
															},
															"color_metadata": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"color_space_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"color_space_passthrough_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"hdr10_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"max_cll": {
																						Type: schema.TypeInt,
																					},
																					"max_fall": {
																						Type: schema.TypeInt,
																					},
																				},
																			},
																		},
																		"rec601_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"rec709_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																	},
																},
															},
															"filter_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"temporal_filter_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"post_filter_sharpening": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"strength": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																				},
																			},
																		},
																	},
																},
															},
															"fixed_afd": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"flicker_aq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerate_denominator": {
																Type: schema.TypeInt,
															},
															"framerate_numerator": {
																Type: schema.TypeInt,
															},
															"gop_closed_cadence": {
																Type: schema.TypeInt,
															},
															"gop_size": {
																Type: schema.TypeFloat,
															},
															"gop_size_units": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"level": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"look_ahead_rate_control": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"max_bitrate": {
																Type: schema.TypeInt,
															},
															"min_i_interval": {
																Type: schema.TypeInt,
															},
															"par_denominator": {
																Type: schema.TypeInt,
															},
															"par_numerator": {
																Type: schema.TypeInt,
															},
															"profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"qvbr_quality_level": {
																Type: schema.TypeInt,
															},
															"rate_control_mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scan_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scene_change_detect": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"slices": {
																Type: schema.TypeInt,
															},
															"tier": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timecode_insertion": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"mpeg2_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"afd_signaling": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"color_metadata": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"color_space": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"display_aspect_ratio": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"filter_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"temporal_filter_settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"post_filter_sharpening": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"strength": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																				},
																			},
																		},
																	},
																},
															},
															"fixed_afd": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerate_denominator": {
																Type: schema.TypeInt,
															},
															"framerate_numerator": {
																Type: schema.TypeInt,
															},
															"gop_closed_cadence": {
																Type: schema.TypeInt,
															},
															"gop_num_b_frames": {
																Type: schema.TypeInt,
															},
															"gop_size": {
																Type: schema.TypeFloat,
															},
															"gop_size_units": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scan_type": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"subgop_length": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timecode_insertion": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
											},
										},
									},
									"height": {
										Type: schema.TypeInt,
									},
									"name": {
										Type: schema.TypeString,
									},
									"respond_to_afd": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"scaling_behavior": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"sharpness": {
										Type: schema.TypeInt,
									},
									"width": {
										Type: schema.TypeInt,
									},
								},
							},
						},
					},
				},
			},
			"input_attachments": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automatic_input_failover_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error_clear_time_msec": {
										Type: schema.TypeInt,
									},
									"failover_conditions": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"failover_condition_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"audio_silence_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audio_selector_name": {
																			Type: schema.TypeString,
																		},
																		"audio_silence_threshold_msec": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"input_loss_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"input_loss_threshold_msec": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"video_black_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"black_detect_threshold": {
																			Type: schema.TypeFloat,
																		},
																		"video_black_threshold_msec": {
																			Type: schema.TypeInt,
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
									"input_preference": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"secondary_input_id": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"input_attachment_name": {
							Type: schema.TypeString,
						},
						"input_id": {
							Type: schema.TypeString,
						},
						"input_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audio_selectors": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString,
												},
												"selector_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"audio_hls_rendition_selection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"group_id": {
																			Type: schema.TypeString,
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"audio_language_selection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"language_code": {
																			Type: schema.TypeString,
																		},
																		"language_selection_policy": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																	},
																},
															},
															"audio_pid_selection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"pid": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"audio_track_selection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"tracks": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"track": {
																						Type: schema.TypeInt,
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
									"caption_selectors": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"language_code": {
													Type: schema.TypeString,
												},
												"name": {
													Type: schema.TypeString,
												},
												"selector_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ancillary_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"source_ancillary_channel_number": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"arib_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
															},
															"dvb_sub_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ocr_language": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"pid": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"embedded_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"convert608_to708": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"scte20_detection": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"source608_channel_number": {
																			Type: schema.TypeInt,
																		},
																		"source608_track_number": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"scte20_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"convert608_to708": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"source608_channel_number": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"scte27_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ocr_language": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"pid": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"teletext_source_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"output_rectangle": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"height": {
																						Type: schema.TypeFloat,
																					},
																					"left_offset": {
																						Type: schema.TypeFloat,
																					},
																					"top_offset": {
																						Type: schema.TypeFloat,
																					},
																					"width": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"page_number": {
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
									"deblock_filter": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"denoise_filter": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"filter_strength": {
										Type: schema.TypeInt,
									},
									"input_filter": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"network_input_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hls_input_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bandwidth": {
																Type: schema.TypeInt,
															},
															"buffer_segments": {
																Type: schema.TypeInt,
															},
															"retries": {
																Type: schema.TypeInt,
															},
															"retry_interval": {
																Type: schema.TypeInt,
															},
															"scte35_source": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"server_validation": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
											},
										},
									},
									"smpte2038_data_preference": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"source_end_behavior": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"video_selector": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"color_space": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"color_space_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hdr10_settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"max_cll": {
																			Type: schema.TypeInt,
																		},
																		"max_fall": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
														},
													},
												},
												"color_space_usage": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"selector_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"video_selector_pid": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"pid": {
																			Type: schema.TypeInt,
																		},
																	},
																},
															},
															"video_selector_program_id": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"program_id": {
																			Type: schema.TypeInt,
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
					},
				},
			},
			"input_specification": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"codec": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
						"maximum_bitrate": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
						"resolution": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
					},
				},
			},
			"log_level": {
				Type: schema.TypeString, // TODO: enum, add validation
			},
			"maintenance": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maintenance_day": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
						"maintenance_start_time": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString,
			},
			"reserved": {
				Type: schema.TypeString,
			},
			"role_arn": {
				Type: schema.TypeString,
			},
			"vpc": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public_address_allocation_ids": {
							Type:  schema.TypeList,
							Elems: &schema.Schema{Type: schema.TypeString},
						},
						"security_group_ids": {
							Type:  schema.TypeList,
							Elems: &schema.Schema{Type: schema.TypeString},
						},
						"subnet_ids": {
							Type:  schema.TypeList,
							Elems: &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}
