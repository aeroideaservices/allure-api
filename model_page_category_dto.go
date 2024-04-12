/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type PageCategoryDto struct {
	TotalElements    int64         `json:"totalElements,omitempty"`
	TotalPages       int32         `json:"totalPages,omitempty"`
	Pageable         *Pageable     `json:"pageable,omitempty"`
	Size             int32         `json:"size,omitempty"`
	Content          []CategoryDto `json:"content,omitempty"`
	Number           int32         `json:"number,omitempty"`
	Sort             *SortObject   `json:"sort,omitempty"`
	NumberOfElements int32         `json:"numberOfElements,omitempty"`
	First            bool          `json:"first,omitempty"`
	Last             bool          `json:"last,omitempty"`
	Empty            bool          `json:"empty,omitempty"`
}
