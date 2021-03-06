package dcdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDcdnService invokes the dcdn.ModifyDcdnService API synchronously
func (client *Client) ModifyDcdnService(request *ModifyDcdnServiceRequest) (response *ModifyDcdnServiceResponse, err error) {
	response = CreateModifyDcdnServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDcdnServiceWithChan invokes the dcdn.ModifyDcdnService API asynchronously
func (client *Client) ModifyDcdnServiceWithChan(request *ModifyDcdnServiceRequest) (<-chan *ModifyDcdnServiceResponse, <-chan error) {
	responseChan := make(chan *ModifyDcdnServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDcdnService(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDcdnServiceWithCallback invokes the dcdn.ModifyDcdnService API asynchronously
func (client *Client) ModifyDcdnServiceWithCallback(request *ModifyDcdnServiceRequest, callback func(response *ModifyDcdnServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDcdnServiceResponse
		var err error
		defer close(result)
		response, err = client.ModifyDcdnService(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDcdnServiceRequest is the request struct for api ModifyDcdnService
type ModifyDcdnServiceRequest struct {
	*requests.RpcRequest
	WebsocketBillType string           `position:"Query" name:"WebsocketBillType"`
	BillType          string           `position:"Query" name:"BillType"`
	SecurityToken     string           `position:"Query" name:"SecurityToken"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDcdnServiceResponse is the response struct for api ModifyDcdnService
type ModifyDcdnServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDcdnServiceRequest creates a request to invoke ModifyDcdnService API
func CreateModifyDcdnServiceRequest() (request *ModifyDcdnServiceRequest) {
	request = &ModifyDcdnServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "ModifyDcdnService", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDcdnServiceResponse creates a response to parse from ModifyDcdnService response
func CreateModifyDcdnServiceResponse() (response *ModifyDcdnServiceResponse) {
	response = &ModifyDcdnServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
