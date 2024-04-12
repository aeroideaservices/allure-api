# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Copy**](SharedStepScenarioControllerApi.md#Copy) | **Post** /sharedstep/step/{id}/copy | Copy scenario step
[**Create19**](SharedStepScenarioControllerApi.md#Create19) | **Post** /sharedstep/step | Create scenario step
[**DeleteById2**](SharedStepScenarioControllerApi.md#DeleteById2) | **Delete** /sharedstep/step/{id} | Delete a specified scenario step
[**FindOne15**](SharedStepScenarioControllerApi.md#FindOne15) | **Get** /sharedstep/{id}/step | Get scenario for shared step
[**Move2**](SharedStepScenarioControllerApi.md#Move2) | **Post** /sharedstep/step/{id}/move | Move scenario step
[**PatchById1**](SharedStepScenarioControllerApi.md#PatchById1) | **Patch** /sharedstep/step/{id} | Patch a specified scenario step
[**Usage**](SharedStepScenarioControllerApi.md#Usage) | **Get** /sharedstep/{id}/usage | Get testcases with usage of shared step

# **Copy**
> NormalizedScenarioDto Copy(ctx, body, id)
Copy scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepCopyDto**](ScenarioStepCopyDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create19**
> ScenarioStepCreatedResponseDto Create19(ctx, body, optional)
Create scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepCreateDto**](ScenarioStepCreateDto.md)|  | 
 **optional** | ***SharedStepScenarioControllerApiCreate19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharedStepScenarioControllerApiCreate19Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beforeId** | **optional.**|  | 
 **afterId** | **optional.**|  | 

### Return type

[**ScenarioStepCreatedResponseDto**](ScenarioStepCreatedResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteById2**
> NormalizedScenarioDto DeleteById2(ctx, id)
Delete a specified scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne15**
> NormalizedScenarioDto FindOne15(ctx, id)
Get scenario for shared step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Move2**
> NormalizedScenarioDto Move2(ctx, body, id)
Move scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepMoveDto**](ScenarioStepMoveDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchById1**
> NormalizedScenarioDto PatchById1(ctx, body, id)
Patch a specified scenario step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScenarioStepPatchDto**](ScenarioStepPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**NormalizedScenarioDto**](NormalizedScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Usage**
> PageTestCaseRowDto Usage(ctx, id, optional)
Get testcases with usage of shared step

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***SharedStepScenarioControllerApiUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharedStepScenarioControllerApiUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;createdDate,DESC&quot;]]

### Return type

[**PageTestCaseRowDto**](PageTestCaseRowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

