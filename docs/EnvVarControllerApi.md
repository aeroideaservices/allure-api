# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create40**](EnvVarControllerApi.md#Create40) | **Post** /ev | Create a new environment variable
[**Create41**](EnvVarControllerApi.md#Create41) | **Post** /environment | Create a new environment variable
[**Delete36**](EnvVarControllerApi.md#Delete36) | **Delete** /environment/{id} | Delete environment variable by id
[**Delete37**](EnvVarControllerApi.md#Delete37) | **Delete** /ev/{id} | Delete environment variable by id
[**FindAll38**](EnvVarControllerApi.md#FindAll38) | **Get** /ev | Find all environment variables
[**FindAll39**](EnvVarControllerApi.md#FindAll39) | **Get** /environment | Find all environment variables
[**FindOne32**](EnvVarControllerApi.md#FindOne32) | **Get** /environment/{id} | Find environment variable by id
[**FindOne33**](EnvVarControllerApi.md#FindOne33) | **Get** /ev/{id} | Find environment variable by id
[**Merge1**](EnvVarControllerApi.md#Merge1) | **Post** /ev/merge | Merge environment variables
[**Merge2**](EnvVarControllerApi.md#Merge2) | **Post** /environment/merge | Merge environment variables
[**Patch35**](EnvVarControllerApi.md#Patch35) | **Patch** /environment/{id} | Patch environment variable
[**Patch36**](EnvVarControllerApi.md#Patch36) | **Patch** /ev/{id} | Patch environment variable
[**Suggest17**](EnvVarControllerApi.md#Suggest17) | **Get** /environment/suggest | Suggest environment variables
[**Suggest18**](EnvVarControllerApi.md#Suggest18) | **Get** /ev/suggest | Suggest environment variables

# **Create40**
> EnvVarDto Create40(ctx, body)
Create a new environment variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarCreateDto**](EnvVarCreateDto.md)|  | 

### Return type

[**EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create41**
> EnvVarDto Create41(ctx, body)
Create a new environment variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarCreateDto**](EnvVarCreateDto.md)|  | 

### Return type

[**EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete36**
> Delete36(ctx, id)
Delete environment variable by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete37**
> Delete37(ctx, id)
Delete environment variable by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll38**
> []EnvVarDto FindAll38(ctx, )
Find all environment variables

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAll39**
> []EnvVarDto FindAll39(ctx, )
Find all environment variables

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne32**
> EnvVarDto FindOne32(ctx, id)
Find environment variable by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOne33**
> EnvVarDto FindOne33(ctx, id)
Find environment variable by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Merge1**
> Merge1(ctx, body)
Merge environment variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarMergeDto**](EnvVarMergeDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Merge2**
> Merge2(ctx, body)
Merge environment variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarMergeDto**](EnvVarMergeDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch35**
> EnvVarDto Patch35(ctx, body, id)
Patch environment variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarPatchDto**](EnvVarPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Patch36**
> EnvVarDto Patch36(ctx, body, id)
Patch environment variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvVarPatchDto**](EnvVarPatchDto.md)|  | 
  **id** | **int64**|  | 

### Return type

[**EnvVarDto**](EnvVarDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest17**
> PageIdAndNameOnlyDto Suggest17(ctx, optional)
Suggest environment variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnvVarControllerApiSuggest17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvVarControllerApiSuggest17Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **id** | [**optional.Interface of []int64**](int64.md)|  | 
 **ignoreId** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageIdAndNameOnlyDto**](PageIdAndNameOnlyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest18**
> PageIdAndNameOnlyDto Suggest18(ctx, optional)
Suggest environment variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnvVarControllerApiSuggest18Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvVarControllerApiSuggest18Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **id** | [**optional.Interface of []int64**](int64.md)|  | 
 **ignoreId** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | [default to [&quot;name,ASC&quot;]]

### Return type

[**PageIdAndNameOnlyDto**](PageIdAndNameOnlyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

