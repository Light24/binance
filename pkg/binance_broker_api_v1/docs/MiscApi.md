# {{classname}}

All URIs are relative to *http://example.com/sapi/v1/broker*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrokerAccountInformation**](MiscApi.md#BrokerAccountInformation) | **Get** /info | Broker Account Information
[**ChangeSubAccountCommission**](MiscApi.md#ChangeSubAccountCommission) | **Post** /subAccountApi/commission | Change Sub Account Commission
[**ChangeSubAccountFuturesCommissionAdjustment**](MiscApi.md#ChangeSubAccountFuturesCommissionAdjustment) | **Post** /subAccountApi/commission/futures | Change Sub Account Futures Commission Adjustment
[**CreateApiKeyforSubAccount**](MiscApi.md#CreateApiKeyforSubAccount) | **Post** /subAccountApi | Create Api Key for Sub Account
[**CreateaSubAccount**](MiscApi.md#CreateaSubAccount) | **Post** /subAccount | Create a Sub Account
[**DeleteSubAccountApiKey**](MiscApi.md#DeleteSubAccountApiKey) | **Delete** /subAccountApi | Delete Sub Account Api Key
[**EnableFuturesforSubAccount**](MiscApi.md#EnableFuturesforSubAccount) | **Post** /subAccount/futures | Enable Futures for Sub Account
[**EnableMarginforSubAccount**](MiscApi.md#EnableMarginforSubAccount) | **Post** /subAccount/margin | Enable Margin for Sub Account
[**EnableOrDisableBNBBurnforSubAccountMarginInterest**](MiscApi.md#EnableOrDisableBNBBurnforSubAccountMarginInterest) | **Post** /subAccount/bnbBurn/marginInterest | Enable Or Disable BNB Burn for Sub Account Margin Interest
[**EnableOrDisableBNBBurnforSubAccountSPOTandMARGIN**](MiscApi.md#EnableOrDisableBNBBurnforSubAccountSPOTandMARGIN) | **Post** /subAccount/bnbBurn/spot | Enable Or Disable BNB Burn for Sub Account SPOT and MARGIN
[**GenerateBrokerCommissionRebateHistory**](MiscApi.md#GenerateBrokerCommissionRebateHistory) | **Post** /rebate/historicalRecord | Generate Broker Commission Rebate History
[**GetBNBBurnStatusforSubAccount**](MiscApi.md#GetBNBBurnStatusforSubAccount) | **Get** /subAccount/bnbBurn/status | Get BNB Burn Status for Sub Account
[**GetSubAccountDepositHistory**](MiscApi.md#GetSubAccountDepositHistory) | **Get** /subAccount/depositHist | Get Sub Account Deposit History
[**QueryBrokerCommissionRebateHistory**](MiscApi.md#QueryBrokerCommissionRebateHistory) | **Get** /rebate/historicalRecord | Query Broker Commission Rebate History
[**QueryBrokerCommissionRebateRecentRecord**](MiscApi.md#QueryBrokerCommissionRebateRecentRecord) | **Get** /rebate/recentRecord | Query Broker Commission Rebate Recent Record
[**QuerySubAccount**](MiscApi.md#QuerySubAccount) | **Get** /subAccount | Query Sub Account
[**QuerySubAccountApiKey**](MiscApi.md#QuerySubAccountApiKey) | **Get** /subAccountApi | Query Sub Account Api Key
[**QuerySubAccountFuturesCommissionAdjustment**](MiscApi.md#QuerySubAccountFuturesCommissionAdjustment) | **Get** /subAccountApi/commission/futures | Query Sub Account Futures Commission Adjustment
[**QuerySubAccountSpotAssetinfo**](MiscApi.md#QuerySubAccountSpotAssetinfo) | **Get** /subAccount/spotSummary | Query Sub Account Spot Asset info
[**QuerySubAccountTransferHistory**](MiscApi.md#QuerySubAccountTransferHistory) | **Get** /transfer | Query Sub Account Transfer History
[**QuerySubaccountFuturesAssetinfo**](MiscApi.md#QuerySubaccountFuturesAssetinfo) | **Get** /subAccount/futuresSummary | Query Subaccount Futures Asset info
[**QuerySubaccountMarginAssetinfo**](MiscApi.md#QuerySubaccountMarginAssetinfo) | **Get** /subAccount/marginSummary | Query Subaccount Margin Asset info
[**SubAccountTransfer**](MiscApi.md#SubAccountTransfer) | **Post** /transfer | Sub Account Transfer

# **BrokerAccountInformation**
> BrokerAccountInformation(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Broker Account Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeSubAccountCommission**
> ChangeSubAccountCommission(ctx, subAccountId, makerCommission, takerCommission, timestamp, signature, contentType, xMBXAPIKEY)
Change Sub Account Commission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **makerCommission** | **string**|  | 
  **takerCommission** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeSubAccountFuturesCommissionAdjustment**
> ChangeSubAccountFuturesCommissionAdjustment(ctx, subAccountId, symbol, makerAdjustment, takerAdjustment, timestamp, signature, contentType, xMBXAPIKEY)
Change Sub Account Futures Commission Adjustment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **symbol** | **string**|  | 
  **makerAdjustment** | **string**|  | 
  **takerAdjustment** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApiKeyforSubAccount**
> CreateApiKeyforSubAccount(ctx, subAccountId, canTrade, timestamp, signature, contentType, xMBXAPIKEY)
Create Api Key for Sub Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **canTrade** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateaSubAccount**
> CreateaSubAccount(ctx, tag, timestamp, signature, contentType, xMBXAPIKEY)
Create a Sub Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| tag length should be less than 32 | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubAccountApiKey**
> DeleteSubAccountApiKey(ctx, subAccountId, subAccountApiKey, timestamp, signature, contentType, xMBXAPIKEY)
Delete Sub Account Api Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **subAccountApiKey** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableFuturesforSubAccount**
> EnableFuturesforSubAccount(ctx, subAccountId, futures, timestamp, signature, contentType, xMBXAPIKEY)
Enable Futures for Sub Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **futures** | **string**| only true for now | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableMarginforSubAccount**
> EnableMarginforSubAccount(ctx, subAccountId, margin, timestamp, signature, contentType, xMBXAPIKEY)
Enable Margin for Sub Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **margin** | **string**| only true for now | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableOrDisableBNBBurnforSubAccountMarginInterest**
> EnableOrDisableBNBBurnforSubAccountMarginInterest(ctx, subAccountId, spotBNBBurn, timestamp, signature, contentType, xMBXAPIKEY)
Enable Or Disable BNB Burn for Sub Account Margin Interest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **spotBNBBurn** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableOrDisableBNBBurnforSubAccountSPOTandMARGIN**
> EnableOrDisableBNBBurnforSubAccountSPOTandMARGIN(ctx, subAccountId, spotBNBBurn, timestamp, signature, contentType, xMBXAPIKEY)
Enable Or Disable BNB Burn for Sub Account SPOT and MARGIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **spotBNBBurn** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateBrokerCommissionRebateHistory**
> GenerateBrokerCommissionRebateHistory(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Generate Broker Commission Rebate History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBNBBurnStatusforSubAccount**
> GetBNBBurnStatusforSubAccount(ctx, subAccountId, timestamp, signature, contentType, xMBXAPIKEY)
Get BNB Burn Status for Sub Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubAccountDepositHistory**
> GetSubAccountDepositHistory(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Sub Account Deposit History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryBrokerCommissionRebateHistory**
> QueryBrokerCommissionRebateHistory(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Broker Commission Rebate History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryBrokerCommissionRebateRecentRecord**
> QueryBrokerCommissionRebateRecentRecord(ctx, subAccountId, startTime, endTime, page, size, timestamp, signature, contentType, xMBXAPIKEY)
Query Broker Commission Rebate Recent Record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **startTime** | **string**|  | 
  **endTime** | **string**|  | 
  **page** | **string**| default 1 | 
  **size** | **string**| default 500, max 500 | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubAccount**
> QuerySubAccount(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Sub Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubAccountApiKey**
> QuerySubAccountApiKey(ctx, subAccountId, timestamp, signature, contentType, xMBXAPIKEY)
Query Sub Account Api Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubAccountFuturesCommissionAdjustment**
> QuerySubAccountFuturesCommissionAdjustment(ctx, subAccountId, timestamp, signature, contentType, xMBXAPIKEY)
Query Sub Account Futures Commission Adjustment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountId** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubAccountSpotAssetinfo**
> QuerySubAccountSpotAssetinfo(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Sub Account Spot Asset info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubAccountTransferHistory**
> QuerySubAccountTransferHistory(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Sub Account Transfer History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubaccountFuturesAssetinfo**
> QuerySubaccountFuturesAssetinfo(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Subaccount Futures Asset info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySubaccountMarginAssetinfo**
> QuerySubaccountMarginAssetinfo(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Subaccount Margin Asset info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubAccountTransfer**
> SubAccountTransfer(ctx, asset, amount, timestamp, signature, contentType, xMBXAPIKEY)
Sub Account Transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **amount** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

