/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestCaseScenarioStepDto struct {
	Name           string                     `json:"name,omitempty"`
	Keyword        string                     `json:"keyword,omitempty"`
	ExpectedResult string                     `json:"expectedResult,omitempty"`
	Attachments    []TestCaseAttachmentRowDto `json:"attachments,omitempty"`
	Steps          []TestCaseScenarioStepDto  `json:"steps,omitempty"`
	Leaf           bool                       `json:"leaf,omitempty"`
	StepsCount     int64                      `json:"stepsCount,omitempty"`
	HasContent     bool                       `json:"hasContent,omitempty"`
}
