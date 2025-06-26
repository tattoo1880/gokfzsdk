package request

import (
        "topsdk/ability229/domain"
        "topsdk/util"
    )

type CainiaoWaybillCloudprintNetprintBindRequest struct {
    /*
        req     */
    Params  *domain.CainiaoWaybillCloudprintNetprintBindCloudPrinterBindRequest `json:"params" required:"true" `
}

func (s *CainiaoWaybillCloudprintNetprintBindRequest) SetParams(v domain.CainiaoWaybillCloudprintNetprintBindCloudPrinterBindRequest) *CainiaoWaybillCloudprintNetprintBindRequest {
    s.Params = &v
    return s
}

func (req *CainiaoWaybillCloudprintNetprintBindRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Params != nil) {
        paramMap["params"] = util.ConvertStruct(*req.Params)
    }
    return paramMap
}

func (req *CainiaoWaybillCloudprintNetprintBindRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}