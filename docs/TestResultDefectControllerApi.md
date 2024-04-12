# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefect**](TestResultDefectControllerApi.md#CreateDefect) | **Post** /testresult/{testResultId}/defect | 
[**GetCandidates**](TestResultDefectControllerApi.md#GetCandidates) | **Get** /testresult/{id}/defect/candidate | 
[**GetDefects**](TestResultDefectControllerApi.md#GetDefects) | **Get** /testresult/{testResultId}/defect | 
[**GetMatches**](TestResultDefectControllerApi.md#GetMatches) | **Get** /testresult/defect/match | 
[**GetSimilar**](TestResultDefectControllerApi.md#GetSimilar) | **Get** /testresult/{id}/defect/similar | 
[**Link**](TestResultDefectControllerApi.md#Link) | **Post** /testresult/{testResultId}/defect/{defectId} | 
[**Unlink**](TestResultDefectControllerApi.md#Unlink) | **Delete** /testresult/{testResultId}/defect/{defectId} | 

# **CreateDefect**
> DefectDto CreateDefect(ctx, body, testResultId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestResultNewDefectDto**](TestResultNewDefectDto.md)|  | 
  **testResultId** | **int64**|  | 

### Return type

[**DefectDto**](DefectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCandidates**
> PageDefectDto GetCandidates(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestResultDefectControllerApiGetCandidatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultDefectControllerApiGetCandidatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threshold** | **optional.Float64**|  | [default to 0.7]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageDefectDto**](PageDefectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefects**
> PageDefectRowDto GetDefects(ctx, testResultId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 
 **optional** | ***TestResultDefectControllerApiGetDefectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultDefectControllerApiGetDefectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageDefectRowDto**](PageDefectRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatches**
> PageTestResultDefectMatchDto GetMatches(ctx, launchId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **launchId** | **int64**|  | 
 **optional** | ***TestResultDefectControllerApiGetMatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultDefectControllerApiGetMatchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageRegex** | **optional.String**|  | 
 **traceRegex** | **optional.String**|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageTestResultDefectMatchDto**](PageTestResultDefectMatchDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSimilar**
> PageTestResultRowDto GetSimilar(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***TestResultDefectControllerApiGetSimilarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultDefectControllerApiGetSimilarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threshold** | **optional.Float64**|  | [default to 0.7]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PageTestResultRowDto**](PageTestResultRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Link**
> DefectDto Link(ctx, testResultId, defectId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 
  **defectId** | **int64**|  | 
 **optional** | ***TestResultDefectControllerApiLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestResultDefectControllerApiLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TestResultLinkDefectDto**](TestResultLinkDefectDto.md)|  | 

### Return type

[**DefectDto**](DefectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unlink**
> Unlink(ctx, testResultId, defectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testResultId** | **int64**|  | 
  **defectId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

