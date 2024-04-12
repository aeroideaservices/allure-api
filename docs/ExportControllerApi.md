# {{classname}}

All URIs are relative to */api/rs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Generate1**](ExportControllerApi.md#Generate1) | **Post** /export/launch/pdf | Generate launch pdf report
[**GenerateTestCaseCsvExport**](ExportControllerApi.md#GenerateTestCaseCsvExport) | **Post** /export/testcase/csv | Generate test cases csv report
[**GenerateTestCasePdfExport**](ExportControllerApi.md#GenerateTestCasePdfExport) | **Post** /export/testcase/pdf | Generate test cases pdf report
[**GenerateTestResultCsvExport**](ExportControllerApi.md#GenerateTestResultCsvExport) | **Post** /export/testresult/csv | Generate test results csv report
[**GetSupportedLaunchPdfContent**](ExportControllerApi.md#GetSupportedLaunchPdfContent) | **Get** /export/launch/pdf/structure | Get supported launch pdf report parts
[**GetSupportedTCFields**](ExportControllerApi.md#GetSupportedTCFields) | **Get** /export/testcase/csv/mapping | Get supported test case export fields
[**GetSupportedTRFields**](ExportControllerApi.md#GetSupportedTRFields) | **Get** /export/testresult/csv/mapping | Get supported test result export fields

# **Generate1**
> ExportRequestDto Generate1(ctx, body, optional)
Generate launch pdf report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LaunchPdfOptions**](LaunchPdfOptions.md)|  | 
 **optional** | ***ExportControllerApiGenerate1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportControllerApiGenerate1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shared** | **optional.**|  | [default to true]

### Return type

[**ExportRequestDto**](ExportRequestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateTestCaseCsvExport**
> ExportRequestDto GenerateTestCaseCsvExport(ctx, body, optional)
Generate test cases csv report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCaseCsvExportOptions**](TestCaseCsvExportOptions.md)|  | 
 **optional** | ***ExportControllerApiGenerateTestCaseCsvExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportControllerApiGenerateTestCaseCsvExportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shared** | **optional.**|  | [default to true]

### Return type

[**ExportRequestDto**](ExportRequestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateTestCasePdfExport**
> ExportRequestDto GenerateTestCasePdfExport(ctx, body, optional)
Generate test cases pdf report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestCasePdfOptions**](TestCasePdfOptions.md)|  | 
 **optional** | ***ExportControllerApiGenerateTestCasePdfExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportControllerApiGenerateTestCasePdfExportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shared** | **optional.**|  | [default to true]

### Return type

[**ExportRequestDto**](ExportRequestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateTestResultCsvExport**
> ExportRequestDto GenerateTestResultCsvExport(ctx, body, optional)
Generate test results csv report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestResultCsvExportOptions**](TestResultCsvExportOptions.md)|  | 
 **optional** | ***ExportControllerApiGenerateTestResultCsvExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportControllerApiGenerateTestResultCsvExportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shared** | **optional.**|  | [default to true]

### Return type

[**ExportRequestDto**](ExportRequestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedLaunchPdfContent**
> LaunchPdfStructure GetSupportedLaunchPdfContent(ctx, )
Get supported launch pdf report parts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LaunchPdfStructure**](LaunchPdfStructure.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedTCFields**
> []TestCaseExportField GetSupportedTCFields(ctx, )
Get supported test case export fields

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TestCaseExportField**](TestCaseExportField.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedTRFields**
> []TestResultExportField GetSupportedTRFields(ctx, )
Get supported test result export fields

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TestResultExportField**](TestResultExportField.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

