package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func resourceInstallPackage() *schema.Resource {
	return &schema.Resource{
		Description: "Configure InstallPackage.",

		CreateContext: resourceInstallPackageCreate,
		ReadContext:   resourceInstallPackageRead,
		// UpdateContext: resourceInstallPackageUpdate,
		DeleteContext: resourceInstallPackageDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"adom": {
				Type:        schema.TypeString,
				Description: "Specifies the ADOM to target.",
				Optional:    true,
				ForceNew:    true,
				Default:     "root",
			},
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the VDOM to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"deviceparam": {
				Type:        schema.TypeString,
				Description: "Specifies the device to which the resource will be applied to on the FortiManager. If you want to inherit the configuration from the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"device": {
				Type:        schema.TypeString,
				Description: "Name of the device.",
				Required:    true,
				ForceNew:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "Name of VDOM on device",
				Optional:    true,
				ForceNew:    true,
				Default:     "root",
			},
			"package": {
				Type:        schema.TypeString,
				Description: "Name of package to install.",
				Required:    true,
				ForceNew:    true,
			},
			"flags": {
				Type:        schema.TypeList,
				Description: "Dynamic mapping.",
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"force_recreate": {
				Type:        schema.TypeString,
				Description: "Changing this value will force package reinstall",
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceInstallPackageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	adom := ""
	if v, ok := d.GetOk("adom"); ok {
		if s, ok := v.(string); ok {
			adom = s
		}
	}
	urlparams.Adom = adom

	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	deviceparam := ""
	if v, ok := d.GetOk("deviceparam"); ok {
		if s, ok := v.(string); ok {
			deviceparam = s
		}
	}
	urlparams.Scope = deviceparam

	obj, diags := getObjectInstallPackage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err = c.Cmdb.CreateInstallPackage(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId(d.Get("device").(string))

	return diags
}

// func resourceInstallPackageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	mkey := d.Id()
// 	c := meta.(*apiClient).Client
// 	// c.Retries = 1
// 	urlparams := &models.CmdbRequestParams{}
// 	adom := ""
// 	if v, ok := d.GetOk("adom"); ok {
// 		if s, ok := v.(string); ok {
// 			adom = s
// 		}
// 	}
// 	urlparams.Adom = adom

// 	vdomparam := ""
// 	if v, ok := d.GetOk("vdomparam"); ok {
// 		if s, ok := v.(string); ok {
// 			vdomparam = s
// 		}
// 	}
// 	urlparams.Vdom = vdomparam

// 	allow_append := false
// 	if v, ok := d.GetOk("allow_append"); ok {
// 		if s, ok := v.(bool); ok {
// 			allow_append = s
// 		}
// 	}
// 	urlparams.AllowAppend = &allow_append

// 	deviceparam := ""
// 	if v, ok := d.GetOk("deviceparam"); ok {
// 		if s, ok := v.(string); ok {
// 			deviceparam = s
// 		}
// 	}
// 	urlparams.Scope = deviceparam

// 	obj, diags := getObjectInstallPackage(d, c.Config.Fv)
// 	if diags.HasError() {
// 		return diags
// 	}

// 	_, err := c.Cmdb.UpdateInstallPackage(mkey, obj, urlparams)
// 	if err != nil {
// 		return diag.Errorf("error updating InstallPackage resource: %v", err)
// 	}

// 	// log.Printf(strconv.Itoa(c.Retries))
// 	d.SetId(d.Get("name").(string))

// 	return resourceInstallPackageRead(ctx, d, meta)
// }

func resourceInstallPackageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourceInstallPackageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// 	mkey := d.Id()

	// 	c := meta.(*apiClient).Client
	// 	// c.Retries = 1

	// 	urlparams := &models.CmdbRequestParams{}
	// 	adom := ""
	// 	if v, ok := d.GetOk("adom"); ok {
	// 		if s, ok := v.(string); ok {
	// 			adom = s
	// 		}
	// 	}
	// 	urlparams.Adom = adom

	// 	vdomparam := ""
	// 	if v, ok := d.GetOk("vdomparam"); ok {
	// 		if s, ok := v.(string); ok {
	// 			vdomparam = s
	// 		}
	// 	}
	// 	urlparams.Vdom = vdomparam

	// 	deviceparam := ""
	// 	if v, ok := d.GetOk("deviceparam"); ok {
	// 		if s, ok := v.(string); ok {
	// 			deviceparam = s
	// 		}
	// 	}
	// 	urlparams.Scope = deviceparam

	// 	ptp := true
	// 	urlparams.PlainTextPassword = &ptp

	// 	o, err := c.Cmdb.ReadInstallPackage(mkey, urlparams)
	// 	if err != nil {
	// 		return diag.Errorf("error reading InstallPackage resource: %v", err)
	// 	}

	// 	if o == nil {
	// 		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
	// 		d.SetId("")
	// 		return nil
	// 	}

	// 	sort := false
	// 	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
	// 		if b, ok := v.(bool); ok {
	// 			sort = b
	// 		}
	// 	}

	// diags := refreshObjectInstallPackage(d, o, c.Config.Fv, sort)
	// if diags.HasError() {
	// 	return diags
	// }
	return diag.Diagnostics{}
}

// func refreshObjectInstallPackage(d *schema.ResourceData, o *models.InstallPackage, sv string, sort bool) diag.Diagnostics {
// 	var err error

// 	if o.Description != nil {
// 		v := *o.Description

// 		if err = d.Set("description", v); err != nil {
// 			return diag.Errorf("error reading description: %v", err)
// 		}
// 	}

// 	if o.Name != nil {
// 		v := *o.Name

// 		if err = d.Set("name", v); err != nil {
// 			return diag.Errorf("error reading name: %v", err)
// 		}
// 	}

// 	if o.Value != nil {
// 		v := *o.Value

// 		if err = d.Set("value", v); err != nil {
// 			return diag.Errorf("error reading value: %v", err)
// 		}
// 	}

// 	if o.DynamicMapping != nil {
// 		if err = d.Set("dynamic_mapping", flattenInstallPackageDynamicMapping(d, &o.DynamicMapping, "dynamic_mapping", sort)); err != nil {
// 			return diag.Errorf("error reading dynamic_mapping: %v", err)
// 		}
// 	}

// 	return nil
// }

func getObjectInstallPackage(d *schema.ResourceData, sv string) (*models.InstallPackagePayload, diag.Diagnostics) {
	obj := models.InstallPackagePayload{}
	diags := diag.Diagnostics{}

	scope := models.Scope{}

	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			scope.Name = &v2
		}
	}

	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(string); ok {
			scope.Vdom = &v2
		}
	}

	obj.Scope = []models.Scope{scope}

	if v1, ok := d.GetOk("adom"); ok {
		if v2, ok := v1.(string); ok {
			obj.Adom = &v2
		}
	}

	if v1, ok := d.GetOk("package"); ok {
		if v2, ok := v1.(string); ok {
			obj.Pkg = &v2
		}
	}

	if v1, ok := d.GetOk("flags"); ok {
		if v2, ok := v1.([]string); ok {
			obj.Flags = v2
		}
	}

	return &obj, diags
}
