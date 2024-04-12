/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type IntegrationCreateDto struct {
	Name                   string                  `json:"name,omitempty"`
	Type_                  *IntegrationTypeDto     `json:"type"`
	Settings               *IntegrationSettingsDto `json:"settings"`
	DefaultProjectSettings *IntegrationSettingsDto `json:"defaultProjectSettings,omitempty"`
	DefaultSecret          *IntegrationSettingsDto `json:"defaultSecret,omitempty"`
	EnabledByDefault       bool                    `json:"enabledByDefault,omitempty"`
}
