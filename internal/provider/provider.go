package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/auth"
	"github.com/poroping/fortimanager-devicedb-sdk-go/client"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		if s.Required {
			desc = "(Required) " + desc
		}
		if s.ForceNew {
			desc += " Changing this forces a new resource to be created."
		}
		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"hostname": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The hostname/IP address of the FortiOS to be connected.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_ACCESS_HOSTNAME", ""),
				},
				"username": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The FortiOS API access username.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_ACCESS_USERNAME", ""),
				},
				"password": {
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					Description: "The FortiOS API access password.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_ACCESS_PASSWORD", ""),
				},
				"insecure": {
					Type:        schema.TypeBool,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_INSECURE", false),
				},
				"cabundlefile": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_CA_CABUNDLE", ""),
				},
				"peerauth": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Enable/disable peer authentication, can be 'enable' or 'disable'.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_CA_PEERAUTH", ""),
				},
				"cacert": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "CA certificate(Optional).",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_CA_CACERT", ""),
				},
				"clientcert": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "User certificate.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_CA_CLIENTCERT", ""),
				},
				"clientkey": {
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					Description: "User private key.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_CA_CLIENTKEY", ""),
				},
				"vdom": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default VDOM for API calls.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_VDOM", ""),
				},
				"os_version": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Manually set the FortiOS version. Will overwrite auto_version.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_OS_VERSION", nil),
				},
				"auto_version": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Will determine FortiOS version on client creation.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_AUTO_VERSION", true),
				},
				"default_device": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Set default target device.",
					DefaultFunc: schema.EnvDefaultFunc("TF_FMG_DEVICEDB_DEFAULT_DEVICE", ""),
				},
			},
			DataSourcesMap: providerDataSourcesMerged(),
			ResourcesMap:   providerResourcesMerged(),
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func providerDataSourcesMerged() map[string]*schema.Resource {
	a := providerDataSources()
	// b := providerDataSourcesAliases()
	// c := providerDataSourcesCustom()
	// for k, v := range b {
	// 	a[k] = v
	// }
	// for k, v := range c {
	// 	a[k] = v
	// }
	delete(a, "fortios_firewall_vendormacsummary") // shit resource that is empty?
	return a
}

func providerResourcesMerged() map[string]*schema.Resource {
	a := providerResources()
	// b := providerResourcesAliases()
	// c := providerResourcesCustom()
	// for k, v := range b {
	// 	a[k] = v
	// }
	// for k, v := range c {
	// 	a[k] = v
	// }
	delete(a, "fortios_firewall_vendormacsummary") // shit resource that is empty?
	return a
}

type apiClient struct {
	Client *client.FortiSDKClient
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		auth := &auth.Auth{
			Hostname: d.Get("hostname").(string),
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
			Vdom:     d.Get("vdom").(string),
			CABundle: d.Get("cabundlefile").(string),
			Insecure: d.Get("insecure").(bool),

			PeerAuth:   d.Get("peerauth").(string),
			CaCert:     d.Get("cacert").(string),
			ClientCert: d.Get("clientcert").(string),
			ClientKey:  d.Get("clientkey").(string),
		}

		tmp, err := client.NewClient(auth)
		diags = diag.FromErr(err)

		c := &apiClient{
			Client: tmp,
		}

		userAgent := p.UserAgent("terraform-provider-fortimanagerdvdb", version)
		log.Printf("[DEBUG] Setting User Agent to %s", userAgent)
		c.Client.Config.UserAgent = userAgent

		dd := d.Get("default_device").(string)
		c.Client.Config.FwTarget = dd

		if v, ok := d.GetOk("os_version"); ok {
			log.Printf("[INFO] Manually setting FortiOS version to %s", v.(string))
			c.Client.Config.Fv = v.(string)
		}

		return c, diags
	}
}
