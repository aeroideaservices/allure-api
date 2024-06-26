/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestCaseRelationTypeDto string

// List of TestCaseRelationTypeDto
const (
	RELATED_TO_TestCaseRelationTypeDto       TestCaseRelationTypeDto = "related to"
	CLONES_TestCaseRelationTypeDto           TestCaseRelationTypeDto = "clones"
	IS_CLONED_BY_TestCaseRelationTypeDto     TestCaseRelationTypeDto = "is cloned by"
	DUPLICATES_TestCaseRelationTypeDto       TestCaseRelationTypeDto = "duplicates"
	IS_DUPLICATED_BY_TestCaseRelationTypeDto TestCaseRelationTypeDto = "is duplicated by"
	AUTOMATES_TestCaseRelationTypeDto        TestCaseRelationTypeDto = "automates"
	IS_AUTOMATED_BY_TestCaseRelationTypeDto  TestCaseRelationTypeDto = "is automated by"
)
