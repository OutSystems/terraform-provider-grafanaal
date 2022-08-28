package main

import (
	"encoding/json"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceNotificationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotificationPolicyCreate,
		Read:   resourceNotificationPolicyRead,
		Update: resourceNotificationPolicyUpdate,
		Delete: resourceNotificationPolicyDelete,

		Schema: map[string]*schema.Schema{
			"continue": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by_str": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mute_time_intervals": {
				Type:     schema.TypeString,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"receiver": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"routes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_interval": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_wait": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_matchers": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provenance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repeat_interval": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func convertResourceDataToNotificationPolicy(d *schema.ResourceData) *SpecificPolicy {
	notificationPolicy := &SpecificPolicy{}

	if v, ok := d.GetOk("receiver"); ok {
		notificationPolicy.Receiver = v.(string)
	}
	if v, ok := d.GetOk("group_by_str"); ok {
		notificationPolicy.GroupBy = v.([]string)
	}
	if v, ok := d.GetOk("object_matchers"); ok {

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(v.(string)), &jsonMap)
		notificationPolicy.ObjectMatchers = jsonMap
	}
	if v, ok := d.GetOk("mute_time_intervals"); ok {
		notificationPolicy.MuteTimeIntervals = v.([]string)
	}
	if v, ok := d.GetOk("continue"); ok {
		notificationPolicy.Continue = v.(bool)
	}
	if v, ok := d.GetOk("provenance"); ok {
		notificationPolicy.Provenance = v.(string)
	}
	if v, ok := d.GetOk("group_interval"); ok {
		convertedInt, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		notificationPolicy.GroupWait = convertedInt
	}
	if v, ok := d.GetOk("group_wait"); ok {
		convertedInt, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		notificationPolicy.GroupInterval = convertedInt
	}
	if v, ok := d.GetOk("repeat_interval"); ok {
		convertedInt, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		notificationPolicy.RepeatInterval = convertedInt
	}
	if v, ok := d.GetOk("routes"); ok {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(v.(string)), &jsonMap)

		notificationPolicy.Routes = jsonMap
	}

	return notificationPolicy
}

func resourceNotificationPolicyCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*Client)

	notificationPolicy := convertResourceDataToNotificationPolicy(d)
	err := client.SetNotificationPolicyTree(notificationPolicy)
	if err != nil {
		return err
	}

	return resourceNotificationPolicyRead(d, m)
}

func resourceNotificationPolicyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNotificationPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	notificationPolicy := convertResourceDataToNotificationPolicy(d)
	err := client.SetNotificationPolicyTree(notificationPolicy)
	if err != nil {
		return err
	}

	return resourceNotificationPolicyRead(d, m)
}

func resourceNotificationPolicyDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*Client)

	err := client.ResetNotificationPolicyTree()
	if err != nil {
		return err
	}
	return nil
}
