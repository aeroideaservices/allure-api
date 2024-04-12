/*
 * allure-report-service
 *
 * Branch **HEAD**  Commit **1858ebdc07ef5d4e4615f9b1f2695e658df2e4ca**
 *
 * API version: 4.24.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package allure_api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the allure-report-service API v4.24.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccessGroupBulkControllerApi *AccessGroupBulkControllerApiService

	AccessGroupControllerApi *AccessGroupControllerApiService

	AnalyticControllerApi *AnalyticControllerApiService

	BlobStorageControllerApi *BlobStorageControllerApiService

	BlobStorageUpgradeControllerApi *BlobStorageUpgradeControllerApiService

	CategoryControllerApi *CategoryControllerApiService

	CategoryMatcherControllerApi *CategoryMatcherControllerApiService

	CleanerSchemaControllerApi *CleanerSchemaControllerApiService

	CleanupControllerApi *CleanupControllerApiService

	CommentControllerApi *CommentControllerApiService

	CustomFieldControllerApi *CustomFieldControllerApiService

	CustomFieldSchemaControllerApi *CustomFieldSchemaControllerApiService

	CustomFieldValueControllerApi *CustomFieldValueControllerApiService

	DashboardControllerApi *DashboardControllerApiService

	DefectBulkControllerApi *DefectBulkControllerApiService

	DefectControllerApi *DefectControllerApiService

	DefectMatcherControllerApi *DefectMatcherControllerApiService

	EnvVarControllerApi *EnvVarControllerApiService

	EnvVarSchemaControllerApi *EnvVarSchemaControllerApiService

	EnvVarValueControllerApi *EnvVarValueControllerApiService

	ExportControllerApi *ExportControllerApiService

	ExportRequestControllerApi *ExportRequestControllerApiService

	FilterControllerApi *FilterControllerApiService

	GlobalSettingsControllerApi *GlobalSettingsControllerApiService

	IdeTestCaseControllerApi *IdeTestCaseControllerApiService

	IframeControllerApi *IframeControllerApiService

	ImportRequestControllerApi *ImportRequestControllerApiService

	IntegrationControllerApi *IntegrationControllerApiService

	IntegrationExportControllerApi *IntegrationExportControllerApiService

	IntegrationIssueControllerApi *IntegrationIssueControllerApiService

	IntegrationSelectControllerApi *IntegrationSelectControllerApiService

	IntegrationTmsControllerApi *IntegrationTmsControllerApiService

	IntegrationWebhookControllerApi *IntegrationWebhookControllerApiService

	IssueControllerApi *IssueControllerApiService

	IssueSchemaControllerApi *IssueSchemaControllerApiService

	JobControllerApi *JobControllerApiService

	JobRunControllerApi *JobRunControllerApiService

	JobRunTestPlanControllerApi *JobRunTestPlanControllerApiService

	LaunchControllerApi *LaunchControllerApiService

	LaunchDiffControllerApi *LaunchDiffControllerApiService

	LaunchIssueControllerApi *LaunchIssueControllerApiService

	LaunchSearchControllerApi *LaunchSearchControllerApiService

	LaunchTagControllerApi *LaunchTagControllerApiService

	LaunchUploadControllerApi *LaunchUploadControllerApiService

	MarkdownPreviewControllerApi *MarkdownPreviewControllerApiService

	MemberControllerApi *MemberControllerApiService

	MuteControllerApi *MuteControllerApiService

	PermissionControllerApi *PermissionControllerApiService

	PermissionSetControllerApi *PermissionSetControllerApiService

	ProjectAccessControllerApi *ProjectAccessControllerApiService

	ProjectCategoryControllerApi *ProjectCategoryControllerApiService

	ProjectCategoryMatcherControllerApi *ProjectCategoryMatcherControllerApiService

	ProjectCollaboratorControllerApi *ProjectCollaboratorControllerApiService

	ProjectControllerApi *ProjectControllerApiService

	ProjectPropertyControllerApi *ProjectPropertyControllerApiService

	ProjectSettingsControllerApi *ProjectSettingsControllerApiService

	RoleControllerApi *RoleControllerApiService

	RoleSchemaControllerApi *RoleSchemaControllerApiService

	SharedStepAttachmentControllerApi *SharedStepAttachmentControllerApiService

	SharedStepControllerApi *SharedStepControllerApiService

	SharedStepScenarioControllerApi *SharedStepScenarioControllerApiService

	StatusControllerApi *StatusControllerApiService

	TestCaseAttachmentControllerApi *TestCaseAttachmentControllerApiService

	TestCaseAuditControllerApi *TestCaseAuditControllerApiService

	TestCaseBulkControllerApi *TestCaseBulkControllerApiService

	TestCaseCloneControllerApi *TestCaseCloneControllerApiService

	TestCaseControllerApi *TestCaseControllerApiService

	TestCaseCsvImportControllerApi *TestCaseCsvImportControllerApiService

	TestCaseCustomFieldControllerApi *TestCaseCustomFieldControllerApiService

	TestCaseDefectControllerApi *TestCaseDefectControllerApiService

	TestCaseExampleControllerApi *TestCaseExampleControllerApiService

	TestCaseExportControllerApi *TestCaseExportControllerApiService

	TestCaseIssueControllerApi *TestCaseIssueControllerApiService

	TestCaseMembersControllerApi *TestCaseMembersControllerApiService

	TestCaseOverviewControllerApi *TestCaseOverviewControllerApiService

	TestCaseRelationControllerApi *TestCaseRelationControllerApiService

	TestCaseScenarioControllerApi *TestCaseScenarioControllerApiService

	TestCaseSearchControllerApi *TestCaseSearchControllerApiService

	TestCaseSyncControllerApi *TestCaseSyncControllerApiService

	TestCaseTagControllerApi *TestCaseTagControllerApiService

	TestCaseTestKeyControllerApi *TestCaseTestKeyControllerApiService

	TestCaseTreeControllerApi *TestCaseTreeControllerApiService

	TestCaseTreeSelectionControllerApi *TestCaseTreeSelectionControllerApiService

	TestCaseUpdateSchemaControllerApi *TestCaseUpdateSchemaControllerApiService

	TestFixtureResultAttachmentControllerApi *TestFixtureResultAttachmentControllerApiService

	TestKeyControllerApi *TestKeyControllerApiService

	TestKeySchemaControllerApi *TestKeySchemaControllerApiService

	TestLayerControllerApi *TestLayerControllerApiService

	TestLayerSchemaControllerApi *TestLayerSchemaControllerApiService

	TestPlanControllerApi *TestPlanControllerApiService

	TestResultAttachmentControllerApi *TestResultAttachmentControllerApiService

	TestResultBulkControllerApi *TestResultBulkControllerApiService

	TestResultControllerApi *TestResultControllerApiService

	TestResultCustomFieldControllerApi *TestResultCustomFieldControllerApiService

	TestResultDefectControllerApi *TestResultDefectControllerApiService

	TestResultEnvVarControllerApi *TestResultEnvVarControllerApiService

	TestResultFixtureControllerApi *TestResultFixtureControllerApiService

	TestResultIssueControllerApi *TestResultIssueControllerApiService

	TestResultMembersControllerApi *TestResultMembersControllerApiService

	TestResultMuteControllerApi *TestResultMuteControllerApiService

	TestResultRerunControllerApi *TestResultRerunControllerApiService

	TestResultRunControllerApi *TestResultRunControllerApiService

	TestResultSearchControllerApi *TestResultSearchControllerApiService

	TestResultTestKeyControllerApi *TestResultTestKeyControllerApiService

	TestResultTreeControllerApi *TestResultTreeControllerApiService

	TestResultTreeSelectionControllerApi *TestResultTreeSelectionControllerApiService

	TestTagControllerApi *TestTagControllerApiService

	TreeControllerApi *TreeControllerApiService

	TreePathControllerApi *TreePathControllerApiService

	UploadControllerApi *UploadControllerApiService

	WidgetControllerApi *WidgetControllerApiService

	WorkflowControllerApi *WorkflowControllerApiService

	WorkflowSchemaControllerApi *WorkflowSchemaControllerApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccessGroupBulkControllerApi = (*AccessGroupBulkControllerApiService)(&c.common)
	c.AccessGroupControllerApi = (*AccessGroupControllerApiService)(&c.common)
	c.AnalyticControllerApi = (*AnalyticControllerApiService)(&c.common)
	c.BlobStorageControllerApi = (*BlobStorageControllerApiService)(&c.common)
	c.BlobStorageUpgradeControllerApi = (*BlobStorageUpgradeControllerApiService)(&c.common)
	c.CategoryControllerApi = (*CategoryControllerApiService)(&c.common)
	c.CategoryMatcherControllerApi = (*CategoryMatcherControllerApiService)(&c.common)
	c.CleanerSchemaControllerApi = (*CleanerSchemaControllerApiService)(&c.common)
	c.CleanupControllerApi = (*CleanupControllerApiService)(&c.common)
	c.CommentControllerApi = (*CommentControllerApiService)(&c.common)
	c.CustomFieldControllerApi = (*CustomFieldControllerApiService)(&c.common)
	c.CustomFieldSchemaControllerApi = (*CustomFieldSchemaControllerApiService)(&c.common)
	c.CustomFieldValueControllerApi = (*CustomFieldValueControllerApiService)(&c.common)
	c.DashboardControllerApi = (*DashboardControllerApiService)(&c.common)
	c.DefectBulkControllerApi = (*DefectBulkControllerApiService)(&c.common)
	c.DefectControllerApi = (*DefectControllerApiService)(&c.common)
	c.DefectMatcherControllerApi = (*DefectMatcherControllerApiService)(&c.common)
	c.EnvVarControllerApi = (*EnvVarControllerApiService)(&c.common)
	c.EnvVarSchemaControllerApi = (*EnvVarSchemaControllerApiService)(&c.common)
	c.EnvVarValueControllerApi = (*EnvVarValueControllerApiService)(&c.common)
	c.ExportControllerApi = (*ExportControllerApiService)(&c.common)
	c.ExportRequestControllerApi = (*ExportRequestControllerApiService)(&c.common)
	c.FilterControllerApi = (*FilterControllerApiService)(&c.common)
	c.GlobalSettingsControllerApi = (*GlobalSettingsControllerApiService)(&c.common)
	c.IdeTestCaseControllerApi = (*IdeTestCaseControllerApiService)(&c.common)
	c.IframeControllerApi = (*IframeControllerApiService)(&c.common)
	c.ImportRequestControllerApi = (*ImportRequestControllerApiService)(&c.common)
	c.IntegrationControllerApi = (*IntegrationControllerApiService)(&c.common)
	c.IntegrationExportControllerApi = (*IntegrationExportControllerApiService)(&c.common)
	c.IntegrationIssueControllerApi = (*IntegrationIssueControllerApiService)(&c.common)
	c.IntegrationSelectControllerApi = (*IntegrationSelectControllerApiService)(&c.common)
	c.IntegrationTmsControllerApi = (*IntegrationTmsControllerApiService)(&c.common)
	c.IntegrationWebhookControllerApi = (*IntegrationWebhookControllerApiService)(&c.common)
	c.IssueControllerApi = (*IssueControllerApiService)(&c.common)
	c.IssueSchemaControllerApi = (*IssueSchemaControllerApiService)(&c.common)
	c.JobControllerApi = (*JobControllerApiService)(&c.common)
	c.JobRunControllerApi = (*JobRunControllerApiService)(&c.common)
	c.JobRunTestPlanControllerApi = (*JobRunTestPlanControllerApiService)(&c.common)
	c.LaunchControllerApi = (*LaunchControllerApiService)(&c.common)
	c.LaunchDiffControllerApi = (*LaunchDiffControllerApiService)(&c.common)
	c.LaunchIssueControllerApi = (*LaunchIssueControllerApiService)(&c.common)
	c.LaunchSearchControllerApi = (*LaunchSearchControllerApiService)(&c.common)
	c.LaunchTagControllerApi = (*LaunchTagControllerApiService)(&c.common)
	c.LaunchUploadControllerApi = (*LaunchUploadControllerApiService)(&c.common)
	c.MarkdownPreviewControllerApi = (*MarkdownPreviewControllerApiService)(&c.common)
	c.MemberControllerApi = (*MemberControllerApiService)(&c.common)
	c.MuteControllerApi = (*MuteControllerApiService)(&c.common)
	c.PermissionControllerApi = (*PermissionControllerApiService)(&c.common)
	c.PermissionSetControllerApi = (*PermissionSetControllerApiService)(&c.common)
	c.ProjectAccessControllerApi = (*ProjectAccessControllerApiService)(&c.common)
	c.ProjectCategoryControllerApi = (*ProjectCategoryControllerApiService)(&c.common)
	c.ProjectCategoryMatcherControllerApi = (*ProjectCategoryMatcherControllerApiService)(&c.common)
	c.ProjectCollaboratorControllerApi = (*ProjectCollaboratorControllerApiService)(&c.common)
	c.ProjectControllerApi = (*ProjectControllerApiService)(&c.common)
	c.ProjectPropertyControllerApi = (*ProjectPropertyControllerApiService)(&c.common)
	c.ProjectSettingsControllerApi = (*ProjectSettingsControllerApiService)(&c.common)
	c.RoleControllerApi = (*RoleControllerApiService)(&c.common)
	c.RoleSchemaControllerApi = (*RoleSchemaControllerApiService)(&c.common)
	c.SharedStepAttachmentControllerApi = (*SharedStepAttachmentControllerApiService)(&c.common)
	c.SharedStepControllerApi = (*SharedStepControllerApiService)(&c.common)
	c.SharedStepScenarioControllerApi = (*SharedStepScenarioControllerApiService)(&c.common)
	c.StatusControllerApi = (*StatusControllerApiService)(&c.common)
	c.TestCaseAttachmentControllerApi = (*TestCaseAttachmentControllerApiService)(&c.common)
	c.TestCaseAuditControllerApi = (*TestCaseAuditControllerApiService)(&c.common)
	c.TestCaseBulkControllerApi = (*TestCaseBulkControllerApiService)(&c.common)
	c.TestCaseCloneControllerApi = (*TestCaseCloneControllerApiService)(&c.common)
	c.TestCaseControllerApi = (*TestCaseControllerApiService)(&c.common)
	c.TestCaseCsvImportControllerApi = (*TestCaseCsvImportControllerApiService)(&c.common)
	c.TestCaseCustomFieldControllerApi = (*TestCaseCustomFieldControllerApiService)(&c.common)
	c.TestCaseDefectControllerApi = (*TestCaseDefectControllerApiService)(&c.common)
	c.TestCaseExampleControllerApi = (*TestCaseExampleControllerApiService)(&c.common)
	c.TestCaseExportControllerApi = (*TestCaseExportControllerApiService)(&c.common)
	c.TestCaseIssueControllerApi = (*TestCaseIssueControllerApiService)(&c.common)
	c.TestCaseMembersControllerApi = (*TestCaseMembersControllerApiService)(&c.common)
	c.TestCaseOverviewControllerApi = (*TestCaseOverviewControllerApiService)(&c.common)
	c.TestCaseRelationControllerApi = (*TestCaseRelationControllerApiService)(&c.common)
	c.TestCaseScenarioControllerApi = (*TestCaseScenarioControllerApiService)(&c.common)
	c.TestCaseSearchControllerApi = (*TestCaseSearchControllerApiService)(&c.common)
	c.TestCaseSyncControllerApi = (*TestCaseSyncControllerApiService)(&c.common)
	c.TestCaseTagControllerApi = (*TestCaseTagControllerApiService)(&c.common)
	c.TestCaseTestKeyControllerApi = (*TestCaseTestKeyControllerApiService)(&c.common)
	c.TestCaseTreeControllerApi = (*TestCaseTreeControllerApiService)(&c.common)
	c.TestCaseTreeSelectionControllerApi = (*TestCaseTreeSelectionControllerApiService)(&c.common)
	c.TestCaseUpdateSchemaControllerApi = (*TestCaseUpdateSchemaControllerApiService)(&c.common)
	c.TestFixtureResultAttachmentControllerApi = (*TestFixtureResultAttachmentControllerApiService)(&c.common)
	c.TestKeyControllerApi = (*TestKeyControllerApiService)(&c.common)
	c.TestKeySchemaControllerApi = (*TestKeySchemaControllerApiService)(&c.common)
	c.TestLayerControllerApi = (*TestLayerControllerApiService)(&c.common)
	c.TestLayerSchemaControllerApi = (*TestLayerSchemaControllerApiService)(&c.common)
	c.TestPlanControllerApi = (*TestPlanControllerApiService)(&c.common)
	c.TestResultAttachmentControllerApi = (*TestResultAttachmentControllerApiService)(&c.common)
	c.TestResultBulkControllerApi = (*TestResultBulkControllerApiService)(&c.common)
	c.TestResultControllerApi = (*TestResultControllerApiService)(&c.common)
	c.TestResultCustomFieldControllerApi = (*TestResultCustomFieldControllerApiService)(&c.common)
	c.TestResultDefectControllerApi = (*TestResultDefectControllerApiService)(&c.common)
	c.TestResultEnvVarControllerApi = (*TestResultEnvVarControllerApiService)(&c.common)
	c.TestResultFixtureControllerApi = (*TestResultFixtureControllerApiService)(&c.common)
	c.TestResultIssueControllerApi = (*TestResultIssueControllerApiService)(&c.common)
	c.TestResultMembersControllerApi = (*TestResultMembersControllerApiService)(&c.common)
	c.TestResultMuteControllerApi = (*TestResultMuteControllerApiService)(&c.common)
	c.TestResultRerunControllerApi = (*TestResultRerunControllerApiService)(&c.common)
	c.TestResultRunControllerApi = (*TestResultRunControllerApiService)(&c.common)
	c.TestResultSearchControllerApi = (*TestResultSearchControllerApiService)(&c.common)
	c.TestResultTestKeyControllerApi = (*TestResultTestKeyControllerApiService)(&c.common)
	c.TestResultTreeControllerApi = (*TestResultTreeControllerApiService)(&c.common)
	c.TestResultTreeSelectionControllerApi = (*TestResultTreeSelectionControllerApiService)(&c.common)
	c.TestTagControllerApi = (*TestTagControllerApiService)(&c.common)
	c.TreeControllerApi = (*TreeControllerApiService)(&c.common)
	c.TreePathControllerApi = (*TreePathControllerApiService)(&c.common)
	c.UploadControllerApi = (*UploadControllerApiService)(&c.common)
	c.WidgetControllerApi = (*WidgetControllerApiService)(&c.common)
	c.WorkflowControllerApi = (*WorkflowControllerApiService)(&c.common)
	c.WorkflowSchemaControllerApi = (*WorkflowSchemaControllerApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}