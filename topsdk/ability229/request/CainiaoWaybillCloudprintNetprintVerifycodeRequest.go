package request

import (
        "topsdk/ability229/domain"
        "topsdk/util"
    )

type CainiaoWaybillCloudprintNetprintVerifycodeRequest struct {
    /*
        请求     */
    Printer  *domain.CainiaoWaybillCloudprintNetprintVerifycodeCloudPrinterVerifyCodeRequest `json:"printer" required:"true" `
}

func (s *CainiaoWaybillCloudprintNetprintVerifycodeRequest) SetPrinter(v domain.CainiaoWaybillCloudprintNetprintVerifycodeCloudPrinterVerifyCodeRequest) *CainiaoWaybillCloudprintNetprintVerifycodeRequest {
    s.Printer = &v
    return s
}

func (req *CainiaoWaybillCloudprintNetprintVerifycodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Printer != nil) {
        paramMap["printer"] = util.ConvertStruct(*req.Printer)
    }
    return paramMap
}

func (req *CainiaoWaybillCloudprintNetprintVerifycodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}