# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMembers**](TestResultMembersControllerApi.md#GetMembers) | **Get** /testresult/{testResultId}/members | Find user roles for test result
[**SetMembers**](TestResultMembersControllerApi.md#SetMembers) | **Post** /testresult/{testResultId}/members | Set user roles for test result

# **GetMembers**
> []MemberDto GetMembers(ctx, testResultId)
Find user roles for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 

### Return type

[**[]MemberDto**](MemberDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMembers**
> []MemberDto SetMembers(ctx, body, testResultId)
Set user roles for test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MemberDto**](MemberDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**[]MemberDto**](MemberDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

