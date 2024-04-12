/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

type TestCaseBulkCloneDto struct {
	Selection         *TestCaseTreeSelectionDto `json:"selection"`
	ToProjectId       int64                     `json:"toProjectId,omitempty"`
	StatusId          int64                     `json:"statusId,omitempty"`
	WorkflowId        int64                     `json:"workflowId,omitempty"`
	NameSuffix        string                    `json:"nameSuffix,omitempty"`
	IgnoreAttachments bool                      `json:"ignoreAttachments,omitempty"`
	IgnoreCfv         bool                      `json:"ignoreCfv,omitempty"`
	IgnoreIssueLinks  bool                      `json:"ignoreIssueLinks,omitempty"`
	IgnoreLinks       bool                      `json:"ignoreLinks,omitempty"`
	IgnoreMembers     bool                      `json:"ignoreMembers,omitempty"`
	IgnoreParameters  bool                      `json:"ignoreParameters,omitempty"`
	IgnoreRelations   bool                      `json:"ignoreRelations,omitempty"`
	IgnoreScenario    bool                      `json:"ignoreScenario,omitempty"`
	IgnoreTags        bool                      `json:"ignoreTags,omitempty"`
	IgnoreTestKeys    bool                      `json:"ignoreTestKeys,omitempty"`
}