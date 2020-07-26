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
)

func resourceIapClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceIapClientCreate,
		Read:   resourceIapClientRead,
		Delete: resourceIapClientDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIapClientImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"brand": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Identifier of the brand to which this client
is attached to. The format is
'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Human-friendly name given to the OAuth client.`,
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Unique identifier of the OAuth client.`,
			},
			"secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Client secret of the OAuth client.`,
				Sensitive:   true,
			},
		},
	}
}

func resourceIapClientCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	displayNameProp, err := expandIapClientDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := replaceVars(d, config, "{{IapBasePath}}{{brand}}/identityAwareProxyClients")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Client: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Client: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Client %q: %#v", d.Id(), res)

	brand := d.Get("brand")
	clientId := flattenIapClientClientId(res["name"], d, config)

	d.Set("client_id", clientId)
	d.SetId(fmt.Sprintf("%s/identityAwareProxyClients/%s", brand, clientId))

	return resourceIapClientRead(d, meta)
}

func resourceIapClientRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{IapBasePath}}{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("IapClient %q", d.Id()))
	}

	if err := d.Set("secret", flattenIapClientSecret(res["secret"], d, config)); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}
	if err := d.Set("display_name", flattenIapClientDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}
	if err := d.Set("client_id", flattenIapClientClientId(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}

	return nil
}

func resourceIapClientDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	url, err := replaceVars(d, config, "{{IapBasePath}}{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Client %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Client")
	}

	log.Printf("[DEBUG] Finished deleting Client %q: %#v", d.Id(), res)
	return nil
}

func resourceIapClientImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<brand>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("brand").(string), "/")
	if len(nameParts) != 6 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("brand").(string),
			"projects/{{project_number}}/brands/{{brand_id}}/identityAwareProxyClients/{{client_id}}",
		)
	}

	d.Set("brand", fmt.Sprintf("projects/%s/brands/%s", nameParts[1], nameParts[3]))
	d.Set("client_id", nameParts[5])
	return []*schema.ResourceData{d}, nil
}

func flattenIapClientSecret(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIapClientDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIapClientClientId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandIapClientDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
