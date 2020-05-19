package rendering

import (
	"regexp"

	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
)

func trimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(`\r?\n|\| `)
	return re.ReplaceAllString(s, " ")
}

// RenderMarkdown generate a Markdown file
func RenderMarkdown(ruleGroups []v1.RuleGroup) string {
	document := "<!-- Generated by promdoc. DO NOT EDIT. -->\n\n"
	document += "# Alerts\n"

	var currentGroup string
	for _, ruleGroup := range ruleGroups {
		if currentGroup != ruleGroup.Name {
			currentGroup = ruleGroup.Name
			document += "\n## " + ruleGroup.Name + "\n"
			document += "|Name|Summary|Message|Severity|Runbook\n"
			document += "|---|---|---|---|---|\n"
		}

		for _, rule := range ruleGroup.Rules {
			if rule.Alert == "" {
				continue
			}

			document += "|" + rule.Alert + "|" + rule.Annotations["summary"] + "|" + trimSpaceNewlineInString(
				rule.Annotations["message"]) + "|" + rule.Labels["severity"] + "|" + rule.Annotations["runbook_url"] + "|\n"
		}
	}

	return document
}

func RenderCSV(ruleGroups []v1.RuleGroup) string {
	var currentGroup string
	document := "Name,Summary,Message,Severity,Runbook\n"
	for _, ruleGroup := range ruleGroups {
		if currentGroup != ruleGroup.Name {
			currentGroup = ruleGroup.Name
			document += "\n## " + ruleGroup.Name + "\n"
		}

		for _, rule := range ruleGroup.Rules {
			if rule.Alert == "" {
				continue
			}

			document += rule.Alert + ";" + rule.Annotations["summary"] + ";" + trimSpaceNewlineInString(
				rule.Annotations["message"]) + ";" + rule.Labels["severity"] +
				";" + rule.Annotations["runbook_url"] + "\n"
		}
	}

	return document
}