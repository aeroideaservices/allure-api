/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type AnalyticTrByStatusTrendDto struct {
	Date    int64 `json:"date,omitempty"`
	Failed  int64 `json:"failed,omitempty"`
	Broken  int64 `json:"broken,omitempty"`
	Passed  int64 `json:"passed,omitempty"`
	Skipped int64 `json:"skipped,omitempty"`
}
