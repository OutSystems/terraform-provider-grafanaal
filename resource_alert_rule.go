package main

import (
	"encoding/json"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlertRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertRuleCreate,
		Read:   resourceAlertRuleRead,
		Update: resourceAlertRuleUpdate,
		Delete: resourceAlertRuleDelete,

		Schema: map[string]*schema.Schema{
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"condition": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_source_uid": {
							Type:     schema.TypeString,
							Required: true,
						},
						"model": {
							Type:     schema.TypeString,
							Required: true,
						},
						"query_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ref_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"relative_time_range": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: "it should contain for and to keys olny ex: {from:'test',to:'test'}",
						},
					},
				},
			},
			"exec_err_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"folder_uid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alert_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"no_data_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"for_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provenance": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func expandMap(mapInterface map[string]interface{}) (keyValuePair map[string]string) {
	keyValuePair = make(map[string]string)
	for key, val := range mapInterface {
		keyValuePair[key] = val.(string)
	}
	return
}

func convertResourceDataToAlertRule(d *schema.ResourceData) *AlertRule {
	alertRule := &AlertRule{}
	if v, ok := d.GetOk("annotations"); ok {
		alertRule.Annotations = expandMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("condition"); ok {
		alertRule.Condition = v.(string)
	}
	v := d.Get("query").([]interface{})

	var queryList []*AlertQuery
	for _, element := range v {
		queryData := &AlertQuery{}
		i := element.(map[string]interface{})
		queryData.DatasourceUID = i["data_source_uid"].(string)

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(i["model"].(string)), &jsonMap)

		queryData.Model = jsonMap

		queryData.QueryType = i["query_type"].(string)
		queryData.RefID = i["ref_id"].(string)
		fromTo := expandMap(i["relative_time_range"].(map[string]interface{}))

		relativeTimeRange := &RelativeTimeRange{}
		from, err := strconv.ParseInt(fromTo["from"], 10, 64)
		if err != nil {
			panic(err)
		}
		relativeTimeRange.From = from
		to, err := strconv.ParseInt(fromTo["to"], 10, 64)
		if err != nil {
			panic(err)
		}
		relativeTimeRange.To = to

		queryData.RelativeTimeRange = *relativeTimeRange
		queryList = append(queryList, queryData)
	}
	alertRule.Data = queryList

	if v, ok := d.GetOk("exec_err_state"); ok {
		alertRule.ExecErrState = ExecErrState(v.(string))
	}
	if v, ok := d.GetOk("folder_uid"); ok {
		alertRule.FolderUID = v.(string)
	}
	if v, ok := d.GetOk("alert_id"); ok {
		convertedInt, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		alertRule.ID = convertedInt
	}
	if v, ok := d.GetOk("labels"); ok {
		alertRule.Labels = expandMap(v.(map[string]interface{}))
	}
	if v, ok := d.GetOk("no_data_state"); ok {
		alertRule.NoDataState = NoDataState(v.(string))
	}
	if v, ok := d.GetOk("org_id"); ok {
		convertedInt, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		alertRule.OrgID = convertedInt
	}
	if v, ok := d.GetOk("rule_group"); ok {
		alertRule.RuleGroup = v.(string)
	}
	if v, ok := d.GetOk("title"); ok {
		alertRule.Title = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		alertRule.UID = v.(string)
	}
	if v, ok := d.GetOk("updated"); ok {
		alertRule.Updated = v.(string)
	}
	if v, ok := d.GetOk("for_time"); ok {
		convertedInt, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		alertRule.ForDuration = convertedInt
	}
	if v, ok := d.GetOk("provenance"); ok {
		alertRule.Provenance = v.(string)
	}

	return alertRule
}

func resourceAlertRuleCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	alertRule := convertResourceDataToAlertRule(d)
	id, err := client.NewAlertRule(alertRule)
	if err != nil {
		return err
	}
	d.SetId(id)

	return resourceAlertRuleRead(d, m)
}

func resourceAlertRuleRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	id := d.Id()

	alertRule, err := client.AlertRule(id)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(alertRule.UID)
	return nil
}

func resourceAlertRuleUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	alertRule := convertResourceDataToAlertRule(d)
	err := client.UpdateAlertRule(alertRule)
	if err != nil {
		return err
	}
	return resourceAlertRuleRead(d, m)
}

func resourceAlertRuleDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	id := d.Id()
	err := client.DeleteAlertRule(id)
	if err != nil {
		return err
	}
	return nil
}
