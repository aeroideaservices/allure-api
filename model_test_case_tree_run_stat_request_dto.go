/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestCaseTreeRunStatRequestDto struct {
	Inverted      bool         `json:"inverted,omitempty"`
	GroupsInclude [][]int64    `json:"groupsInclude,omitempty"`
	GroupsExclude [][]int64    `json:"groupsExclude,omitempty"`
	LeafsInclude  []int64      `json:"leafsInclude,omitempty"`
	LeafsExclude  []int64      `json:"leafsExclude,omitempty"`
	Path          []int64      `json:"path,omitempty"`
	ProjectId     int64        `json:"projectId"`
	TreeId        int64        `json:"treeId,omitempty"`
	FilterId      int64        `json:"filterId,omitempty"`
	Search        string       `json:"search,omitempty"`
	JobsMapping   []JobMapping `json:"jobsMapping,omitempty"`
}
