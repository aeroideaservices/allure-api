/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type AnalyticTcStatusCountDto struct {
	StatusId    int64  `json:"statusId,omitempty"`
	StatusName  string `json:"statusName,omitempty"`
	StatusColor string `json:"statusColor,omitempty"`
	Count       int64  `json:"count,omitempty"`
}
