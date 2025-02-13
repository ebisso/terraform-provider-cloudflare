// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessApplicationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate to this application using their WARP session.  When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.",
				Optional:    true,
			},
			"app_launcher_logo_url": schema.StringAttribute{
				Description: "The image URL of the logo shown in the App Launcher header.",
				Optional:    true,
			},
			"bg_color": schema.StringAttribute{
				Description: "The background color of the App Launcher page.",
				Optional:    true,
			},
			"custom_deny_message": schema.StringAttribute{
				Description: "The custom error message shown to a user when they are denied access to the application.",
				Optional:    true,
			},
			"custom_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing identity-based rules.",
				Optional:    true,
			},
			"custom_non_identity_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing non-identity rules.",
				Optional:    true,
			},
			"domain": schema.StringAttribute{
				Description: "The primary hostname and path secured by Access. This domain will be displayed if the app is visible in the App Launcher.",
				Optional:    true,
			},
			"header_bg_color": schema.StringAttribute{
				Description: "The background color of the App Launcher header.",
				Optional:    true,
			},
			"logo_url": schema.StringAttribute{
				Description: "The image URL for the logo shown in the App Launcher dashboard.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the application.",
				Optional:    true,
			},
			"options_preflight_bypass": schema.BoolAttribute{
				Description: "Allows options preflight requests to bypass Access authentication and go directly to the origin. Cannot turn on if cors_headers is set.",
				Optional:    true,
			},
			"same_site_cookie_attribute": schema.StringAttribute{
				Description: "Sets the SameSite cookie setting, which provides increased security against CSRF attacks.",
				Optional:    true,
			},
			"service_auth_401_redirect": schema.BoolAttribute{
				Description: "Returns a 401 status code when the request is blocked by a Service Auth policy.",
				Optional:    true,
			},
			"skip_interstitial": schema.BoolAttribute{
				Description: "Enables automatic authentication through cloudflared.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The application type.",
				Optional:    true,
			},
			"allowed_idps": schema.ListAttribute{
				Description: "The identity providers your users can select when connecting to this application. Defaults to all IdPs configured in your account.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"custom_pages": schema.ListAttribute{
				Description: "The custom pages that will be displayed when applicable for this application",
				Optional:    true,
				ElementType: types.StringType,
			},
			"self_hosted_domains": schema.ListAttribute{
				Description: "List of public domains that Access will secure. This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.\n",
				Optional:    true,
				ElementType: types.StringType,
			},
			"tags": schema.ListAttribute{
				Description: "The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"app_launcher_visible": schema.BoolAttribute{
				Description: "Displays the application in the App Launcher.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login. You must specify only one identity provider in allowed_idps.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enable_binding_cookie": schema.BoolAttribute{
				Description: "Enables the binding cookie, which increases security against compromised authorization tokens and CSRF attacks.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"http_only_cookie_attribute": schema.BoolAttribute{
				Description: "Enables the HttpOnly cookie attribute, which increases security against XSS attacks.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"path_cookie_attribute": schema.BoolAttribute{
				Description: "Enables cookie paths to scope an application's JWT to the application path. If disabled, the JWT will scope to the hostname by default",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for this application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("24h"),
			},
			"skip_app_launcher_login_page": schema.BoolAttribute{
				Description: "Determines when to skip the App Launcher landing page.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"cors_headers": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationCORSHeadersModel](ctx),
				Attributes: map[string]schema.Attribute{
					"allow_all_headers": schema.BoolAttribute{
						Description: "Allows all HTTP request headers.",
						Optional:    true,
					},
					"allow_all_methods": schema.BoolAttribute{
						Description: "Allows all HTTP request methods.",
						Optional:    true,
					},
					"allow_all_origins": schema.BoolAttribute{
						Description: "Allows all origins.",
						Optional:    true,
					},
					"allow_credentials": schema.BoolAttribute{
						Description: "When set to `true`, includes credentials (cookies, authorization headers, or TLS client certificates) with requests.",
						Optional:    true,
					},
					"allowed_headers": schema.ListAttribute{
						Description: "Allowed HTTP request headers.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"allowed_methods": schema.ListAttribute{
						Description: "Allowed HTTP request methods.",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"GET",
									"POST",
									"HEAD",
									"PUT",
									"DELETE",
									"CONNECT",
									"OPTIONS",
									"TRACE",
									"PATCH",
								),
							),
						},
						ElementType: types.StringType,
					},
					"allowed_origins": schema.ListAttribute{
						Description: "Allowed origins.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"max_age": schema.Float64Attribute{
						Description: "The maximum number of seconds the results of a preflight request can be cached.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(-1, 86400),
						},
					},
				},
			},
			"destinations": schema.ListNestedAttribute{
				Description: "List of destinations secured by Access. This supersedes `self_hosted_domains` to allow for more flexibility in defining different types of domains. If `destinations` are provided, then `self_hosted_domains` will be ignored.\n",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationDestinationsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("public", "private"),
							},
						},
						"uri": schema.StringAttribute{
							Description: "The URI of the destination. Public destinations' URIs can include a domain and path with [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).\n",
							Optional:    true,
						},
						"cidr": schema.StringAttribute{
							Description: "The CIDR range of the destination. Single IPs will be computed as /32.",
							Optional:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname of the destination. Matches a valid SNI served by an HTTPS origin.",
							Optional:    true,
						},
						"l4_protocol": schema.StringAttribute{
							Description: "The L4 protocol of the destination. When omitted, both UDP and TCP traffic will match.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("tcp", "udp"),
							},
						},
						"port_range": schema.StringAttribute{
							Description: "The port range of the destination. Can be a single port or a range of ports. When omitted, all ports will match.\n",
							Optional:    true,
						},
						"vnet_id": schema.StringAttribute{
							Description: "The VNET ID to match the destination. When omitted, all VNETs will match.",
							Optional:    true,
						},
					},
				},
			},
			"footer_links": schema.ListNestedAttribute{
				Description: "The links in the App Launcher footer.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationFooterLinksModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "The hypertext in the footer link.",
							Required:    true,
						},
						"url": schema.StringAttribute{
							Description: "the hyperlink in the footer link.",
							Required:    true,
						},
					},
				},
			},
			"landing_page_design": schema.SingleNestedAttribute{
				Description: "The design of the App Launcher landing page shown to users when they log in.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationLandingPageDesignModel](ctx),
				Attributes: map[string]schema.Attribute{
					"button_color": schema.StringAttribute{
						Description: "The background color of the log in button on the landing page.",
						Optional:    true,
					},
					"button_text_color": schema.StringAttribute{
						Description: "The color of the text in the log in button on the landing page.",
						Optional:    true,
					},
					"image_url": schema.StringAttribute{
						Description: "The URL of the image shown on the landing page.",
						Optional:    true,
					},
					"message": schema.StringAttribute{
						Description: "The message shown on the landing page.",
						Optional:    true,
					},
					"title": schema.StringAttribute{
						Description: "The title shown on the landing page.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("Welcome!"),
					},
				},
			},
			"policies": schema.ListNestedAttribute{
				Description: "The policies that Access applies to the application, in ascending order of precedence. Items can reference existing policies or create new policies exclusive to the application.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the policy",
							Optional:    true,
						},
						"precedence": schema.Int64Attribute{
							Description: "The order of execution for this policy. Must be unique for each policy within an app.",
							Optional:    true,
						},
						"decision": schema.StringAttribute{
							Description: "The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"allow",
									"deny",
									"non_identity",
									"bypass",
								),
							},
						},
						"include": schema.ListNestedAttribute{
							Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesIncludeModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGroupModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Required:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAuthContextModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Required:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAuthMethodModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Required:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeAzureADModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Optional:   true,
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeCommonNameModel](ctx),
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Required:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGeoModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Required:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeDevicePostureModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Required:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEmailDomainModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Required:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEmailListModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Required:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeEmailModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Required:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Required:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Required:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Required:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Optional:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeGSuiteModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Required:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeLoginMethodModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Required:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeIPListModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Required:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeIPModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Required:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeOktaModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Required:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeSAMLModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Required:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Required:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesIncludeServiceTokenModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Required:    true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Access policy.",
							Optional:    true,
						},
						"connection_rules": schema.SingleNestedAttribute{
							Description: "The rules that define how users may connect to the targets secured by your application.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesConnectionRulesModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ssh": schema.SingleNestedAttribute{
									Description: "The SSH-specific rules that define how users may connect to the targets secured by your application.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesConnectionRulesSSHModel](ctx),
									Attributes: map[string]schema.Attribute{
										"usernames": schema.ListAttribute{
											Description: "Contains the Unix usernames that may be used when connecting over SSH.",
											Required:    true,
											ElementType: types.StringType,
										},
										"allow_email_alias": schema.BoolAttribute{
											Description: "Enables using Identity Provider email alias as SSH username.",
											Optional:    true,
										},
									},
								},
							},
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesExcludeModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGroupModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Required:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAuthContextModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Required:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAuthMethodModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Required:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeAzureADModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Optional:   true,
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeCommonNameModel](ctx),
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Required:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGeoModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Required:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeDevicePostureModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Required:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEmailDomainModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Required:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEmailListModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Required:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeEmailModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Required:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Required:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Required:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Required:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Optional:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeGSuiteModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Required:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeLoginMethodModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Required:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeIPListModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Required:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeIPModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Required:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeOktaModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Required:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeSAMLModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Required:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Required:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesExcludeServiceTokenModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Required:    true,
											},
										},
									},
								},
							},
						},
						"require": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesRequireModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGroupModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Required:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAuthContextModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Required:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAuthMethodModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Required:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireAzureADModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Optional:   true,
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireCommonNameModel](ctx),
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Required:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGeoModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Required:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireDevicePostureModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Required:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEmailDomainModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Required:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEmailListModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Required:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireEmailModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Required:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Required:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Required:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Required:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Optional:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireGSuiteModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Required:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireLoginMethodModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Required:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireIPListModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Required:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireIPModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Required:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireOktaModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Required:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireSAMLModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Required:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Required:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										Optional:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationPoliciesRequireServiceTokenModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Required:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"saas_app": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppModel](ctx),
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Description: "Optional identifier indicating the authentication protocol used for the saas app. Required for OIDC. Default if unset is \"saml\"",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("saml", "oidc"),
						},
					},
					"consumer_service_url": schema.StringAttribute{
						Description: "The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.",
						Optional:    true,
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"custom_attributes": schema.ListNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSaaSAppCustomAttributesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"friendly_name": schema.StringAttribute{
									Description: "The SAML FriendlyName of the attribute.",
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the attribute.",
									Optional:    true,
								},
								"name_format": schema.StringAttribute{
									Description: "A globally unique name for an identity or service provider.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified",
											"urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
											"urn:oasis:names:tc:SAML:2.0:attrname-format:uri",
										),
									},
								},
								"required": schema.BoolAttribute{
									Description: "If the attribute is required when building a SAML assertion.",
									Optional:    true,
								},
								"source": schema.SingleNestedAttribute{
									Computed:   true,
									Optional:   true,
									CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "The name of the IdP attribute.",
											Optional:    true,
										},
										"name_by_idp": schema.ListNestedAttribute{
											Description: "A mapping from IdP ID to attribute name.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceNameByIdPModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"idp_id": schema.StringAttribute{
														Description: "The UID of the IdP.",
														Optional:    true,
													},
													"source_name": schema.StringAttribute{
														Description: "The name of the IdP provided attribute.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"default_relay_state": schema.StringAttribute{
						Description: "The URL that the user will be redirected to after a successful login for IDP initiated logins.",
						Optional:    true,
					},
					"idp_entity_id": schema.StringAttribute{
						Description: "The unique identifier for your SaaS application.",
						Computed:    true,
						Optional:    true,
					},
					"name_id_format": schema.StringAttribute{
						Description: "The format of the name identifier sent to the SaaS application.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "email"),
						},
					},
					"name_id_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.\n",
						Optional:    true,
					},
					"public_key": schema.StringAttribute{
						Description: "The Access public certificate that will be used to verify your identity.",
						Computed:    true,
						Optional:    true,
					},
					"saml_attribute_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata] (https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.\n",
						Optional:    true,
					},
					"sp_entity_id": schema.StringAttribute{
						Description: "A globally unique name for an identity or service provider.",
						Optional:    true,
					},
					"sso_endpoint": schema.StringAttribute{
						Description: "The endpoint where your SaaS application will send login requests.",
						Computed:    true,
						Optional:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"access_token_lifetime": schema.StringAttribute{
						Description: "The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must be greater than or equal to 1m and less than or equal to 24h.",
						Optional:    true,
					},
					"allow_pkce_without_client_secret": schema.BoolAttribute{
						Description: "If client secret should be required on the token endpoint when authorization_code_with_pkce grant is used.",
						Optional:    true,
					},
					"app_launcher_url": schema.StringAttribute{
						Description: "The URL where this applications tile redirects users",
						Optional:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "The application client id",
						Computed:    true,
						Optional:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "The application client secret, only returned on POST request.",
						Computed:    true,
						Optional:    true,
					},
					"custom_claims": schema.ListNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSaaSAppCustomClaimsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "The name of the claim.",
									Optional:    true,
								},
								"required": schema.BoolAttribute{
									Description: "If the claim is required when building an OIDC token.",
									Optional:    true,
								},
								"scope": schema.StringAttribute{
									Description: "The scope of the claim.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"groups",
											"profile",
											"email",
											"openid",
										),
									},
								},
								"source": schema.SingleNestedAttribute{
									Computed:   true,
									Optional:   true,
									CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "The name of the IdP claim.",
											Optional:    true,
										},
										"name_by_idp": schema.MapAttribute{
											Description: "A mapping from IdP ID to claim name.",
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
						},
					},
					"grant_types": schema.ListAttribute{
						Description: "The OIDC flows supported by this application",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"authorization_code",
									"authorization_code_with_pkce",
									"refresh_tokens",
									"hybrid",
									"implicit",
								),
							),
						},
						ElementType: types.StringType,
					},
					"group_filter_regex": schema.StringAttribute{
						Description: "A regex to filter Cloudflare groups returned in ID token and userinfo endpoint",
						Optional:    true,
					},
					"hybrid_and_implicit_options": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"return_access_token_from_authorization_endpoint": schema.BoolAttribute{
								Description: "If an Access Token should be returned from the OIDC Authorization endpoint",
								Optional:    true,
							},
							"return_id_token_from_authorization_endpoint": schema.BoolAttribute{
								Description: "If an ID Token should be returned from the OIDC Authorization endpoint",
								Optional:    true,
							},
						},
					},
					"redirect_uris": schema.ListAttribute{
						Description: "The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens",
						Optional:    true,
						ElementType: types.StringType,
					},
					"refresh_token_options": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"lifetime": schema.StringAttribute{
								Description: "How long a refresh token will be valid for after creation. Valid units are m,h,d. Must be longer than 1m.",
								Optional:    true,
							},
						},
					},
					"scopes": schema.ListAttribute{
						Description: "Define the user information shared with access, \"offline_access\" scope will be automatically enabled if refresh tokens are enabled",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"openid",
									"groups",
									"email",
									"profile",
								),
							),
						},
						ElementType: types.StringType,
					},
				},
			},
			"scim_config": schema.SingleNestedAttribute{
				Description: "Configuration for provisioning to this application via SCIM. This is currently in closed beta.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationSCIMConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"idp_uid": schema.StringAttribute{
						Description: "The UID of the IdP to use as the source for SCIM resources to provision to this application.",
						Required:    true,
					},
					"remote_uri": schema.StringAttribute{
						Description: "The base URI for the application's SCIM-compatible API.",
						Required:    true,
					},
					"authentication": schema.SingleNestedAttribute{
						Description: "Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationSCIMConfigAuthenticationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"password": schema.StringAttribute{
								Description: "Password used to authenticate with the remote SCIM service.",
								Optional:    true,
							},
							"scheme": schema.StringAttribute{
								Description: "The authentication scheme to use when making SCIM requests to this application.",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"httpbasic",
										"oauthbearertoken",
										"oauth2",
										"access_service_token",
									),
								},
							},
							"user": schema.StringAttribute{
								Description: "User name used to authenticate with the remote SCIM service.",
								Optional:    true,
							},
							"token": schema.StringAttribute{
								Description: "Token used to authenticate with the remote SCIM service.",
								Optional:    true,
							},
							"authorization_url": schema.StringAttribute{
								Description: "URL used to generate the auth code used during token generation.",
								Optional:    true,
							},
							"client_id": schema.StringAttribute{
								Description: "Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.",
								Optional:    true,
							},
							"client_secret": schema.StringAttribute{
								Description: "Secret used to authenticate when generating a token for authenticating with the remove SCIM service.",
								Optional:    true,
							},
							"token_url": schema.StringAttribute{
								Description: "URL used to generate the token used to authenticate with the remote SCIM service.",
								Optional:    true,
							},
							"scopes": schema.ListAttribute{
								Description: "The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"deactivate_on_delete": schema.BoolAttribute{
						Description: "If false, propagates DELETE requests to the target application for SCIM resources. If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.",
						Optional:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether SCIM provisioning is turned on for this application.",
						Optional:    true,
					},
					"mappings": schema.ListNestedAttribute{
						Description: "A list of mappings to apply to SCIM resources before provisioning them in this application. These can transform or filter the resources to be provisioned.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationSCIMConfigMappingsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"schema": schema.StringAttribute{
									Description: "Which SCIM resource type this mapping applies to.",
									Required:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether or not this mapping is enabled.",
									Optional:    true,
								},
								"filter": schema.StringAttribute{
									Description: "A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.",
									Optional:    true,
								},
								"operations": schema.SingleNestedAttribute{
									Description: "Whether or not this mapping applies to creates, updates, or deletes.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel](ctx),
									Attributes: map[string]schema.Attribute{
										"create": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to create (POST) operations.",
											Optional:    true,
										},
										"delete": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to DELETE operations.",
											Optional:    true,
										},
										"update": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to update (PATCH/PUT) operations.",
											Optional:    true,
										},
									},
								},
								"strictness": schema.StringAttribute{
									Description: "The level of adherence to outbound resource schemas when provisioning to this mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown values to the target.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("strict", "passthrough"),
									},
								},
								"transform_jsonata": schema.StringAttribute{
									Description: "A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.",
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"target_criteria": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationTargetCriteriaModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Int64Attribute{
							Description: "The port that the targets use for the chosen communication protocol. A port cannot be assigned to multiple protocols.",
							Required:    true,
						},
						"protocol": schema.StringAttribute{
							Description: "The communication protocol your application secures.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("ssh"),
							},
						},
						"target_attributes": schema.MapAttribute{
							Description: "Contains a map of target attribute keys to target attribute values.",
							Required:    true,
							ElementType: types.ListType{
								ElemType: types.StringType,
							},
						},
					},
				},
			},
			"aud": schema.StringAttribute{
				Description: "Audience tag.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustAccessApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessApplicationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
