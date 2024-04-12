/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestResultRowDto struct {
	Id         int64       `json:"id,omitempty"`
	TestCaseId int64       `json:"testCaseId,omitempty"`
	Name       string      `json:"name,omitempty"`
	Status     *TestStatus `json:"status,omitempty"`
	Duration   int64       `json:"duration,omitempty"`
	Assignee   string      `json:"assignee,omitempty"`
}
