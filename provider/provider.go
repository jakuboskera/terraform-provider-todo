package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jakuboskera/terraform-provider-todo/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TODO_URL", ""),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TODO_API_KEY", ""),
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				DefaultFunc: schema.EnvDefaultFunc("TODO_IGNORE_CERT", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"todo_task": resourceTask(),
		},
		// DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var apiPath string = "/api/v1"

	url := d.Get("url").(string)
	api_key := d.Get("api_key").(string)
	insecure := d.Get("insecure").(bool)

	if strings.HasSuffix(url, "/") {
		url = strings.Trim(url, "/")
	}

	return client.NewClient(url+apiPath, api_key, insecure), nil
}
