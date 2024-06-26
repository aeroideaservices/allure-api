/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestCaseDiff struct {
	ProjectId          *DiffValueChangeLong    `json:"projectId,omitempty"`
	Name               *DiffValueChangeString  `json:"name,omitempty"`
	Automated          *DiffValueChangeBoolean `json:"automated,omitempty"`
	Deleted            *DiffValueChangeBoolean `json:"deleted,omitempty"`
	Description        *DiffValueChangeString  `json:"description,omitempty"`
	DescriptionHtml    *DiffValueChangeString  `json:"descriptionHtml,omitempty"`
	Precondition       *DiffValueChangeString  `json:"precondition,omitempty"`
	PreconditionHtml   *DiffValueChangeString  `json:"preconditionHtml,omitempty"`
	ExpectedResult     *DiffValueChangeString  `json:"expectedResult,omitempty"`
	ExpectedResultHtml *DiffValueChangeString  `json:"expectedResultHtml,omitempty"`
	StatusId           *DiffValueChangeLong    `json:"statusId,omitempty"`
	WorkflowId         *DiffValueChangeLong    `json:"workflowId,omitempty"`
	TestLayerId        *DiffValueChangeLong    `json:"testLayerId,omitempty"`
}
