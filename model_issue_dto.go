/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type IssueDto struct {
	Id            int64  `json:"id,omitempty"`
	IntegrationId int64  `json:"integrationId,omitempty"`
	Name          string `json:"name,omitempty"`
	Url           string `json:"url,omitempty"`
	Summary       string `json:"summary,omitempty"`
	Status        string `json:"status,omitempty"`
	Closed        bool   `json:"closed,omitempty"`
}
