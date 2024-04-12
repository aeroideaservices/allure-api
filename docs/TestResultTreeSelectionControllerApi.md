# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountLeaves**](TestResultTreeSelectionControllerApi.md#CountLeaves) | **Post** /testresulttree/select | Count test cases by tree select

# **CountLeaves**
> int64 CountLeaves(ctx, body)
Count test cases by tree select

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestResultTreeSelectionDto**](TestResultTreeSelectionDto.md)|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

