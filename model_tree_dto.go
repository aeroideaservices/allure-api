/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TreeDto struct {
	Id               int64            `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	ProjectId        int64            `json:"projectId,omitempty"`
	Fields           []CustomFieldDto `json:"fields,omitempty"`
	CreatedDate      int64            `json:"createdDate,omitempty"`
	LastModifiedDate int64            `json:"lastModifiedDate,omitempty"`
	CreatedBy        string           `json:"createdBy,omitempty"`
	LastModifiedBy   string           `json:"lastModifiedBy,omitempty"`
}
