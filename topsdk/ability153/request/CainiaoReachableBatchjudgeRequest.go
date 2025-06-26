package request

import (
        "topsdk/ability153/domain"
        "topsdk/util"
    )

type CainiaoReachableBatchjudgeRequest struct {
    /*
        1:快递 2:快运     */
    AddressType  *int64 `json:"address_type" required:"true" `
    /*
        收发信息     */
    Data  *domain.CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto `json:"data" required:"true" `
    /*
        调用方对象     */
    ClientInfo  *domain.CainiaoReachableBatchjudgeClientInfoDto `json:"client_info" required:"true" `
}

func (s *CainiaoReachableBatchjudgeRequest) SetAddressType(v int64) *CainiaoReachableBatchjudgeRequest {
    s.AddressType = &v
    return s
}
func (s *CainiaoReachableBatchjudgeRequest) SetData(v domain.CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto) *CainiaoReachableBatchjudgeRequest {
    s.Data = &v
    return s
}
func (s *CainiaoReachableBatchjudgeRequest) SetClientInfo(v domain.CainiaoReachableBatchjudgeClientInfoDto) *CainiaoReachableBatchjudgeRequest {
    s.ClientInfo = &v
    return s
}

func (req *CainiaoReachableBatchjudgeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AddressType != nil) {
        paramMap["address_type"] = *req.AddressType
    }
    if(req.Data != nil) {
        paramMap["data"] = util.ConvertStruct(*req.Data)
    }
    if(req.ClientInfo != nil) {
        paramMap["client_info"] = util.ConvertStruct(*req.ClientInfo)
    }
    return paramMap
}

func (req *CainiaoReachableBatchjudgeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}