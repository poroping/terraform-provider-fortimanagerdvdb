// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS custom signature.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/suppressors"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/utils"
)

func resourceIpsCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS custom signature.",

		CreateContext: resourceIpsCustomCreate,
		ReadContext:   resourceIpsCustomRead,
		UpdateContext: resourceIpsCustomUpdate,
		DeleteContext: resourceIpsCustomDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
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
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "block"}, false),

				Description: "Default action (pass or block) for this signature.",
				Optional:    true,
				Computed:    true,
			},
			"application": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Applications to be protected. Blank for all applications.",
				Optional:         true,
				Computed:         true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Protect client or server traffic.",
				Optional:         true,
				Computed:         true,
			},
			"log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging.",
				Optional:    true,
				Computed:    true,
			},
			"log_packet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable packet logging.",
				Optional:    true,
				Computed:    true,
			},
			"os": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Operating system(s) that the signature protects. Blank for all operating systems.",
				Optional:         true,
				Computed:         true,
			},
			"protocol": {
				Type: schema.TypeString,

				Description: "Protocol(s) that the signature scans. Blank for all protocols.",
				Optional:    true,
				Computed:    true,
			},
			"rule_id": {
				Type: schema.TypeInt,

				Description: "Signature ID.",
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Type: schema.TypeString,

				Description: "Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.",
				Optional:    true,
				Computed:    true,
			},
			"signature": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),

				Description: "Custom signature enclosed in single quotes.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable this signature.",
				Optional:    true,
				Computed:    true,
			},
			"tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Signature tag.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceIpsCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
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

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "tag"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating IpsCustom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectIpsCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIpsCustom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsCustom")
	}

	return resourceIpsCustomRead(ctx, d, meta)
}

func resourceIpsCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
	c := meta.(*apiClient).Client
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
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

	obj, diags := getObjectIpsCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIpsCustom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IpsCustom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsCustom")
	}

	return resourceIpsCustomRead(ctx, d, meta)
}

func resourceIpsCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
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

	err := c.Cmdb.DeleteIpsCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IpsCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
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

	o, err := c.Cmdb.ReadIpsCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsCustom resource: %v", err)
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

	diags := refreshObjectIpsCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectIpsCustom(d *schema.ResourceData, o *models.IpsCustom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Application != nil {
		v := *o.Application

		if err = d.Set("application", v); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Location != nil {
		v := *o.Location

		if err = d.Set("location", v); err != nil {
			return diag.Errorf("error reading location: %v", err)
		}
	}

	if o.Log != nil {
		v := *o.Log

		if err = d.Set("log", v); err != nil {
			return diag.Errorf("error reading log: %v", err)
		}
	}

	if o.LogPacket != nil {
		v := *o.LogPacket

		if err = d.Set("log_packet", v); err != nil {
			return diag.Errorf("error reading log_packet: %v", err)
		}
	}

	if o.Os != nil {
		v := *o.Os

		if err = d.Set("os", v); err != nil {
			return diag.Errorf("error reading os: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.RuleId != nil {
		v := *o.RuleId

		if err = d.Set("rule_id", v); err != nil {
			return diag.Errorf("error reading rule_id: %v", err)
		}
	}

	if o.Severity != nil {
		v := *o.Severity

		if err = d.Set("severity", v); err != nil {
			return diag.Errorf("error reading severity: %v", err)
		}
	}

	if o.Signature != nil {
		v := *o.Signature

		if err = d.Set("signature", v); err != nil {
			return diag.Errorf("error reading signature: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tag != nil {
		v := *o.Tag

		if err = d.Set("tag", v); err != nil {
			return diag.Errorf("error reading tag: %v", err)
		}
	}

	return nil
}

func getObjectIpsCustom(d *schema.ResourceData, sv string) (*models.IpsCustom, diag.Diagnostics) {
	obj := models.IpsCustom{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("application"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application", sv)
				diags = append(diags, e)
			}
			obj.Application = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("location", sv)
				diags = append(diags, e)
			}
			obj.Location = &v2
		}
	}
	if v1, ok := d.GetOk("log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log", sv)
				diags = append(diags, e)
			}
			obj.Log = &v2
		}
	}
	if v1, ok := d.GetOk("log_packet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_packet", sv)
				diags = append(diags, e)
			}
			obj.LogPacket = &v2
		}
	}
	if v1, ok := d.GetOk("os"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("os", sv)
				diags = append(diags, e)
			}
			obj.Os = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v1, ok := d.GetOk("rule_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rule_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RuleId = &tmp
		}
	}
	if v1, ok := d.GetOk("severity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("severity", sv)
				diags = append(diags, e)
			}
			obj.Severity = &v2
		}
	}
	if v1, ok := d.GetOk("signature"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("signature", sv)
				diags = append(diags, e)
			}
			obj.Signature = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tag", sv)
				diags = append(diags, e)
			}
			obj.Tag = &v2
		}
	}
	return &obj, diags
}
