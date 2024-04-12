/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type Pageable struct {
	PageNumber int32       `json:"pageNumber,omitempty"`
	PageSize   int32       `json:"pageSize,omitempty"`
	Offset     int64       `json:"offset,omitempty"`
	Sort       *SortObject `json:"sort,omitempty"`
	Unpaged    bool        `json:"unpaged,omitempty"`
	Paged      bool        `json:"paged,omitempty"`
}