/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestResultBulkResolveDto struct {
	Selection  *TestResultTreeSelectionDto `json:"selection"`
	Status     *TestStatus                 `json:"status"`
	Message    string                      `json:"message,omitempty"`
	Trace      string                      `json:"trace,omitempty"`
	CategoryId int64                       `json:"categoryId,omitempty"`
}