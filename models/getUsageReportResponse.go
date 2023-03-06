package models

type GetUsageReportResponse struct {
    Url              *string `json:"url"`
    UsageReportUrl   *string `json:"usage_report_url"`
    GroupedReportUrl *string `json:"grouped_report_url"`
}
