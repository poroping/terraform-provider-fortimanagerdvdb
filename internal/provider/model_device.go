package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/utils"
)

func resourceModelDevice() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Model Device.",

		CreateContext: resourceModelDeviceCreate,
		ReadContext:   resourceModelDeviceRead,
		UpdateContext: resourceModelDeviceUpdate,
		DeleteContext: resourceModelDeviceDelete,

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
				Description: "Name of the device.",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of the device.",
				Optional:    true,
			},
			"serial": {
				Type:        schema.TypeString,
				Description: "Serial number of the device.",
				Required:    true,
				ForceNew:    true,
			},
			"version": {
				Type:        schema.TypeString,
				Description: "Version of the device.",
				Optional:    true,
			},
			"device_blueprint": {
				Type:        schema.TypeString,
				Description: "Device blueprint of the device.",
				Optional:    true,
			},
			"authorization_template": {
				Type:        schema.TypeString,
				Description: "Authz template of the device.",
				Optional:    true,
			},
			"prefer_img_ver": {
				Type:        schema.TypeString,
				Description: "Preferred image version of the device.",
				Optional:    true,
			},
			"mgmt_mode": {
				Type:        schema.TypeString,
				Description: "mgmt_mode of the device.",
				Optional:    true,
			},
			"os_type": {
				Type:        schema.TypeString,
				Description: "os_type of the device.",
				Optional:    true,
			},
		},
	}
}

func resourceModelDeviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		d.SetId(d.Get("name").(string))
		return resourceModelDeviceUpdate(ctx, d, meta)
	}

	deviceparam := ""
	if v, ok := d.GetOk("deviceparam"); ok {
		if s, ok := v.(string); ok {
			deviceparam = s
		}
	}
	urlparams.Scope = deviceparam

	obj, diags := getObjectModelDevice(d, c.Config.Fv)

	da := "add_model"
	obj.DeviceAction = &da
	flags := []string{"linked_to_model", "is_model"}
	obj.Flags = flags
	adm := "admin"
	obj.AdmUSR = &adm

	if diags.HasError() {
		return diags
	}

	obj2 := &models.ModelDevice{}
	obj2.Adom = &adom
	obj2.Device = obj
	_, err = c.Cmdb.CreateModelDevice(obj2, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId(d.Get("name").(string))

	return resourceModelDeviceRead(ctx, d, meta)
}

func resourceModelDeviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectModelDevice(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	obj2 := &models.ModelDevice{}
	obj2.Adom = &adom
	obj2.Device = obj

	_, err := c.Cmdb.UpdateModelDevice(mkey, obj2, urlparams)
	if err != nil {
		return diag.Errorf("error updating Model Device resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	d.SetId(d.Get("name").(string))

	return resourceModelDeviceRead(ctx, d, meta)
}

func resourceModelDeviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteModelDevice(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting Model Device resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceModelDeviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadModelDevice(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading Model Device resource: %v", err)
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

	diags := refreshObjectDevice(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectDevice(d *schema.ResourceData, o *models.Device, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Desc != nil {
		v := *o.Desc

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

	if o.Sn != nil {
		v := *o.Sn

		if err = d.Set("serial", v); err != nil {
			return diag.Errorf("error reading serial: %v", err)
		}
	}

	if o.PreferImgVer != nil {
		v := *o.PreferImgVer

		if err = d.Set("prefer_img_ver", v); err != nil {
			return diag.Errorf("error reading prefer_img_ver: %v", err)
		}
	}

	if o.MgmtMode != nil {
		v := *o.MgmtMode

		if err = d.Set("mgmt_mode", v); err != nil {
			return diag.Errorf("error reading mgmt_mode: %v", err)
		}
	}

	if o.OSType != nil {
		v := *o.OSType

		if err = d.Set("os_type", v); err != nil {
			return diag.Errorf("error reading os_type: %v", err)
		}
	}

	if o.DeviceBlueprint != nil {
		v := *o.DeviceBlueprint

		if err = d.Set("device_blueprint", v); err != nil {
			return diag.Errorf("error reading device_blueprint: %v", err)
		}
	}

	if o.OSVer != nil && o.Mr != nil && o.Patch != nil {
		osver := *o.OSVer
		mr := *o.Mr
		patch := *o.Patch

		v := fmt.Sprintf("%s.%d.%d", strings.Split(osver, ".")[0], mr, patch)

		if err = d.Set("version", v); err != nil {
			return diag.Errorf("error reading version: %v", err)
		}
	}

	return nil
}

func getObjectModelDevice(d *schema.ResourceData, sv string) (*models.Device, diag.Diagnostics) {
	obj := models.Device{}
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
			obj.Desc = &v2
		}
	}

	if v1, ok := d.GetOk("serial"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("serial", sv)
				diags = append(diags, e)
			}
			obj.Sn = &v2
		}
	}

	if v1, ok := d.GetOk("prefer_img_ver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefer_img_ver", sv)
				diags = append(diags, e)
			}
			obj.PreferImgVer = &v2
		}
	}

	if v1, ok := d.GetOk("mgmt_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mgmt_mode", sv)
				diags = append(diags, e)
			}
			obj.MgmtMode = &v2
		}
	}

	if v1, ok := d.GetOk("os_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("os_type", sv)
				diags = append(diags, e)
			}
			obj.OSType = &v2
		}
	}

	if v1, ok := d.GetOk("device_blueprint"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device_blueprint", sv)
				diags = append(diags, e)
			}
			obj.DeviceBlueprint = &v2
		}
	}

	if v1, ok := d.GetOk("version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("version", sv)
				diags = append(diags, e)
			}
			v3 := strings.Split(v2, ".")
			osver := fmt.Sprintf("%s.0", v3[0])
			obj.OSVer = &osver
			mr, err := strconv.Atoi(v3[1])
			if err != nil {
				obj.Mr = nil
			} else {
				mr2 := int64(mr)
				obj.Mr = &mr2
			}
			patch, err := strconv.Atoi(v3[2])
			if err != nil {
				obj.Patch = nil
			} else {
				patch2 := int64(patch)
				obj.Patch = &patch2
			}
		}
	}

	return &obj, diags
}
