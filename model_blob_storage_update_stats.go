/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type BlobStorageUpdateStats struct {
	TestResultCount                  int32 `json:"testResultCount,omitempty"`
	TestResultAttachmentCount        int32 `json:"testResultAttachmentCount,omitempty"`
	TestFixtureResultCount           int32 `json:"testFixtureResultCount,omitempty"`
	TestFixtureResultAttachmentCount int32 `json:"testFixtureResultAttachmentCount,omitempty"`
	TestCaseAttachmentCount          int32 `json:"testCaseAttachmentCount,omitempty"`
}
