/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type AnalyticAutomationTrendDto struct {
	Date                 int64 `json:"date,omitempty"`
	AutomatedCount       int64 `json:"automatedCount,omitempty"`
	ManualCount          int64 `json:"manualCount,omitempty"`
	SumDurationAutomated int64 `json:"sumDurationAutomated,omitempty"`
	SumDurationManual    int64 `json:"sumDurationManual,omitempty"`
}