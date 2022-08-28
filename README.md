Importing provider (version.tf file):

terraform {
  required_providers {
    grafanaal = {
      source = "OutSystems/grafanaal"
      version = "1.0.5"
    }
  }
}


Provider definition:

provider "grafanaal" {
  url = "grafana_url"
  token = "grafana_token"
}


ALERT RULE :

description of Alert Rule Schema : https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#span-idalert-rulespan-alertrule

Example:

resource "grafanaal_rule" "exaple_rule" {
    annotations = {
        "summary": "Example Alert"
    }
    condition = "C"
    query {
        data_source_uid="uid"
        model="{\"editorMode\":\"code\",\"exemplar\":false,\"expr\":\"sum(increase(outsystems_udp_delivery_messages_success_total{}[$__range]))by(environment)/\\nsum(increase(outsystems_udp_delivery_messages_total{}[$__range]))by(environment)*100\",\"format\":\"time_series\",\"hide\":false,\"instant\":false,\"interval\":\"\",\"intervalMs\":1000,\"legendFormat\":\"__auto\",\"maxDataPoints\":43200,\"range\":true,\"refId\":\"A\"}"
        query_type=""
        ref_id="A"
        relative_time_range={"from":"300","to":"0"}
    }
    query {
        data_source_uid="-100"
        model="{\"conditions\":[{\"evaluator\":{\"params\":[0,0],\"type\":\"gt\"},\"operator\":{\"type\":\"and\"},\"query\":{\"params\":[]},\"reducer\":{\"params\":[],\"type\":\"avg\"},\"type\":\"query\"}],\"datasource\":{\"name\":\"Expression\",\"type\":\"__expr__\",\"uid\":\"__expr__\"},\"expression\":\"A\",\"hide\":false,\"intervalMs\":1000,\"maxDataPoints\":43200,\"reducer\":\"mean\",\"refId\":\"B\",\"settings\":{\"mode\":\"replaceNN\",\"replaceWithValue\":0},\"type\":\"reduce\"}"
        query_type=""
        ref_id="B"
        relative_time_range={"from":"0","to":"0"}
    }
    query {
        data_source_uid="-100"
        model="{\"conditions\":[{\"evaluator\":{\"params\":[0,0],\"type\":\"gt\"},\"operator\":{\"type\":\"and\"},\"query\":{\"params\":[]},\"reducer\":{\"params\":[],\"type\":\"avg\"},\"type\":\"query\"}],\"datasource\":{\"name\":\"Expression\",\"type\":\"__expr__\",\"uid\":\"__expr__\"},\"expression\":\"$B<95\",\"hide\":false,\"intervalMs\":1000,\"maxDataPoints\":43200,\"refId\":\"C\",\"type\":\"math\"}"
        query_type=""
        ref_id="C"
        relative_time_range={"from":"0","to":"0"}
    }
    exec_err_state="Alerting"
    folder_uid="folder_uid"
    labels={
        "name": "udp_alert"
    }
    no_data_state="OK"
    org_id="1"
    rule_group="test"
    title="delivary rate test < 95%"
    for_time="300000000000"
}

Contact Point :

description of Contact Point Schema : https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#span-idembedded-contact-pointspan-embeddedcontactpoint

Example :

resource "grafanaal_contact_point" "example_contactpoint" {
  name = "test"
  uid="test_uid"
  type= "email"
  settings={"addresses":"vikram.patil@outsystems.com","singleEmail":true}
}

Notification Policy :

description of Contact Point Schema : https://grafana.com/docs/grafana/latest/developers/http_api/alerting_provisioning/#span-idroutespan-route



