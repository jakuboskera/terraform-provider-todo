package client

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jakuboskera/terraform-provider-todo/models"
)

// TaskBody return a json body
func TaskBody(d *schema.ResourceData) models.TaskBody {
	return models.TaskBody{
		Text:   d.Get("text").(string),
		IsDone: d.Get("is_done").(bool),
	}
}
