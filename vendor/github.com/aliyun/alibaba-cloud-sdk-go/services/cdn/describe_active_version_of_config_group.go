package cdn

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

// DescribeActiveVersionOfConfigGroup invokes the cdn.DescribeActiveVersionOfConfigGroup API synchronously
// api document: https://help.aliyun.com/api/cdn/describeactiveversionofconfiggroup.html
func (client *Client) DescribeActiveVersionOfConfigGroup(request *DescribeActiveVersionOfConfigGroupRequest) (response *DescribeActiveVersionOfConfigGroupResponse, err error) {
	response = CreateDescribeActiveVersionOfConfigGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActiveVersionOfConfigGroupWithChan invokes the cdn.DescribeActiveVersionOfConfigGroup API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeactiveversionofconfiggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveVersionOfConfigGroupWithChan(request *DescribeActiveVersionOfConfigGroupRequest) (<-chan *DescribeActiveVersionOfConfigGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeActiveVersionOfConfigGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActiveVersionOfConfigGroup(request)
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

// DescribeActiveVersionOfConfigGroupWithCallback invokes the cdn.DescribeActiveVersionOfConfigGroup API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeactiveversionofconfiggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveVersionOfConfigGroupWithCallback(request *DescribeActiveVersionOfConfigGroupRequest, callback func(response *DescribeActiveVersionOfConfigGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActiveVersionOfConfigGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeActiveVersionOfConfigGroup(request)
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

// DescribeActiveVersionOfConfigGroupRequest is the request struct for api DescribeActiveVersionOfConfigGroup
type DescribeActiveVersionOfConfigGroupRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Env           string           `position:"Query" name:"Env"`
	ConfigGroupId string           `position:"Query" name:"ConfigGroupId"`
}

// DescribeActiveVersionOfConfigGroupResponse is the response struct for api DescribeActiveVersionOfConfigGroup
type DescribeActiveVersionOfConfigGroupResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	VersionId     string `json:"VersionId" xml:"VersionId"`
	ConfigGroupId string `json:"ConfigGroupId" xml:"ConfigGroupId"`
	BaseVersionId string `json:"BaseVersionId" xml:"BaseVersionId"`
	Description   string `json:"Description" xml:"Description"`
	SeqId         int64  `json:"SeqId" xml:"SeqId"`
	Status        string `json:"Status" xml:"Status"`
	Operator      string `json:"Operator" xml:"Operator"`
	CreateTime    string `json:"CreateTime" xml:"CreateTime"`
	UpdateTime    string `json:"UpdateTime" xml:"UpdateTime"`
}

// CreateDescribeActiveVersionOfConfigGroupRequest creates a request to invoke DescribeActiveVersionOfConfigGroup API
func CreateDescribeActiveVersionOfConfigGroupRequest() (request *DescribeActiveVersionOfConfigGroupRequest) {
	request = &DescribeActiveVersionOfConfigGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeActiveVersionOfConfigGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeActiveVersionOfConfigGroupResponse creates a response to parse from DescribeActiveVersionOfConfigGroup response
func CreateDescribeActiveVersionOfConfigGroupResponse() (response *DescribeActiveVersionOfConfigGroupResponse) {
	response = &DescribeActiveVersionOfConfigGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
