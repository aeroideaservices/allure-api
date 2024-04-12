# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMembers2**](TestCaseMembersControllerApi.md#GetMembers2) | **Get** /testcase/{testCaseId}/members | Find user roles for test case
[**SetMembers1**](TestCaseMembersControllerApi.md#SetMembers1) | **Post** /testcase/{testCaseId}/members | Set user roles for test case

# **GetMembers2**
> []MemberDto GetMembers2(ctx, testCaseId)
Find user roles for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testCaseId** | **int64**|  | 

### Return type

[**[]MemberDto**](MemberDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMembers1**
> []MemberDto SetMembers1(ctx, body, testCaseId)
Set user roles for test case

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MemberDto**](MemberDto.md)|  | 
  **testCaseId** | **int64**|  | 

### Return type

[**[]MemberDto**](MemberDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

