// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeForwardingRuleCreate,
		Read:   resourceComputeForwardingRuleRead,
		Update: resourceComputeForwardingRuleUpdate,
		Delete: resourceComputeForwardingRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeForwardingRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"ip_address": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateIpAddress,
				Description: `The IP address that this forwarding rule is serving on behalf of.

Addresses are restricted based on the forwarding rule's load balancing
scheme (EXTERNAL or INTERNAL) and scope (global or regional).

When the load balancing scheme is EXTERNAL, for global forwarding
rules, the address must be a global IP, and for regional forwarding
rules, the address must live in the same region as the forwarding
rule. If this field is empty, an ephemeral IPv4 address from the same
scope (global or regional) will be assigned. A regional forwarding
rule supports IPv4 only. A global forwarding rule supports either IPv4
or IPv6.

When the load balancing scheme is INTERNAL, this can only be an RFC
1918 IP address belonging to the network/subnet configured for the
forwarding rule. By default, if this field is empty, an ephemeral
internal IP address will be automatically allocated from the IP range
of the subnet or network configured for this forwarding rule.

An address must be specified by a literal IP address. ~> **NOTE:** While
the API allows you to specify various resource paths for an address resource
instead, Terraform requires this to specifically be an IP address to
avoid needing to fetching the IP address from resource paths on refresh
or unnecessary diffs.`,
			},
			"ip_protocol": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"TCP", "UDP", "ESP", "AH", "SCTP", "ICMP", ""}, false),
				DiffSuppressFunc: caseDiffSuppress,
				Description: `The IP protocol to which this rule applies.

When the load balancing scheme is INTERNAL, only TCP and UDP are
valid. Possible values: ["TCP", "UDP", "ESP", "AH", "SCTP", "ICMP"]`,
			},
			"all_ports": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `For internal TCP/UDP load balancing (i.e. load balancing scheme is
INTERNAL and protocol is TCP/UDP), set this to true to allow packets
addressed to any ports to be forwarded to the backends configured
with this forwarding rule. Used with backend service. Cannot be set
if port or portRange are set.`,
			},
			"allow_global_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If true, clients can access ILB from all regions.
Otherwise only allows from the local region the ILB is located at.`,
			},
			"backend_service": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `A BackendService to receive the matched traffic. This is used only
for INTERNAL load balancing.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"load_balancing_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"EXTERNAL", "INTERNAL", "INTERNAL_MANAGED", ""}, false),
				Description: `This signifies what the ForwardingRule will be used for and can be
EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is used for Classic
Cloud VPN gateways, protocol forwarding to VMs from an external IP address,
and HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP load balancers.
INTERNAL is used for protocol forwarding to VMs from an internal IP address,
and internal TCP/UDP load balancers.
INTERNAL_MANAGED is used for internal HTTP(S) load balancers. Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL", "INTERNAL_MANAGED"]`,
				Default: "EXTERNAL",
			},
			"network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `For internal load balancing, this field identifies the network that
the load balanced IP should belong to for this Forwarding Rule. If
this field is not specified, the default network will be used.
This field is only used for INTERNAL load balancing.`,
			},
			"network_tier": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PREMIUM", "STANDARD", ""}, false),
				Description: `The networking tier used for configuring this address. If this field is not
specified, it is assumed to be PREMIUM. Possible values: ["PREMIUM", "STANDARD"]`,
			},
			"port_range": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: portRangeDiffSuppress,
				Description: `This field is used along with the target field for TargetHttpProxy,
TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
TargetPool, TargetInstance.

Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
addressed to ports in the specified range will be forwarded to target.
Forwarding rules with the same [IPAddress, IPProtocol] pair must have
disjoint port ranges.

Some types of forwarding target have constraints on the acceptable
ports:

* TargetHttpProxy: 80, 8080
* TargetHttpsProxy: 443
* TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
                  1883, 5222
* TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
                  1883, 5222
* TargetVpnGateway: 500, 4500`,
			},
			"ports": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Description: `This field is used along with the backend_service field for internal
load balancing.

When the load balancing scheme is INTERNAL, a single port or a comma
separated list of ports can be configured. Only packets addressed to
these ports will be forwarded to the backends configured with this
forwarding rule.

You may specify a maximum of up to 5 ports.`,
				MaxItems: 5,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `A reference to the region where the regional forwarding rule resides.
This field is not applicable to global forwarding rules.`,
			},
			"service_label": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `An optional prefix to the service name for this Forwarding Rule.
If specified, will be the first label of the fully qualified service
name.

The label must be 1-63 characters long, and comply with RFC1035.
Specifically, the label must be 1-63 characters long and match the
regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
character must be a lowercase letter, and all following characters
must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.

This field is only used for INTERNAL load balancing.`,
			},
			"subnetwork": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The subnetwork that the load balanced IP should belong to for this
Forwarding Rule.  This field is only used for INTERNAL load balancing.

If the network specified is in auto subnet mode, this field is
optional. However, if the network is in custom subnet mode, a
subnetwork must be specified.`,
			},
			"target": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
				Description: `The URL of the target resource to receive the matched traffic.
The target must live in the same region as the forwarding rule.
The forwarded traffic must be of a type appropriate to the target
object.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"service_name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The internal fully qualified service name for this Forwarding Rule.
This field is only used for INTERNAL load balancing.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeForwardingRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_protocol"); !isEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	backendServiceProp, err := expandComputeForwardingRuleBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(backendServiceProp)) && (ok || !reflect.DeepEqual(v, backendServiceProp)) {
		obj["backendService"] = backendServiceProp
	}
	loadBalancingSchemeProp, err := expandComputeForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeForwardingRuleNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	portRangeProp, err := expandComputeForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port_range"); !isEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	portsProp, err := expandComputeForwardingRulePorts(d.Get("ports"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ports"); !isEmptyValue(reflect.ValueOf(portsProp)) && (ok || !reflect.DeepEqual(v, portsProp)) {
		obj["ports"] = portsProp
	}
	subnetworkProp, err := expandComputeForwardingRuleSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !isEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	targetProp, err := expandComputeForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	allowGlobalAccessProp, err := expandComputeForwardingRuleAllowGlobalAccess(d.Get("allow_global_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_global_access"); ok || !reflect.DeepEqual(v, allowGlobalAccessProp) {
		obj["allowGlobalAccess"] = allowGlobalAccessProp
	}
	allPortsProp, err := expandComputeForwardingRuleAllPorts(d.Get("all_ports"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("all_ports"); !isEmptyValue(reflect.ValueOf(allPortsProp)) && (ok || !reflect.DeepEqual(v, allPortsProp)) {
		obj["allPorts"] = allPortsProp
	}
	networkTierProp, err := expandComputeForwardingRuleNetworkTier(d.Get("network_tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_tier"); !isEmptyValue(reflect.ValueOf(networkTierProp)) && (ok || !reflect.DeepEqual(v, networkTierProp)) {
		obj["networkTier"] = networkTierProp
	}
	serviceLabelProp, err := expandComputeForwardingRuleServiceLabel(d.Get("service_label"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_label"); !isEmptyValue(reflect.ValueOf(serviceLabelProp)) && (ok || !reflect.DeepEqual(v, serviceLabelProp)) {
		obj["serviceLabel"] = serviceLabelProp
	}
	regionProp, err := expandComputeForwardingRuleRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ForwardingRule: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ForwardingRule: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating ForwardingRule",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ForwardingRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ForwardingRule %q: %#v", d.Id(), res)

	return resourceComputeForwardingRuleRead(d, meta)
}

func resourceComputeForwardingRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeForwardingRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeForwardingRuleCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("description", flattenComputeForwardingRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("ip_address", flattenComputeForwardingRuleIPAddress(res["IPAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("ip_protocol", flattenComputeForwardingRuleIPProtocol(res["IPProtocol"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("backend_service", flattenComputeForwardingRuleBackendService(res["backendService"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("load_balancing_scheme", flattenComputeForwardingRuleLoadBalancingScheme(res["loadBalancingScheme"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("name", flattenComputeForwardingRuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("network", flattenComputeForwardingRuleNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("port_range", flattenComputeForwardingRulePortRange(res["portRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("ports", flattenComputeForwardingRulePorts(res["ports"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("subnetwork", flattenComputeForwardingRuleSubnetwork(res["subnetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("target", flattenComputeForwardingRuleTarget(res["target"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("allow_global_access", flattenComputeForwardingRuleAllowGlobalAccess(res["allowGlobalAccess"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("all_ports", flattenComputeForwardingRuleAllPorts(res["allPorts"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("network_tier", flattenComputeForwardingRuleNetworkTier(res["networkTier"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("service_label", flattenComputeForwardingRuleServiceLabel(res["serviceLabel"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("service_name", flattenComputeForwardingRuleServiceName(res["serviceName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("region", flattenComputeForwardingRuleRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ForwardingRule: %s", err)
	}

	return nil
}

func resourceComputeForwardingRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("target") {
		obj := make(map[string]interface{})

		targetProp, err := expandComputeForwardingRuleTarget(d.Get("target"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
			obj["target"] = targetProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}/setTarget")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating ForwardingRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ForwardingRule %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating ForwardingRule",
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("allow_global_access") {
		obj := make(map[string]interface{})

		allowGlobalAccessProp, err := expandComputeForwardingRuleAllowGlobalAccess(d.Get("allow_global_access"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("allow_global_access"); ok || !reflect.DeepEqual(v, allowGlobalAccessProp) {
			obj["allowGlobalAccess"] = allowGlobalAccessProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating ForwardingRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ForwardingRule %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating ForwardingRule",
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeForwardingRuleRead(d, meta)
}

func resourceComputeForwardingRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ForwardingRule %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ForwardingRule")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting ForwardingRule",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ForwardingRule %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeForwardingRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/forwardingRules/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeForwardingRuleCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleIPAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleIPProtocol(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleBackendService(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeForwardingRuleLoadBalancingScheme(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeForwardingRulePortRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRulePorts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeForwardingRuleSubnetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeForwardingRuleTarget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleAllowGlobalAccess(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleAllPorts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleNetworkTier(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleServiceLabel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleServiceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeForwardingRuleRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeForwardingRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIPAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIPProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleBackendService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	// This method returns a full self link from a partial self link.
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		// Anything that starts with a URL scheme is assumed to be a self link worth using.
		return v, nil
	} else if strings.HasPrefix(v.(string), "projects/") {
		// If the self link references a project, we'll just stuck the compute prefix on it
		url, err := replaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
		if err != nil {
			return "", err
		}
		return url, nil
	} else if strings.HasPrefix(v.(string), "regions/") || strings.HasPrefix(v.(string), "zones/") {
		// For regional or zonal resources which include their region or zone, just put the project in front.
		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/")
		if err != nil {
			return nil, err
		}
		return url + v.(string), nil
	}
	// Anything else is assumed to be a regional resource, with a partial link that begins with the resource name.
	// This isn't very likely - it's a last-ditch effort to extract something useful here.  We can do a better job
	// as soon as MultiResourceRefs are working since we'll know the types that this field is supposed to point to.
	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/")
	if err != nil {
		return nil, err
	}
	return url + v.(string), nil
}

func expandComputeForwardingRuleLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeForwardingRulePortRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRulePorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeForwardingRuleSubnetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeForwardingRuleTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	// This method returns a full self link from a partial self link.
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		// Anything that starts with a URL scheme is assumed to be a self link worth using.
		return v, nil
	} else if strings.HasPrefix(v.(string), "projects/") {
		// If the self link references a project, we'll just stuck the compute prefix on it
		url, err := replaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
		if err != nil {
			return "", err
		}
		return url, nil
	} else if strings.HasPrefix(v.(string), "regions/") || strings.HasPrefix(v.(string), "zones/") {
		// For regional or zonal resources which include their region or zone, just put the project in front.
		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/")
		if err != nil {
			return nil, err
		}
		return url + v.(string), nil
	}
	// Anything else is assumed to be a regional resource, with a partial link that begins with the resource name.
	// This isn't very likely - it's a last-ditch effort to extract something useful here.  We can do a better job
	// as soon as MultiResourceRefs are working since we'll know the types that this field is supposed to point to.
	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/")
	if err != nil {
		return nil, err
	}
	return url + v.(string), nil
}

func expandComputeForwardingRuleAllowGlobalAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleAllPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNetworkTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceLabel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
