//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

func TestRouterAPIService_GetAllRouters(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	templateName := "test_template_get_all_routers"

	tests := []struct {
		name    string
		param   map[string]string
		given   func(m *MockAPIClientHandler)
		want    models.GetAllNetworkRouter
		wantErr bool
	}{
		{
			name: "Normal Test case 1: Get all Routers",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"networkRouters": [{
							"id": 1,
							"name": "test_template_get_all_routers"
						}]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.GetAllNetworkRouter{
				NetworkRouters: []models.GetNetworkRouter{
					{
						ID:   1,
						Name: templateName,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Failed Test case 2: Error in prepare request",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(nil, errors.New("prepare error request"))
			},
			want:    models.GetAllNetworkRouter{},
			wantErr: true,
		},
		{
			name: "Failed Test case 3: Error in callAPI",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.GetAllNetworkRouter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			n := RouterAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := n.GetAllRouter(ctx, tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("RouterAPIService.GetAllRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterAPIService.GetAllRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterAPIService_GetNetworkRouterTypes(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	templateName := "test_template_get_all_router_types"

	tests := []struct {
		name    string
		param   map[string]string
		given   func(m *MockAPIClientHandler)
		want    models.GetNetworlRouterTypes
		wantErr bool
	}{
		{
			name: "Normal Test case 1: Get all Router Types",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/network-router-types"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"networkRouterTypes": [{
							"id": 1,
							"name": "test_template_get_all_router_types"
						}]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.GetNetworlRouterTypes{
				NetworkRouterTypes: []models.NetworkRouterTypes{
					{
						ID:   1,
						Name: templateName,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Failed Test case 2: Error in prepare request",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/network-router-types"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(nil, errors.New("prepare error request"))
			},
			want:    models.GetNetworlRouterTypes{},
			wantErr: true,
		},
		{
			name: "Failed Test case 3: Error in callAPI",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/network-router-types"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.GetNetworlRouterTypes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			n := RouterAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := n.GetRouterTypes(ctx, tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("RouterAPIService.GetRouterTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterAPIService.GetRouterTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRouterAPIService_CreateRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		args    models.CreateRouterRequest
		given   func(m *MockAPIClientHandler)
		want    models.CreateRouterResp
		wantErr bool
	}{
		{
			name: "Normal test case 1: Create Router",
			args: models.CreateRouterRequest{
				NetworkRouter: models.CreateRouterRequestRouter{
					Name: "tf_router",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers"
				method := "POST"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)

				m.EXPECT().prepareRequest(gomock.Any(), path, method,
					models.CreateRouterRequest{
						NetworkRouter: models.CreateRouterRequestRouter{
							Name: "tf_router",
						},
					},
					headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"success": true,
						"id": 16
					}
				`))),
				}, nil)
			},
			want: models.CreateRouterResp{
				Success: true,
				ID:      16,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			tt.given(mockAPIClient)
			n := RouterAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			got, err := n.CreateRouter(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("RouterAPIService.CreateRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterAPIService.CreateRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterAPIService_UpdateRouter(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	headers := getDefaultHeaders()
	defer ctrl.Finish()
	tests := []struct {
		name     string
		routerID int
		param    models.CreateRouterRequest
		given    func(m *MockAPIClientHandler)
		want     models.SuccessOrErrorMessage
		wantErr  bool
	}{
		{
			name:     "Normal Test case 1: Update a Router",
			routerID: 1,
			param: models.CreateRouterRequest{
				NetworkRouter: models.CreateRouterRequestRouter{
					Name: "test_update_router_name",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers/1"
				method := "PUT"
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"networkRouter": {
						"name": "test_update_router_name"
					}
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
						"success": true
				}`)))
				pBody := models.CreateRouterRequest{
					NetworkRouter: models.CreateRouterRequestRouter{
						Name: "test_update_router_name",
					},
				}
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.SuccessOrErrorMessage{
				Success: true,
			},
			wantErr: false,
		},
		{
			name:     "Failed test case 2: error in prepare request",
			routerID: 1,
			param: models.CreateRouterRequest{
				NetworkRouter: models.CreateRouterRequestRouter{
					Name: "test_update_router_name",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers/1"
				method := "PUT"

				pBody := models.CreateRouterRequest{
					NetworkRouter: models.CreateRouterRequestRouter{
						Name: "test_update_router_name",
					},
				}
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
		{
			name:     "Failed test case 3: error in callAPI",
			routerID: 1,
			param: models.CreateRouterRequest{
				NetworkRouter: models.CreateRouterRequestRouter{
					Name: "test_update_router_name",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/networks/routers/1"
				method := "PUT"
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"networkRouter": {
						"name": "test_update_router_name"
					}
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				}`)))
				pBody := models.CreateRouterRequest{
					NetworkRouter: models.CreateRouterRequestRouter{
						Name: "test_update_router_name",
					},
				}
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := RouterAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.UpdateRouter(ctx, tt.routerID, tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("RouterAPIService.UpdateRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterAPIService.UpdateRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}