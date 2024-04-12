/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type LaunchCreateDto struct {
	ProjectId int64             `json:"projectId"`
	Name      string            `json:"name"`
	External  bool              `json:"external,omitempty"`
	Autoclose bool              `json:"autoclose,omitempty"`
	Tags      []LaunchTagDto    `json:"tags,omitempty"`
	Links     []ExternalLinkDto `json:"links,omitempty"`
	Issues    []IssueDto        `json:"issues,omitempty"`
}