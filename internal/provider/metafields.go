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

func resourceMetafields() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Metafields.",

		CreateContext: resourceMetafieldsCreate,
		ReadContext:   resourceMetafieldsRead,
		UpdateContext: resourceMetafieldsUpdate,
		DeleteContext: resourceMetafieldsDelete,

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
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of the metafield.",
				Optional:    true,
			},
			"value": {
				Type:        schema.TypeString,
				Description: "Default value of the metafield.",
				Optional:    true,
			},
			"dynamic_mapping": {
				Type:        schema.TypeList,
				Description: "Dynamic Mapping.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:        schema.TypeString,
							Description: "Dynamic value",
							Optional:    true,
						},
						"scope": {
							Type:        schema.TypeList,
							Description: "Dynamic mapping.",
							Optional:    true,
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
				},
			},
		},
	}
}

func resourceMetafieldsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return resourceMetafieldsUpdate(ctx, d, meta)
	}

	deviceparam := ""
	if v, ok := d.GetOk("deviceparam"); ok {
		if s, ok := v.(string); ok {
			deviceparam = s
		}
	}
	urlparams.Scope = deviceparam

	obj, diags := getObjectMetafields(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err = c.Cmdb.CreateMetafields(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId(d.Get("name").(string))

	return resourceMetafieldsRead(ctx, d, meta)
}

func resourceMetafieldsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectMetafields(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateMetafields(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating Metafields resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	d.SetId(d.Get("name").(string))

	return resourceMetafieldsRead(ctx, d, meta)
}

func resourceMetafieldsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteMetafields(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting Metafields resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceMetafieldsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadMetafields(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading Metafields resource: %v", err)
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

	diags := refreshObjectMetafields(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenMetafieldsDynamicMapping(d *schema.ResourceData, v *[]models.DynamicMapping, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			if tmp := cfg.Scope; tmp != nil {
				v["scope"] = flattenMetafieldsDynamicMappingScope(d, &tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "scope"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func flattenMetafieldsDynamicMappingScope(d *schema.ResourceData, v *[]models.Scope, prefix string, sort bool) interface{} {
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

func refreshObjectMetafields(d *schema.ResourceData, o *models.Metafields, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Value != nil {
		v := *o.Value

		if err = d.Set("value", v); err != nil {
			return diag.Errorf("error reading value: %v", err)
		}
	}

	if o.DynamicMapping != nil {
		if err = d.Set("dynamic_mapping", flattenMetafieldsDynamicMapping(d, &o.DynamicMapping, "dynamic_mapping", sort)); err != nil {
			return diag.Errorf("error reading dynamic_mapping: %v", err)
		}
	}

	return nil
}

func expandMetafieldsDynamicMapping(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DynamicMapping, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DynamicMapping

	for i := range l {
		tmp := models.DynamicMapping{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scope", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandMetafieldsDynamicMappingScope(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.MetafieldsexpandMetafieldsDynamicMappingScope
			// 	}
			tmp.Scope = *v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandMetafieldsDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.Scope, error) {
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

func getObjectMetafields(d *schema.ResourceData, sv string) (*models.Metafields, diag.Diagnostics) {
	obj := models.Metafields{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}

	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}

	if v1, ok := d.GetOk("value"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("value", sv)
				diags = append(diags, e)
			}
			obj.Value = &v2
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok {
		t, err := expandMetafieldsDynamicMapping(d, v, "dynamic_mapping", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DynamicMapping = *t
		}
	} else if d.HasChange("dynamic_mapping") {
		old, new := d.GetChange("dynamic_mapping")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DynamicMapping = []models.DynamicMapping{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectMetafields(d *schema.ResourceData, sv string) (*models.Metafields, diag.Diagnostics) {
	obj := models.Metafields{}
	diags := diag.Diagnostics{}

	obj.DynamicMapping = []models.DynamicMapping{}

	return &obj, diags
}
