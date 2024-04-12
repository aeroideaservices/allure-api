# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfvAdd**](TestCaseBulkControllerApi.md#CfvAdd) | **Post** /testcase/bulk/cfv/add | Add custom field values for all test cases
[**CfvRemove**](TestCaseBulkControllerApi.md#CfvRemove) | **Post** /testcase/bulk/cfv/remove | Remove custom field values for all test cases
[**CloneAll**](TestCaseBulkControllerApi.md#CloneAll) | **Post** /testcase/bulk/clone | Clone test cases by ids
[**CreateTestPlan**](TestCaseBulkControllerApi.md#CreateTestPlan) | **Post** /testcase/bulk/testplan/create | Create test plan from selected test cases
[**DeleteAll**](TestCaseBulkControllerApi.md#DeleteAll) | **Post** /testcase/bulk/remove | Remove test cases by ids
[**ExternalLinkAdd**](TestCaseBulkControllerApi.md#ExternalLinkAdd) | **Post** /testcase/bulk/externallink/add | Add external link for all test cases
[**IssueAdd**](TestCaseBulkControllerApi.md#IssueAdd) | **Post** /testcase/bulk/issue/add | Add issues for all test cases
[**IssueRemove**](TestCaseBulkControllerApi.md#IssueRemove) | **Post** /testcase/bulk/issue/remove | Remove issues for all test cases
[**LayerSet**](TestCaseBulkControllerApi.md#LayerSet) | **Post** /testcase/bulk/layer/set | Set specified layer for all test cases
[**MemberAdd**](TestCaseBulkControllerApi.md#MemberAdd) | **Post** /testcase/bulk/member/add | Add members for all test cases
[**MemberRemove**](TestCaseBulkControllerApi.md#MemberRemove) | **Post** /testcase/bulk/member/remove | Remove member for all test cases
[**MoveAll**](TestCaseBulkControllerApi.md#MoveAll) | **Post** /testcase/bulk/move | Move test cases to other project
[**MuteAdd**](TestCaseBulkControllerApi.md#MuteAdd) | **Post** /testcase/bulk/mute/add | Add mute for all test cases
[**Run1**](TestCaseBulkControllerApi.md#Run1) | **Post** /testcase/bulk/run/new | Run selected test cases in a new launch
[**Run2**](TestCaseBulkControllerApi.md#Run2) | **Post** /testcase/bulk/run | Run selected test cases in a new launch
[**Run3**](TestCaseBulkControllerApi.md#Run3) | **Post** /testcase/bulk/run/existing | Run selected test cases in an existing launch
[**StatusSet**](TestCaseBulkControllerApi.md#StatusSet) | **Post** /testcase/bulk/status/set | Set specified status for all test cases
[**TagsAdd1**](TestCaseBulkControllerApi.md#TagsAdd1) | **Post** /testcase/bulk/tag/add | Add tags for all test cases
[**TagsRemove1**](TestCaseBulkControllerApi.md#TagsRemove1) | **Post** /testcase/bulk/tag/remove | Remove tags for all test cases

# **CfvAdd**
> CfvAdd(ctx, body)
Add custom field values for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkCfvDto**](TestCaseBulkCfvDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CfvRemove**
> CfvRemove(ctx, body)
Remove custom field values for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkEntityIdsDto**](TestCaseBulkEntityIdsDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloneAll**
> CloneAll(ctx, body)
Clone test cases by ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkCloneDto**](TestCaseBulkCloneDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTestPlan**
> TestPlanDto CreateTestPlan(ctx, body)
Create test plan from selected test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkTestPlanCreateDto**](TestCaseBulkTestPlanCreateDto.md)|  | 

### Return type

[**TestPlanDto**](TestPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAll**
> DeleteAll(ctx, body)
Remove test cases by ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkDto**](TestCaseBulkDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalLinkAdd**
> ExternalLinkAdd(ctx, body)
Add external link for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkExternalLinkDto**](TestCaseBulkExternalLinkDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueAdd**
> IssueAdd(ctx, body)
Add issues for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkIssueDto**](TestCaseBulkIssueDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueRemove**
> IssueRemove(ctx, body)
Remove issues for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkEntityIdsDto**](TestCaseBulkEntityIdsDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LayerSet**
> LayerSet(ctx, body)
Set specified layer for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkLayerDto**](TestCaseBulkLayerDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberAdd**
> MemberAdd(ctx, body)
Add members for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkMemberDto**](TestCaseBulkMemberDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberRemove**
> MemberRemove(ctx, body)
Remove member for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkEntityIdsDto**](TestCaseBulkEntityIdsDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveAll**
> MoveAll(ctx, body)
Move test cases to other project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkProjectChangeDto**](TestCaseBulkProjectChangeDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteAdd**
> MuteAdd(ctx, body)
Add mute for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkMuteDto**](TestCaseBulkMuteDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Run1**
> LaunchDto Run1(ctx, body)
Run selected test cases in a new launch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkRunNewLaunchDto**](TestCaseBulkRunNewLaunchDto.md)|  | 

### Return type

[**LaunchDto**](LaunchDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Run2**
> LaunchDto Run2(ctx, body)
Run selected test cases in a new launch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkRunNewLaunchDto**](TestCaseBulkRunNewLaunchDto.md)|  | 

### Return type

[**LaunchDto**](LaunchDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Run3**
> LaunchDto Run3(ctx, body)
Run selected test cases in an existing launch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkRunExistingLaunchDto**](TestCaseBulkRunExistingLaunchDto.md)|  | 

### Return type

[**LaunchDto**](LaunchDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatusSet**
> StatusSet(ctx, body)
Set specified status for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkStatusDto**](TestCaseBulkStatusDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagsAdd1**
> TagsAdd1(ctx, body)
Add tags for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkTagDto**](TestCaseBulkTagDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagsRemove1**
> TagsRemove1(ctx, body)
Remove tags for all test cases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseBulkEntityIdsDto**](TestCaseBulkEntityIdsDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

