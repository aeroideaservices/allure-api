/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type ImportRequestDto struct {
	Id               int64                  `json:"id,omitempty"`
	ProjectId        int64                  `json:"projectId,omitempty"`
	Type_            *ImportRequestTypeDto  `json:"type,omitempty"`
	State            *ImportRequestStateDto `json:"state,omitempty"`
	StorageKey       string                 `json:"storageKey,omitempty"`
	OriginalFileName string                 `json:"originalFileName,omitempty"`
	ContentType      string                 `json:"contentType,omitempty"`
	ContentLength    int64                  `json:"contentLength,omitempty"`
	ErrorMessage     string                 `json:"errorMessage,omitempty"`
	CountImported    int64                  `json:"countImported,omitempty"`
	CountFailed      int64                  `json:"countFailed,omitempty"`
	CreatedDate      int64                  `json:"createdDate,omitempty"`
	LastModifiedDate int64                  `json:"lastModifiedDate,omitempty"`
	CreatedBy        string                 `json:"createdBy,omitempty"`
	LastModifiedBy   string                 `json:"lastModifiedBy,omitempty"`
}
