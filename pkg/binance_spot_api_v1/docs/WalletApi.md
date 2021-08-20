# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**APIKeyPermission**](WalletApi.md#APIKeyPermission) | **Get** /sapi/v1/account/apiRestrictions | API Key Permission
[**AccountAPITradingStatusSAPIUSERDATA**](WalletApi.md#AccountAPITradingStatusSAPIUSERDATA) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status(SAPI) (USER_DATA)
[**AccountStatusSAPIUSERDATA**](WalletApi.md#AccountStatusSAPIUSERDATA) | **Get** /sapi/v1/account/status | Account Status(SAPI) (USER_DATA)
[**AllCoinsInformationUSERDATA**](WalletApi.md#AllCoinsInformationUSERDATA) | **Get** /sapi/v1/capital/config/getall | All Coins&#x27; Information (USER_DATA)
[**AssetDetailSAPIUSERDATA**](WalletApi.md#AssetDetailSAPIUSERDATA) | **Get** /sapi/v1/asset/assetDetail | Asset Detail(SAPI) (USER_DATA)
[**AssetDividendRecordUSERDATA**](WalletApi.md#AssetDividendRecordUSERDATA) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**DailyAccountSnapshotUSERDATA**](WalletApi.md#DailyAccountSnapshotUSERDATA) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**DepositAddressSupportingnetworkUSERDATA**](WalletApi.md#DepositAddressSupportingnetworkUSERDATA) | **Get** /sapi/v1/capital/deposit/address | Deposit Address (supporting network) (USER_DATA)
[**DepositHistorysupportingnetworkUSERDATA**](WalletApi.md#DepositHistorysupportingnetworkUSERDATA) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History（supporting network） (USER_DATA)
[**DisableFastWithdrawSwitchUSERDATA**](WalletApi.md#DisableFastWithdrawSwitchUSERDATA) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**DustLogUSERDATA**](WalletApi.md#DustLogUSERDATA) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**DustTransferUSERDATA**](WalletApi.md#DustTransferUSERDATA) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)
[**EnableFastWithdrawSwitchUSERDATA**](WalletApi.md#EnableFastWithdrawSwitchUSERDATA) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**FundingWallet**](WalletApi.md#FundingWallet) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet
[**QueryUserUniversalTransferHistory**](WalletApi.md#QueryUserUniversalTransferHistory) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History
[**SystemStatusSystem**](WalletApi.md#SystemStatusSystem) | **Get** /sapi/v1/system/status | System Status (System)
[**TradeFeeSAPIUSERDATA**](WalletApi.md#TradeFeeSAPIUSERDATA) | **Get** /sapi/v1/asset/tradeFee | Trade Fee(SAPI) (USER_DATA)
[**UserUniversalTransfer**](WalletApi.md#UserUniversalTransfer) | **Post** /sapi/v1/asset/transfer | User Universal Transfer
[**WithdrawHistorySupportingnetworkUSERDATA**](WalletApi.md#WithdrawHistorySupportingnetworkUSERDATA) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)
[**WithdrawSAPI**](WalletApi.md#WithdrawSAPI) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw [SAPI]

# **APIKeyPermission**
> APIKeyPermission(ctx, timestamp, signature, contentType, xMBXAPIKEY)
API Key Permission

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

# **AccountAPITradingStatusSAPIUSERDATA**
> AccountAPITradingStatusSAPIUSERDATA(ctx, timestamp, signature, xMBXAPIKEY, contentType)
Account API Trading Status(SAPI) (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountStatusSAPIUSERDATA**
> AccountStatusSAPIUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Account Status(SAPI) (USER_DATA)

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

# **AllCoinsInformationUSERDATA**
> AllCoinsInformationUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
All Coins' Information (USER_DATA)

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

# **AssetDetailSAPIUSERDATA**
> AssetDetailSAPIUSERDATA(ctx, timestamp, signature, xMBXAPIKEY, contentType)
Asset Detail(SAPI) (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDividendRecordUSERDATA**
> AssetDividendRecordUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Asset Dividend Record (USER_DATA)

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

# **DailyAccountSnapshotUSERDATA**
> DailyAccountSnapshotUSERDATA(ctx, type_, timestamp, signature, contentType, xMBXAPIKEY)
Daily Account Snapshot (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| \&quot;SPOT\&quot;, \&quot;MARGIN\&quot;, \&quot;FUTURES\&quot; | 
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

# **DepositAddressSupportingnetworkUSERDATA**
> DepositAddressSupportingnetworkUSERDATA(ctx, coin, timestamp, signature, contentType, xMBXAPIKEY)
Deposit Address (supporting network) (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
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

# **DepositHistorysupportingnetworkUSERDATA**
> DepositHistorysupportingnetworkUSERDATA(ctx, coin, timestamp, signature, contentType, xMBXAPIKEY)
Deposit History（supporting network） (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
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

# **DisableFastWithdrawSwitchUSERDATA**
> DisableFastWithdrawSwitchUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Disable Fast Withdraw Switch (USER_DATA)

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

# **DustLogUSERDATA**
> DustLogUSERDATA(ctx, timestamp, signature, xMBXAPIKEY, contentType)
DustLog(USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DustTransferUSERDATA**
> DustTransferUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Dust Transfer (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**| The asset being converted. | 
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

# **EnableFastWithdrawSwitchUSERDATA**
> EnableFastWithdrawSwitchUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Enable Fast Withdraw Switch (USER_DATA)

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

# **FundingWallet**
> FundingWallet(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Funding Wallet

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

# **QueryUserUniversalTransferHistory**
> QueryUserUniversalTransferHistory(ctx, type_, timestamp, signature, contentType, xMBXAPIKEY)
Query User Universal Transfer History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
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

# **SystemStatusSystem**
> SystemStatusSystem(ctx, contentType)
System Status (System)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeFeeSAPIUSERDATA**
> TradeFeeSAPIUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Trade Fee(SAPI) (USER_DATA)

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

# **UserUniversalTransfer**
> UserUniversalTransfer(ctx, type_, asset, amount, timestamp, signature, contentType, xMBXAPIKEY)
User Universal Transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
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

# **WithdrawHistorySupportingnetworkUSERDATA**
> WithdrawHistorySupportingnetworkUSERDATA(ctx, coin, timestamp, signature, contentType, xMBXAPIKEY)
Withdraw History (supporting network) (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
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

# **WithdrawSAPI**
> WithdrawSAPI(ctx, coin, address, amount, timestamp, signature, contentType, xMBXAPIKEY)
Withdraw [SAPI]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **coin** | **string**|  | 
  **address** | **string**|  | 
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

