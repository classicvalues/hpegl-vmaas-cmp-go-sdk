// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"

	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
	models "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

// Linger please
var (
	_ context.Context
)

type InstancesApiService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
InstancesApiService
Clone an instance and all VM within that instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiCloneAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CloneInstanceBody) -

*/

type InstancesApiCloneAnInstanceOpts struct {
	Body optional.Interface
}

func (a *InstancesApiService) CloneAnInstance(ctx context.Context, instanceId int, localVarOptionals *models.CreateInstanceBody) (models.SuccessOrErrorMessage, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	var cloneResp models.SuccessOrErrorMessage
	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/clone", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return cloneResp, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return cloneResp, err
	}
	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return cloneResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return cloneResp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return cloneResp, err
	}
	if err := json.Unmarshal(localVarBody, &cloneResp); err != nil {
		return cloneResp, err
	}

	return cloneResp, nil
}

/*
InstancesApiService
Creates an image template from an existing instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiCloneToImageOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CloneToImage) -

*/

type InstancesApiCloneToImageOpts struct {
	Body optional.Interface
}

func (a *InstancesApiService) CloneToImage(ctx context.Context, instanceId int, localVarOptionals *InstancesApiCloneToImageOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/clone-image", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Create an Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param optional nil or *InstancesApiCreateAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CreateInstanceBody) -
@return models.GetInstanceResponse
*/

func (a *InstancesApiService) CreateAnInstance(ctx context.Context, localVarOptionals *models.CreateInstanceBody) (models.GetInstanceResponse, error) {
	var (
		localVarHttpMethod     = strings.ToUpper("Post")
		localVarPostBody       interface{}
		localVarFileName       string
		localVarFileBytes      []byte
		createInstanceResponse models.GetInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return createInstanceResponse, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return createInstanceResponse, err
	}
	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return createInstanceResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return createInstanceResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return createInstanceResponse, err
	}

	if err = json.Unmarshal(localVarBody, &createInstanceResponse); err != nil {
		return createInstanceResponse, err
	}
	return createInstanceResponse, nil
}

/*
InstancesApiService
Will delete an instance and all associated monitors and backups.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiDeleteAnIstanceOpts - Optional Parameters:
     * @param "Force" (optional.String) -

*/

type InstancesApiDeleteAnIstanceOpts struct {
	Force optional.String
}

func (a *InstancesApiService) DeleteAnInstance(ctx context.Context, instanceId int) (models.SuccessOrErrorMessage, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return models.SuccessOrErrorMessage{}, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return models.SuccessOrErrorMessage{}, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	// fmt.Println(string(localVarBody))
	var instancesResponse models.SuccessOrErrorMessage
	if err = json.Unmarshal(localVarBody, &instancesResponse); err != nil {
		return models.SuccessOrErrorMessage{}, err
	}
	return instancesResponse, nil
}

/*
InstancesApiService
Get a Specific Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
@return models.GetInstanceResponse
*/
func (a *InstancesApiService) GetASpecificInstance(ctx context.Context, instanceId int) (models.GetInstanceResponse, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		getInstanceResponse models.GetInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return getInstanceResponse, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return getInstanceResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return getInstanceResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return getInstanceResponse, err
	}

	if err = json.Unmarshal(localVarBody, &getInstanceResponse); err != nil {
		return getInstanceResponse, err
	}
	return getInstanceResponse, nil
}

/*
InstancesApiService
Fetch the list of available instance types
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId

*/
func (a *InstancesApiService) GetAllInstanceTypesForProvisioning(ctx context.Context) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstanceTypesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Get All Instances
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param optional nil or *InstancesApiGetAllInstancesOpts - Optional Parameters:
     * @param "Max" (optional.Int32) -  Max number of results to return
     * @param "Name" (optional.String) -  Filter by name
     * @param "Phrase" (optional.String) -  Filter by wildcard search of name and description
     * @param "InstanceType" (optional.Int32) -  Filter by instance type code
     * @param "CreatedBy" (optional.String) -  Filter by Created By (User) ID. Accepts multiple values
     * @param "Status" (optional.String) -  Filter by instance status
     * @param "ShowDeleted" (optional.Bool) -  If true, includes instances in pending removal status.
     * @param "Deleted" (optional.Bool) -  If true, only instances in pending removal status are returned.
     * @param "Labels" (optional.String) -  Filter by labels
     * @param "Tags" (optional.String) -  Filter by tags

*/

type InstancesApiGetAllInstancesOpts struct {
	Max          optional.Int32
	Name         optional.String
	Phrase       optional.String
	InstanceType optional.Int32
	CreatedBy    optional.String
	Status       optional.String
	ShowDeleted  optional.Bool
	Deleted      optional.Bool
	Labels       optional.String
	Tags         optional.String
}

func (a *InstancesApiService) GetAllInstances(ctx context.Context, queryParams map[string]string) (models.Instances, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		instancesResponse  models.Instances
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(queryParams)
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return instancesResponse, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return instancesResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return instancesResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return instancesResponse, err
	}

	if err = json.Unmarshal(localVarBody, &instancesResponse); err != nil {
		return instancesResponse, err
	}
	return instancesResponse, nil
}

/*
InstancesApiService
This endpoint retrieves all the Service Plans available for the specified cloud and instance layout. The response includes details about the plans and their configuration options. It may be used to get the list of available plans when creating a new instance or resizing an existing instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param zoneId Cloud ID
 * @param layoutId Instance Layout ID
 * @param optional nil or *InstancesApiGetAvailableServicePlansForAnInstanceOpts - Optional Parameters:
     * @param "SiteId" (optional.Int32) -  Group ID
@return models.GetServicePlanResponse
*/

type InstancesApiGetAvailableServicePlansForAnInstanceOpts struct {
	SiteId optional.Int32
}

func (a *InstancesApiService) GetAvailableServicePlansForAnInstance(ctx context.Context, zoneId int, layoutId int, localVarOptionals *InstancesApiGetAvailableServicePlansForAnInstanceOpts) (models.GetServicePlanResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetServicePlanResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/service-plans", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("zoneId", parameterToString(zoneId, ""))
	localVarQueryParams.Add("layoutId", parameterToString(layoutId, ""))
	if localVarOptionals != nil && localVarOptionals.SiteId.IsSet() {
		localVarQueryParams.Add("siteId", parameterToString(localVarOptionals.SiteId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InstancesApiService
List all environment variables associated with the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) GetEnvVariables(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/envs", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Retrieves the process history for a specific instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) GetInstanceHistory(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/history", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Lists VMware Snapshot of the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) GetListOfSnapshotsForAnInstance(ctx context.Context, instanceId int) (models.ListSnapshotResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		snapshotResponse   models.ListSnapshotResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/snapshots", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return snapshotResponse, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return snapshotResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return snapshotResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return snapshotResponse, err
	}
	if err = json.Unmarshal(localVarBody, &snapshotResponse); err != nil {
		return snapshotResponse, err
	}
	return snapshotResponse, nil
}

/*
InstancesApiService
Fetch an instance type by ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceTypeId

*/
func (a *InstancesApiService) GetSpecificInstanceTypeForProvisioning(ctx context.Context, instanceTypeId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstanceTypesPath, instanceTypeId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}
	return localVarHttpResponse, nil
}

/*
InstancesApiService
It is possible to import a snapshot of an instance. This creates a Virtual Image of the instance as it currently exists.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiImportSnapshotOfAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ImportSnapshotBody) -

*/

func (a *InstancesApiService) ImportSnapshotOfAnInstance(ctx context.Context, instanceId int, localVarOptionals *models.ImportSnapshotBody) (models.SuccessOrErrorMessage, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/import-snapshot", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return models.SuccessOrErrorMessage{}, err
		}

	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return models.SuccessOrErrorMessage{}, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return models.SuccessOrErrorMessage{}, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	var instanceResponse models.SuccessOrErrorMessage
	if err = json.Unmarshal(localVarBody, &instanceResponse); err != nil {
		return models.SuccessOrErrorMessage{}, err
	}
	return instanceResponse, nil
}

/*
InstancesApiService
&#x27;This will lock the instance. While locked, instances may not be removed.&#x27;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) LockAnInstance(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/lock", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
It is possible to resize VMs within an instance by increasing their memory plan or storage limit. This is done by assigning a new service plan to the VM.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiResizeAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ResizeInstanceBody) -
@return models.GetInstanceResponse
*/

func (a *InstancesApiService) ResizeAnInstance(ctx context.Context, instanceId int, localVarOptionals *models.ResizeInstanceBody) (models.ResizeInstanceResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		resizeResponse     models.ResizeInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/resize", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return resizeResponse, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return resizeResponse, err
	}
	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resizeResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return resizeResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return resizeResponse, err
	}

	if err = json.Unmarshal(localVarBody, &resizeResponse); err != nil {
		return resizeResponse, err
	}
	return resizeResponse, nil
}

/*
InstancesApiService
Restarts all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) RestartAnInstance(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/restart", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Creates VMware Snapshot of the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiSnapshotAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of SnapshotBody) -

*/

func (a *InstancesApiService) SnapshotAnInstance(ctx context.Context, instanceId int, localVarOptionals *models.SnapshotBody) (models.Instances, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFileName     string
		localVarFileBytes    []byte
		snapshotInstanceresp models.Instances
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/snapshot", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return snapshotInstanceresp, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return snapshotInstanceresp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return snapshotInstanceresp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return snapshotInstanceresp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return snapshotInstanceresp, err
	}

	if err := json.Unmarshal(localVarBody, &snapshotInstanceresp); err != nil {
		return snapshotInstanceresp, err
	}
	return snapshotInstanceresp, nil
}

/*
InstancesApiService
Starts all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) StartAnInstance(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/start", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Stops all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) StopAnInstance(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/stop", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Suspends all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) SuspendAnInstance(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/suspend", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Undo the delete of an instance that is in pending removal state
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
@return models.GetInstanceResponse
*/
func (a *InstancesApiService) UndoDeleteOfAnInstance(ctx context.Context, instanceId int) (models.GetInstanceResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/cancel-removal", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InstancesApiService
Unlocks the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId

*/
func (a *InstancesApiService) UnlockAnInstance(ctx context.Context, instanceId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/unlock", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
InstancesApiService
Updating an Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param instanceId
 * @param optional nil or *InstancesApiUpdatingAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of UpdateInstanceBody) -
@return models.GetInstanceResponse
*/

type InstancesApiUpdatingAnInstanceOpts struct {
	Body optional.Interface
}

func (a *InstancesApiService) UpdatingAnInstance(ctx context.Context, instanceId int, localVarOptionals *models.UpdateInstanceBody) (models.UpdateInstanceResponse, error) {
	var (
		localVarHttpMethod     = strings.ToUpper("Put")
		localVarPostBody       interface{}
		localVarFileName       string
		localVarFileBytes      []byte
		updateInstanceResponse models.UpdateInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		consts.InstancesPath, instanceId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return updateInstanceResponse, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return updateInstanceResponse, err
	}
	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return updateInstanceResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return updateInstanceResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return updateInstanceResponse, err
	}
	if err = json.Unmarshal(localVarBody, &updateInstanceResponse); err != nil {
		return updateInstanceResponse, err
	}
	return updateInstanceResponse, nil
}
