package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceContactPoints() *schema.Resource {
	return &schema.Resource{
		Create: resourceContactPointsCreate,
		Read:   resourceContactPointsRead,
		Update: resourceContactPointsUpdate,
		Delete: resourceContactPointsDelete,

		Schema: map[string]*schema.Schema{
			"disable_resolve_message": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provenance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"settings": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func convertResourceDataToContactPoint(d *schema.ResourceData) *ContactPoint {
	contactPoint := &ContactPoint{}
	if v, ok := d.GetOk("disable_resolve_message"); ok {
		contactPoint.DisableResolveMessage = v.(bool)
	}
	if v, ok := d.GetOk("name"); ok {
		contactPoint.Name = v.(string)
	}
	if v, ok := d.GetOk("provenance"); ok {
		contactPoint.Provenance = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		contactPoint.Type = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		contactPoint.UID = v.(string)
	}
	if v, ok := d.GetOk("settings"); ok {
		contactPoint.Settings = v.(map[string]interface{})
	}
	return contactPoint

}

func resourceContactPointsCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*Client)

	contactPoint := convertResourceDataToContactPoint(d)
	id, err := client.NewContactPoint(contactPoint)
	if err != nil {
		return err
	}
	d.SetId(id)

	return resourceContactPointsRead(d, m)

}

func resourceContactPointsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	id := d.Id()

	contactPoint, err := client.ContactPoint(id)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(contactPoint.UID)
	return nil
}

func resourceContactPointsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	contactPoint := convertResourceDataToContactPoint(d)
	err := client.UpdateContactPoint(contactPoint)
	if err != nil {
		return err
	}
	return resourceContactPointsRead(d, m)
}

func resourceContactPointsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	id := d.Id()
	err := client.DeleteContactPoint(id)
	if err != nil {
		return err
	}
	return nil
}
