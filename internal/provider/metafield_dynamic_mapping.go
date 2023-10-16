package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/utils"
)

func resourceMetafieldDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MetafieldDynamicMapping.",

		CreateContext: resourceMetafieldDynamicMappingCreate,
		ReadContext:   resourceMetafieldDynamicMappingRead,
		UpdateContext: resourceMetafieldDynamicMappingUpdate,
		DeleteContext: resourceMetafieldDynamicMappingDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the metafield.",
				Required:    true,
			},
			"value": {
				Type:        schema.TypeString,
				Description: "Dynamic value",
				Optional:    true,
			},
			"scope": {
				Type:        schema.TypeList,
				Description: "Dynamic mapping.",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Device name.",
							Optional:    true,
						},
						"vdom": {
							Type:        schema.TypeString,
							Description: "Device vdom.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceMetafieldDynamicMappingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if s, ok := v.(bool); ok {
			allow_append = s
		}
	}
	urlparams.AllowAppend = &allow_append

	if allow_append {
		d.SetId(fmt.Sprintf("%s/%s", d.Get("scope.0.name").(string), d.Get("scope.0.vdom").(string)))
		return resourceMetafieldDynamicMappingUpdate(ctx, d, meta)
	}

	deviceparam := ""
	if v, ok := d.GetOk("deviceparam"); ok {
		if s, ok := v.(string); ok {
			deviceparam = s
		}
	}
	urlparams.Scope = deviceparam

	name := d.Get("name").(string)

	obj, diags := getObjectMetafieldDynamicMapping(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err = c.Cmdb.CreateMetafieldDynamicMapping(name, obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId(fmt.Sprintf("%s/%s", d.Get("scope.0.name").(string), d.Get("scope.0.vdom").(string)))

	return resourceMetafieldDynamicMappingRead(ctx, d, meta)
}

func resourceMetafieldDynamicMappingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
	c := meta.(*apiClient).Client
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

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if s, ok := v.(bool); ok {
			allow_append = s
		}
	}
	urlparams.AllowAppend = &allow_append

	deviceparam := ""
	if v, ok := d.GetOk("deviceparam"); ok {
		if s, ok := v.(string); ok {
			deviceparam = s
		}
	}
	urlparams.Scope = deviceparam

	name := d.Get("name").(string)

	obj, diags := getObjectMetafieldDynamicMapping(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateMetafieldDynamicMapping(name, mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating MetafieldDynamicMapping resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	d.SetId(fmt.Sprintf("%s/%s", d.Get("scope.0.name").(string), d.Get("scope.0.vdom").(string)))

	return resourceMetafieldDynamicMappingRead(ctx, d, meta)
}

func resourceMetafieldDynamicMappingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
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

	name := d.Get("name").(string)

	err := c.Cmdb.DeleteMetafieldDynamicMapping(name, mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting MetafieldDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceMetafieldDynamicMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
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

	name := d.Get("name").(string)

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadMetafieldDynamicMapping(name, mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading MetafieldDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectMetafieldDynamicMapping(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenMetafieldDynamicMappingScope(d *schema.ResourceData, v *[]models.Scope, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}
			if tmp := cfg.Vdom; tmp != nil {
				v["vdom"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectMetafieldDynamicMapping(d *schema.ResourceData, o *models.DynamicMapping, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Value != nil {
		v := *o.Value

		if err = d.Set("value", v); err != nil {
			return diag.Errorf("error reading value: %v", err)
		}
	}

	if o.Scope != nil {
		if err = d.Set("scope", flattenMetafieldDynamicMappingScope(d, &o.Scope, "scope", sort)); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	return nil
}

func expandMetafieldDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.Scope, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.Scope

	for i := range l {
		tmp := models.Scope{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vdom", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vdom = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectMetafieldDynamicMapping(d *schema.ResourceData, sv string) (*models.DynamicMapping, diag.Diagnostics) {
	obj := models.DynamicMapping{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("value"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("value", sv)
				diags = append(diags, e)
			}
			obj.Value = &v2
		}
	}

	if v, ok := d.GetOk("scope"); ok {
		t, err := expandMetafieldDynamicMappingScope(d, v, "scope", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Scope = *t
		}
	} else if d.HasChange("scope") {
		old, new := d.GetChange("scope")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Scope = []models.Scope{}
		}
	}
	return &obj, diags
}

// // Return an object with explicitly empty objects for tables that have been set.
// func getEmptyObjectMetafieldDynamicMapping(d *schema.ResourceData, sv string) (*models.DynamicMapping, diag.Diagnostics) {
// 	obj := models.DynamicMapping{}
// 	diags := diag.Diagnostics{}

// 	obj.Scope = []models.DynamicMapping{}

// 	return &obj, diags
// }
