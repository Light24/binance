
/*
 * Binance spot API
 *
 * Binance official supported Postman collections.<br/> - API documents: https://binance-docs.github.io/apidocs/spot/en/#change-log - Telegram: https://t.me/binance_api_english - Open Issue at: https://github.com/binance-exchange/binance-api-postman
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type TradeApiService service
/*
TradeApiService Account Information (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) AccountInformationUSERDATA(ctx context.Context, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Account Trade List (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) AccountTradeListUSERDATA(ctx context.Context, symbol string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/myTrades"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService All Orders (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) AllOrdersUSERDATA(ctx context.Context, symbol string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/allOrders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Cancel OCO (TRADE)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) CancelOCOTRADE(ctx context.Context, symbol string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Cancel Order
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param orderId
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) CancelOrder(ctx context.Context, symbol string, orderId string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("orderId", parameterToString(orderId, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Cancel all Open Orders on a Symbol (TRADE)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) CancelallOpenOrdersonaSymbolTRADE(ctx context.Context, symbol string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/openOrders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Current Open Orders (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) CurrentOpenOrdersUSERDATA(ctx context.Context, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/openOrders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService New OCO (TRADE)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param side
 * @param quantity
 * @param price
 * @param stopPrice
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) NewOCOTRADE(ctx context.Context, symbol string, side string, quantity string, price string, stopPrice string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/oco"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("side", parameterToString(side, ""))
	localVarQueryParams.Add("quantity", parameterToString(quantity, ""))
	localVarQueryParams.Add("price", parameterToString(price, ""))
	localVarQueryParams.Add("stopPrice", parameterToString(stopPrice, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService New Order
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param side
 * @param type_
 * @param quantity
 * @param price
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) NewOrder(ctx context.Context, symbol string, side string, type_ string, quantity float64, price int32, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("side", parameterToString(side, ""))
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	localVarQueryParams.Add("quantity", parameterToString(quantity, ""))
	localVarQueryParams.Add("price", parameterToString(price, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Query OCO (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) QueryOCOUSERDATA(ctx context.Context, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Query Open OCO (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) QueryOpenOCOUSERDATA(ctx context.Context, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/openOrderList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Query Order (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) QueryOrderUSERDATA(ctx context.Context, symbol string, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Query all OCO (USER_DATA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) QueryallOCOUSERDATA(ctx context.Context, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/allOrderList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
TradeApiService Test New Order (TRADE)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @param side
 * @param type_
 * @param quantity
 * @param price
 * @param timestamp
 * @param signature
 * @param contentType
 * @param xMBXAPIKEY

*/
func (a *TradeApiService) TestNewOrderTRADE(ctx context.Context, symbol string, side string, type_ string, quantity float64, price int32, timestamp string, signature string, contentType string, xMBXAPIKEY string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/test"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("symbol", parameterToString(symbol, ""))
	localVarQueryParams.Add("side", parameterToString(side, ""))
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	localVarQueryParams.Add("quantity", parameterToString(quantity, ""))
	localVarQueryParams.Add("price", parameterToString(price, ""))
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(signature, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["X-MBX-APIKEY"] = parameterToString(xMBXAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
