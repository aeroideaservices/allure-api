/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestCaseAuditLogEntryDto struct {
	Id         int64                  `json:"id,omitempty"`
	Timestamp  int64                  `json:"timestamp,omitempty"`
	Username   string                 `json:"username,omitempty"`
	ActionType *AuditActionTypeDto    `json:"actionType,omitempty"`
	Data       []TestCaseAuditLogData `json:"data,omitempty"`
}