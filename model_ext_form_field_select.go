/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type ExtFormFieldSelect struct {
	Name            string               `json:"name,omitempty"`
	Required        bool                 `json:"required,omitempty"`
	DependsOnFields []string             `json:"dependsOnFields,omitempty"`
	Type_           string               `json:"type"`
	LabelName       string               `json:"labelName,omitempty"`
	Options         []ExtFormFieldOption `json:"options,omitempty"`
	Multi           bool                 `json:"multi,omitempty"`
	Description     string               `json:"description,omitempty"`
	DefaultValue    string               `json:"defaultValue,omitempty"`
}
