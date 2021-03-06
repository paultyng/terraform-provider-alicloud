package polardb

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

// DescribeSqlLogTrialStatus invokes the polardb.DescribeSqlLogTrialStatus API synchronously
func (client *Client) DescribeSqlLogTrialStatus(request *DescribeSqlLogTrialStatusRequest) (response *DescribeSqlLogTrialStatusResponse, err error) {
	response = CreateDescribeSqlLogTrialStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSqlLogTrialStatusWithChan invokes the polardb.DescribeSqlLogTrialStatus API asynchronously
func (client *Client) DescribeSqlLogTrialStatusWithChan(request *DescribeSqlLogTrialStatusRequest) (<-chan *DescribeSqlLogTrialStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSqlLogTrialStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSqlLogTrialStatus(request)
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

// DescribeSqlLogTrialStatusWithCallback invokes the polardb.DescribeSqlLogTrialStatus API asynchronously
func (client *Client) DescribeSqlLogTrialStatusWithCallback(request *DescribeSqlLogTrialStatusRequest, callback func(response *DescribeSqlLogTrialStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSqlLogTrialStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSqlLogTrialStatus(request)
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

// DescribeSqlLogTrialStatusRequest is the request struct for api DescribeSqlLogTrialStatus
type DescribeSqlLogTrialStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSqlLogTrialStatusResponse is the response struct for api DescribeSqlLogTrialStatus
type DescribeSqlLogTrialStatusResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	EverTrialed    string `json:"EverTrialed" xml:"EverTrialed"`
	OpenTrial      string `json:"OpenTrial" xml:"OpenTrial"`
	RemainDays     string `json:"RemainDays" xml:"RemainDays"`
	UsedDays       string `json:"UsedDays" xml:"UsedDays"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateDescribeSqlLogTrialStatusRequest creates a request to invoke DescribeSqlLogTrialStatus API
func CreateDescribeSqlLogTrialStatusRequest() (request *DescribeSqlLogTrialStatusRequest) {
	request = &DescribeSqlLogTrialStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeSqlLogTrialStatus", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSqlLogTrialStatusResponse creates a response to parse from DescribeSqlLogTrialStatus response
func CreateDescribeSqlLogTrialStatusResponse() (response *DescribeSqlLogTrialStatusResponse) {
	response = &DescribeSqlLogTrialStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
