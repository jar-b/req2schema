package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cdiInputSpecification": {
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
			"channelClass": {
				Type: schema.TypeString, // TODO: enum, add validation
			},
			"destinations": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeString,
						},
						"mediaPackageSettings": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"channelId": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"multiplexSettings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multiplexId": {
										Type: schema.TypeString,
									},
									"programName": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"settings": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"passwordParam": {
										Type: schema.TypeString,
									},
									"streamName": {
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
			"encoderSettings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"audioDescriptions": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audioNormalizationSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"algorithmControl": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"targetLkfs": {
													Type: schema.TypeFloat,
												},
											},
										},
									},
									"audioSelectorName": {
										Type: schema.TypeString,
									},
									"audioType": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"audioTypeControl": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"audioWatermarkingSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nielsenWatermarksSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nielsenCbetSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cbetCheckDigitString": {
																			Type: schema.TypeString,
																		},
																		"cbetStepaside": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"csid": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"nielsenDistributionType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"nielsenNaesIiNwSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"checkDigitString": {
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
									"codecSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aacSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"codingMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"inputType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"rateControlMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"rawFormat": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sampleRate": {
																Type: schema.TypeFloat,
															},
															"spec": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"vbrQuality": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"ac3Settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"bitstreamMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"codingMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"dialnorm": {
																Type: schema.TypeFloat,
															},
															"drcProfile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lfeFilter": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"metadataControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"eac3Settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"attenuationControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"bitstreamMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"codingMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"dcFilter": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"dialnorm": {
																Type: schema.TypeFloat,
															},
															"drcLine": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"drcRf": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lfeControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lfeFilter": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"loRoCenterMixLevel": {
																Type: schema.TypeFloat,
															},
															"loRoSurroundMixLevel": {
																Type: schema.TypeFloat,
															},
															"ltRtCenterMixLevel": {
																Type: schema.TypeFloat,
															},
															"ltRtSurroundMixLevel": {
																Type: schema.TypeFloat,
															},
															"metadataControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"passthroughControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"phaseControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"stereoDownmix": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"surroundExMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"surroundMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"mp2Settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"codingMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sampleRate": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"passThroughSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"wavSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitDepth": {
																Type: schema.TypeFloat,
															},
															"codingMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sampleRate": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
											},
										},
									},
									"languageCode": {
										Type: schema.TypeString,
									},
									"languageCodeControl": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"name": {
										Type: schema.TypeString,
									},
									"remixSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"channelMappings": {
													Type: schema.TypeList,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"inputChannelLevels": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"gain": {
																			Type: schema.TypeFloat,
																		},
																		"inputChannel": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"outputChannel": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"channelsIn": {
													Type: schema.TypeFloat,
												},
												"channelsOut": {
													Type: schema.TypeFloat,
												},
											},
										},
									},
									"streamName": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"availBlanking": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availBlankingImage": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"passwordParam": {
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
						"availConfiguration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"scte35SpliceInsert": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adAvailOffset": {
																Type: schema.TypeFloat,
															},
															"noRegionalBlackoutFlag": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"webDeliveryAllowedFlag": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"scte35TimeSignalApos": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adAvailOffset": {
																Type: schema.TypeFloat,
															},
															"noRegionalBlackoutFlag": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"webDeliveryAllowedFlag": {
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
						"blackoutSlate": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blackoutSlateImage": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"passwordParam": {
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
									"networkEndBlackout": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"networkEndBlackoutImage": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"passwordParam": {
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
									"networkId": {
										Type: schema.TypeString,
									},
									"state": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"captionDescriptions": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"captionSelectorName": {
										Type: schema.TypeString,
									},
									"destinationSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aribDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"burnInDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"backgroundColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"backgroundOpacity": {
																Type: schema.TypeFloat,
															},
															"font": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"passwordParam": {
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
															"fontColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"fontOpacity": {
																Type: schema.TypeFloat,
															},
															"fontResolution": {
																Type: schema.TypeFloat,
															},
															"fontSize": {
																Type: schema.TypeString,
															},
															"outlineColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"outlineSize": {
																Type: schema.TypeFloat,
															},
															"shadowColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"shadowOpacity": {
																Type: schema.TypeFloat,
															},
															"shadowXOffset": {
																Type: schema.TypeFloat,
															},
															"shadowYOffset": {
																Type: schema.TypeFloat,
															},
															"teletextGridControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"xPosition": {
																Type: schema.TypeFloat,
															},
															"yPosition": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"dvbSubDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"backgroundColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"backgroundOpacity": {
																Type: schema.TypeFloat,
															},
															"font": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"passwordParam": {
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
															"fontColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"fontOpacity": {
																Type: schema.TypeFloat,
															},
															"fontResolution": {
																Type: schema.TypeFloat,
															},
															"fontSize": {
																Type: schema.TypeString,
															},
															"outlineColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"outlineSize": {
																Type: schema.TypeFloat,
															},
															"shadowColor": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"shadowOpacity": {
																Type: schema.TypeFloat,
															},
															"shadowXOffset": {
																Type: schema.TypeFloat,
															},
															"shadowYOffset": {
																Type: schema.TypeFloat,
															},
															"teletextGridControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"xPosition": {
																Type: schema.TypeFloat,
															},
															"yPosition": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"ebuTtDDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"copyrightHolder": {
																Type: schema.TypeString,
															},
															"fillLineGap": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"fontFamily": {
																Type: schema.TypeString,
															},
															"styleControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"embeddedDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"embeddedPlusScte20DestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"rtmpCaptionInfoDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"scte20PlusEmbeddedDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"scte27DestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"smpteTtDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"teletextDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"ttmlDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"styleControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"webvttDestinationSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"styleControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
											},
										},
									},
									"languageCode": {
										Type: schema.TypeString,
									},
									"languageDescription": {
										Type: schema.TypeString,
									},
									"name": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"featureActivations": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inputPrepareScheduleActions": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"globalConfiguration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initialAudioGain": {
										Type: schema.TypeFloat,
									},
									"inputEndAction": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"inputLossBehavior": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"blackFrameMsec": {
													Type: schema.TypeFloat,
												},
												"inputLossImageColor": {
													Type: schema.TypeString,
												},
												"inputLossImageSlate": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"passwordParam": {
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
												"inputLossImageType": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"repeatFrameMsec": {
													Type: schema.TypeFloat,
												},
											},
										},
									},
									"outputLockingMode": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"outputTimingSource": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"supportLowFramerateInputs": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"motionGraphicsConfiguration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"motionGraphicsInsertion": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"motionGraphicsSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"htmlMotionGraphicsSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
											},
										},
									},
								},
							},
						},
						"nielsenConfiguration": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"distributorId": {
										Type: schema.TypeString,
									},
									"nielsenPcmToId3Tagging": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
								},
							},
						},
						"outputGroups": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString,
									},
									"outputGroupSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"archiveGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"archiveCdnSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"archiveS3Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"cannedAcl": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"logUploads": {
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
																		"destinationRefId": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"rolloverInterval": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"frameCaptureGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destinationRefId": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"frameCaptureCdnSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frameCaptureS3Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"cannedAcl": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"logUploads": {
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
												"hlsGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adMarkers": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"baseUrlContent": {
																Type: schema.TypeString,
															},
															"baseUrlContent1": {
																Type: schema.TypeString,
															},
															"baseUrlManifest": {
																Type: schema.TypeString,
															},
															"baseUrlManifest1": {
																Type: schema.TypeString,
															},
															"captionLanguageMappings": {
																Type: schema.TypeList,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"captionChannel": {
																			Type: schema.TypeFloat,
																		},
																		"languageCode": {
																			Type: schema.TypeString,
																		},
																		"languageDescription": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"captionLanguageSetting": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"clientCache": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"codecSpecification": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"constantIv": {
																Type: schema.TypeString,
															},
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destinationRefId": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"directoryStructure": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"discontinuityTags": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"encryptionType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"hlsCdnSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hlsAkamaiSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connectionRetryInterval": {
																						Type: schema.TypeFloat,
																					},
																					"filecacheDuration": {
																						Type: schema.TypeFloat,
																					},
																					"httpTransferMode": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"numRetries": {
																						Type: schema.TypeFloat,
																					},
																					"restartDelay": {
																						Type: schema.TypeFloat,
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
																		"hlsBasicPutSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connectionRetryInterval": {
																						Type: schema.TypeFloat,
																					},
																					"filecacheDuration": {
																						Type: schema.TypeFloat,
																					},
																					"numRetries": {
																						Type: schema.TypeFloat,
																					},
																					"restartDelay": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"hlsMediaStoreSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connectionRetryInterval": {
																						Type: schema.TypeFloat,
																					},
																					"filecacheDuration": {
																						Type: schema.TypeFloat,
																					},
																					"mediaStoreStorageClass": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"numRetries": {
																						Type: schema.TypeFloat,
																					},
																					"restartDelay": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"hlsS3Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"cannedAcl": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"logUploads": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																				},
																			},
																		},
																		"hlsWebdavSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connectionRetryInterval": {
																						Type: schema.TypeFloat,
																					},
																					"filecacheDuration": {
																						Type: schema.TypeFloat,
																					},
																					"httpTransferMode": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"numRetries": {
																						Type: schema.TypeFloat,
																					},
																					"restartDelay": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"hlsId3SegmentTagging": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"iFrameOnlyPlaylists": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"incompleteSegmentBehavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"indexNSegments": {
																Type: schema.TypeFloat,
															},
															"inputLossAction": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"ivInManifest": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"ivSource": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"keepSegments": {
																Type: schema.TypeFloat,
															},
															"keyFormat": {
																Type: schema.TypeString,
															},
															"keyFormatVersions": {
																Type: schema.TypeString,
															},
															"keyProviderSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"staticKeySettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"keyProviderServer": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"passwordParam": {
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
																					"staticKeyValue": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"manifestCompression": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"manifestDurationFormat": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"minSegmentLength": {
																Type: schema.TypeFloat,
															},
															"mode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"outputSelection": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"programDateTime": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"programDateTimeClock": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"programDateTimePeriod": {
																Type: schema.TypeFloat,
															},
															"redundantManifest": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"segmentLength": {
																Type: schema.TypeFloat,
															},
															"segmentationMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"segmentsPerSubdirectory": {
																Type: schema.TypeFloat,
															},
															"streamInfResolution": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timedMetadataId3Frame": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timedMetadataId3Period": {
																Type: schema.TypeFloat,
															},
															"timestampDeltaMilliseconds": {
																Type: schema.TypeFloat,
															},
															"tsFileMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"mediaPackageGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destinationRefId": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
												"msSmoothGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"acquisitionPointId": {
																Type: schema.TypeString,
															},
															"audioOnlyTimecodeControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"certificateMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"connectionRetryInterval": {
																Type: schema.TypeFloat,
															},
															"destination": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destinationRefId": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"eventId": {
																Type: schema.TypeString,
															},
															"eventIdMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"eventStopBehavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"filecacheDuration": {
																Type: schema.TypeFloat,
															},
															"fragmentLength": {
																Type: schema.TypeFloat,
															},
															"inputLossAction": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"numRetries": {
																Type: schema.TypeFloat,
															},
															"restartDelay": {
																Type: schema.TypeFloat,
															},
															"segmentationMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sendDelayMs": {
																Type: schema.TypeFloat,
															},
															"sparseTrackType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"streamManifestBehavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timestampOffset": {
																Type: schema.TypeString,
															},
															"timestampOffsetMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"multiplexGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
												},
												"rtmpGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adMarkers": {
																Type:  schema.TypeList,
																Elems: &schema.Schema{Type: schema.TypeString},
															},
															"authenticationScheme": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"cacheFullBehavior": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"cacheLength": {
																Type: schema.TypeFloat,
															},
															"captionData": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"inputLossAction": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"restartDelay": {
																Type: schema.TypeFloat,
															},
														},
													},
												},
												"udpGroupSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"inputLossAction": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timedMetadataId3Frame": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timedMetadataId3Period": {
																Type: schema.TypeFloat,
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
												"audioDescriptionNames": {
													Type:  schema.TypeList,
													Elems: &schema.Schema{Type: schema.TypeString},
												},
												"captionDescriptionNames": {
													Type:  schema.TypeList,
													Elems: &schema.Schema{Type: schema.TypeString},
												},
												"outputName": {
													Type: schema.TypeString,
												},
												"outputSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"archiveOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"containerSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"m2tsSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"absentInputAudioBehavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"arib": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"aribCaptionsPid": {
																									Type: schema.TypeString,
																								},
																								"aribCaptionsPidControl": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audioBufferModel": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audioFramesPerPes": {
																									Type: schema.TypeFloat,
																								},
																								"audioPids": {
																									Type: schema.TypeString,
																								},
																								"audioStreamType": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"bitrate": {
																									Type: schema.TypeFloat,
																								},
																								"bufferModel": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ccDescriptor": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"dvbNitSettings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"networkId": {
																												Type: schema.TypeFloat,
																											},
																											"networkName": {
																												Type: schema.TypeString,
																											},
																											"repInterval": {
																												Type: schema.TypeFloat,
																											},
																										},
																									},
																								},
																								"dvbSdtSettings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"outputSdt": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"repInterval": {
																												Type: schema.TypeFloat,
																											},
																											"serviceName": {
																												Type: schema.TypeString,
																											},
																											"serviceProviderName": {
																												Type: schema.TypeString,
																											},
																										},
																									},
																								},
																								"dvbSubPids": {
																									Type: schema.TypeString,
																								},
																								"dvbTdtSettings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"repInterval": {
																												Type: schema.TypeFloat,
																											},
																										},
																									},
																								},
																								"dvbTeletextPid": {
																									Type: schema.TypeString,
																								},
																								"ebif": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebpAudioInterval": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebpLookaheadMs": {
																									Type: schema.TypeFloat,
																								},
																								"ebpPlacement": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ecmPid": {
																									Type: schema.TypeString,
																								},
																								"esRateInPes": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"etvPlatformPid": {
																									Type: schema.TypeString,
																								},
																								"etvSignalPid": {
																									Type: schema.TypeString,
																								},
																								"fragmentTime": {
																									Type: schema.TypeFloat,
																								},
																								"klv": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"klvDataPids": {
																									Type: schema.TypeString,
																								},
																								"nielsenId3Behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"nullPacketBitrate": {
																									Type: schema.TypeFloat,
																								},
																								"patInterval": {
																									Type: schema.TypeFloat,
																								},
																								"pcrControl": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"pcrPeriod": {
																									Type: schema.TypeFloat,
																								},
																								"pcrPid": {
																									Type: schema.TypeString,
																								},
																								"pmtInterval": {
																									Type: schema.TypeFloat,
																								},
																								"pmtPid": {
																									Type: schema.TypeString,
																								},
																								"programNum": {
																									Type: schema.TypeFloat,
																								},
																								"rateMode": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte27Pids": {
																									Type: schema.TypeString,
																								},
																								"scte35Control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte35Pid": {
																									Type: schema.TypeString,
																								},
																								"segmentationMarkers": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentationStyle": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentationTime": {
																									Type: schema.TypeFloat,
																								},
																								"timedMetadataBehavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"timedMetadataPid": {
																									Type: schema.TypeString,
																								},
																								"transportStreamId": {
																									Type: schema.TypeFloat,
																								},
																								"videoPid": {
																									Type: schema.TypeString,
																								},
																							},
																						},
																					},
																					"rawSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																					},
																				},
																			},
																		},
																		"extension": {
																			Type: schema.TypeString,
																		},
																		"nameModifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"frameCaptureOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nameModifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"hlsOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"h265PackagingType": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"hlsSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"audioOnlyHlsSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"audioGroupId": {
																									Type: schema.TypeString,
																								},
																								"audioOnlyImage": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"passwordParam": {
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
																								"audioTrackType": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentType": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																							},
																						},
																					},
																					"fmp4HlsSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"audioRenditionSets": {
																									Type: schema.TypeString,
																								},
																								"nielsenId3Behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"timedMetadataBehavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																							},
																						},
																					},
																					"frameCaptureHlsSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																					},
																					"standardHlsSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"audioRenditionSets": {
																									Type: schema.TypeString,
																								},
																								"m3u8Settings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"audioFramesPerPes": {
																												Type: schema.TypeFloat,
																											},
																											"audioPids": {
																												Type: schema.TypeString,
																											},
																											"ecmPid": {
																												Type: schema.TypeString,
																											},
																											"nielsenId3Behavior": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"patInterval": {
																												Type: schema.TypeFloat,
																											},
																											"pcrControl": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"pcrPeriod": {
																												Type: schema.TypeFloat,
																											},
																											"pcrPid": {
																												Type: schema.TypeString,
																											},
																											"pmtInterval": {
																												Type: schema.TypeFloat,
																											},
																											"pmtPid": {
																												Type: schema.TypeString,
																											},
																											"programNum": {
																												Type: schema.TypeFloat,
																											},
																											"scte35Behavior": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"scte35Pid": {
																												Type: schema.TypeString,
																											},
																											"timedMetadataBehavior": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"timedMetadataPid": {
																												Type: schema.TypeString,
																											},
																											"transportStreamId": {
																												Type: schema.TypeFloat,
																											},
																											"videoPid": {
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
																		"nameModifier": {
																			Type: schema.TypeString,
																		},
																		"segmentModifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"mediaPackageOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
															},
															"msSmoothOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"h265PackagingType": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"nameModifier": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"multiplexOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"destination": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"destinationRefId": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"rtmpOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"certificateMode": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"connectionRetryInterval": {
																			Type: schema.TypeFloat,
																		},
																		"destination": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"destinationRefId": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"numRetries": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"udpOutputSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"bufferMsec": {
																			Type: schema.TypeFloat,
																		},
																		"containerSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"m2tsSettings": {
																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Elems: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"absentInputAudioBehavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"arib": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"aribCaptionsPid": {
																									Type: schema.TypeString,
																								},
																								"aribCaptionsPidControl": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audioBufferModel": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"audioFramesPerPes": {
																									Type: schema.TypeFloat,
																								},
																								"audioPids": {
																									Type: schema.TypeString,
																								},
																								"audioStreamType": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"bitrate": {
																									Type: schema.TypeFloat,
																								},
																								"bufferModel": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ccDescriptor": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"dvbNitSettings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"networkId": {
																												Type: schema.TypeFloat,
																											},
																											"networkName": {
																												Type: schema.TypeString,
																											},
																											"repInterval": {
																												Type: schema.TypeFloat,
																											},
																										},
																									},
																								},
																								"dvbSdtSettings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"outputSdt": {
																												Type: schema.TypeString, // TODO: enum, add validation
																											},
																											"repInterval": {
																												Type: schema.TypeFloat,
																											},
																											"serviceName": {
																												Type: schema.TypeString,
																											},
																											"serviceProviderName": {
																												Type: schema.TypeString,
																											},
																										},
																									},
																								},
																								"dvbSubPids": {
																									Type: schema.TypeString,
																								},
																								"dvbTdtSettings": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Elems: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"repInterval": {
																												Type: schema.TypeFloat,
																											},
																										},
																									},
																								},
																								"dvbTeletextPid": {
																									Type: schema.TypeString,
																								},
																								"ebif": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebpAudioInterval": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ebpLookaheadMs": {
																									Type: schema.TypeFloat,
																								},
																								"ebpPlacement": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"ecmPid": {
																									Type: schema.TypeString,
																								},
																								"esRateInPes": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"etvPlatformPid": {
																									Type: schema.TypeString,
																								},
																								"etvSignalPid": {
																									Type: schema.TypeString,
																								},
																								"fragmentTime": {
																									Type: schema.TypeFloat,
																								},
																								"klv": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"klvDataPids": {
																									Type: schema.TypeString,
																								},
																								"nielsenId3Behavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"nullPacketBitrate": {
																									Type: schema.TypeFloat,
																								},
																								"patInterval": {
																									Type: schema.TypeFloat,
																								},
																								"pcrControl": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"pcrPeriod": {
																									Type: schema.TypeFloat,
																								},
																								"pcrPid": {
																									Type: schema.TypeString,
																								},
																								"pmtInterval": {
																									Type: schema.TypeFloat,
																								},
																								"pmtPid": {
																									Type: schema.TypeString,
																								},
																								"programNum": {
																									Type: schema.TypeFloat,
																								},
																								"rateMode": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte27Pids": {
																									Type: schema.TypeString,
																								},
																								"scte35Control": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"scte35Pid": {
																									Type: schema.TypeString,
																								},
																								"segmentationMarkers": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentationStyle": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"segmentationTime": {
																									Type: schema.TypeFloat,
																								},
																								"timedMetadataBehavior": {
																									Type: schema.TypeString, // TODO: enum, add validation
																								},
																								"timedMetadataPid": {
																									Type: schema.TypeString,
																								},
																								"transportStreamId": {
																									Type: schema.TypeFloat,
																								},
																								"videoPid": {
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
																					"destinationRefId": {
																						Type: schema.TypeString,
																					},
																				},
																			},
																		},
																		"fecOutputSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"columnDepth": {
																						Type: schema.TypeFloat,
																					},
																					"includeFec": {
																						Type: schema.TypeString, // TODO: enum, add validation
																					},
																					"rowLength": {
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
												"videoDescriptionName": {
													Type: schema.TypeString,
												},
											},
										},
									},
								},
							},
						},
						"timecodeConfig": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"syncThreshold": {
										Type: schema.TypeFloat,
									},
								},
							},
						},
						"videoDescriptions": {
							Type: schema.TypeList,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"codecSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"frameCaptureSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"captureInterval": {
																Type: schema.TypeFloat,
															},
															"captureIntervalUnits": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"h264Settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptiveQuantization": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"afdSignaling": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"bufFillPct": {
																Type: schema.TypeFloat,
															},
															"bufSize": {
																Type: schema.TypeFloat,
															},
															"colorMetadata": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"colorSpaceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"colorSpacePassthroughSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"rec601Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"rec709Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																	},
																},
															},
															"entropyEncoding": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"filterSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"temporalFilterSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"postFilterSharpening": {
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
															"fixedAfd": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"flickerAq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"forceFieldPictures": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerateControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerateDenominator": {
																Type: schema.TypeFloat,
															},
															"framerateNumerator": {
																Type: schema.TypeFloat,
															},
															"gopBReference": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"gopClosedCadence": {
																Type: schema.TypeFloat,
															},
															"gopNumBFrames": {
																Type: schema.TypeFloat,
															},
															"gopSize": {
																Type: schema.TypeFloat,
															},
															"gopSizeUnits": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"level": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lookAheadRateControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"maxBitrate": {
																Type: schema.TypeFloat,
															},
															"minIInterval": {
																Type: schema.TypeFloat,
															},
															"numRefFrames": {
																Type: schema.TypeFloat,
															},
															"parControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"parDenominator": {
																Type: schema.TypeFloat,
															},
															"parNumerator": {
																Type: schema.TypeFloat,
															},
															"profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"qualityLevel": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"qvbrQualityLevel": {
																Type: schema.TypeFloat,
															},
															"rateControlMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scanType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sceneChangeDetect": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"slices": {
																Type: schema.TypeFloat,
															},
															"softness": {
																Type: schema.TypeFloat,
															},
															"spatialAq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"subgopLength": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"syntax": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"temporalAq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timecodeInsertion": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"h265Settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptiveQuantization": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"afdSignaling": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"alternativeTransferFunction": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"bitrate": {
																Type: schema.TypeFloat,
															},
															"bufSize": {
																Type: schema.TypeFloat,
															},
															"colorMetadata": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"colorSpaceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"colorSpacePassthroughSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"hdr10Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"maxCll": {
																						Type: schema.TypeFloat,
																					},
																					"maxFall": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"rec601Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																		"rec709Settings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																		},
																	},
																},
															},
															"filterSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"temporalFilterSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"postFilterSharpening": {
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
															"fixedAfd": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"flickerAq": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerateDenominator": {
																Type: schema.TypeFloat,
															},
															"framerateNumerator": {
																Type: schema.TypeFloat,
															},
															"gopClosedCadence": {
																Type: schema.TypeFloat,
															},
															"gopSize": {
																Type: schema.TypeFloat,
															},
															"gopSizeUnits": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"level": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"lookAheadRateControl": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"maxBitrate": {
																Type: schema.TypeFloat,
															},
															"minIInterval": {
																Type: schema.TypeFloat,
															},
															"parDenominator": {
																Type: schema.TypeFloat,
															},
															"parNumerator": {
																Type: schema.TypeFloat,
															},
															"profile": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"qvbrQualityLevel": {
																Type: schema.TypeFloat,
															},
															"rateControlMode": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scanType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"sceneChangeDetect": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"slices": {
																Type: schema.TypeFloat,
															},
															"tier": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timecodeInsertion": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"mpeg2Settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptiveQuantization": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"afdSignaling": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"colorMetadata": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"colorSpace": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"displayAspectRatio": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"filterSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"temporalFilterSettings": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"postFilterSharpening": {
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
															"fixedAfd": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"framerateDenominator": {
																Type: schema.TypeFloat,
															},
															"framerateNumerator": {
																Type: schema.TypeFloat,
															},
															"gopClosedCadence": {
																Type: schema.TypeFloat,
															},
															"gopNumBFrames": {
																Type: schema.TypeFloat,
															},
															"gopSize": {
																Type: schema.TypeFloat,
															},
															"gopSizeUnits": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"scanType": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"subgopLength": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
															"timecodeInsertion": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
											},
										},
									},
									"height": {
										Type: schema.TypeFloat,
									},
									"name": {
										Type: schema.TypeString,
									},
									"respondToAfd": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"scalingBehavior": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"sharpness": {
										Type: schema.TypeFloat,
									},
									"width": {
										Type: schema.TypeFloat,
									},
								},
							},
						},
					},
				},
			},
			"inputAttachments": {
				Type: schema.TypeList,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automaticInputFailoverSettings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"errorClearTimeMsec": {
										Type: schema.TypeFloat,
									},
									"failoverConditions": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"failoverConditionSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"audioSilenceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audioSelectorName": {
																			Type: schema.TypeString,
																		},
																		"audioSilenceThresholdMsec": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"inputLossSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"inputLossThresholdMsec": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"videoBlackSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"blackDetectThreshold": {
																			Type: schema.TypeFloat,
																		},
																		"videoBlackThresholdMsec": {
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
									"inputPreference": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"secondaryInputId": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"inputAttachmentName": {
							Type: schema.TypeString,
						},
						"inputId": {
							Type: schema.TypeString,
						},
						"inputSettings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elems: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audioSelectors": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString,
												},
												"selectorSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"audioHlsRenditionSelection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"groupId": {
																			Type: schema.TypeString,
																		},
																		"name": {
																			Type: schema.TypeString,
																		},
																	},
																},
															},
															"audioLanguageSelection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"languageCode": {
																			Type: schema.TypeString,
																		},
																		"languageSelectionPolicy": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																	},
																},
															},
															"audioPidSelection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"audioTrackSelection": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"tracks": {
																			Type: schema.TypeList,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"track": {
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
											},
										},
									},
									"captionSelectors": {
										Type: schema.TypeList,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"languageCode": {
													Type: schema.TypeString,
												},
												"name": {
													Type: schema.TypeString,
												},
												"selectorSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ancillarySourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"sourceAncillaryChannelNumber": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"aribSourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
															},
															"dvbSubSourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ocrLanguage": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"embeddedSourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"convert608To708": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"scte20Detection": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"source608ChannelNumber": {
																			Type: schema.TypeFloat,
																		},
																		"source608TrackNumber": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"scte20SourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"convert608To708": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"source608ChannelNumber": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"scte27SourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ocrLanguage": {
																			Type: schema.TypeString, // TODO: enum, add validation
																		},
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"teletextSourceSettings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"outputRectangle": {
																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Elems: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"height": {
																						Type: schema.TypeFloat,
																					},
																					"leftOffset": {
																						Type: schema.TypeFloat,
																					},
																					"topOffset": {
																						Type: schema.TypeFloat,
																					},
																					"width": {
																						Type: schema.TypeFloat,
																					},
																				},
																			},
																		},
																		"pageNumber": {
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
									"deblockFilter": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"denoiseFilter": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"filterStrength": {
										Type: schema.TypeFloat,
									},
									"inputFilter": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"networkInputSettings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hlsInputSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bandwidth": {
																Type: schema.TypeFloat,
															},
															"bufferSegments": {
																Type: schema.TypeFloat,
															},
															"retries": {
																Type: schema.TypeFloat,
															},
															"retryInterval": {
																Type: schema.TypeFloat,
															},
															"scte35Source": {
																Type: schema.TypeString, // TODO: enum, add validation
															},
														},
													},
												},
												"serverValidation": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
											},
										},
									},
									"smpte2038DataPreference": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"sourceEndBehavior": {
										Type: schema.TypeString, // TODO: enum, add validation
									},
									"videoSelector": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elems: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"colorSpace": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"colorSpaceSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hdr10Settings": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"maxCll": {
																			Type: schema.TypeFloat,
																		},
																		"maxFall": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
														},
													},
												},
												"colorSpaceUsage": {
													Type: schema.TypeString, // TODO: enum, add validation
												},
												"selectorSettings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elems: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"videoSelectorPid": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"pid": {
																			Type: schema.TypeFloat,
																		},
																	},
																},
															},
															"videoSelectorProgramId": {
																Type:     schema.TypeList,
																MaxItems: 1,
																Elems: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"programId": {
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
								},
							},
						},
					},
				},
			},
			"inputSpecification": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"codec": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
						"maximumBitrate": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
						"resolution": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
					},
				},
			},
			"logLevel": {
				Type: schema.TypeString, // TODO: enum, add validation
			},
			"maintenance": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maintenanceDay": {
							Type: schema.TypeString, // TODO: enum, add validation
						},
						"maintenanceStartTime": {
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
			"roleArn": {
				Type: schema.TypeString,
			},
			"vpc": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elems: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"publicAddressAllocationIds": {
							Type:  schema.TypeList,
							Elems: &schema.Schema{Type: schema.TypeString},
						},
						"securityGroupIds": {
							Type:  schema.TypeList,
							Elems: &schema.Schema{Type: schema.TypeString},
						},
						"subnetIds": {
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
