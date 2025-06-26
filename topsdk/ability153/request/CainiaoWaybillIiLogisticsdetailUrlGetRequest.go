package request


type CainiaoWaybillIiLogisticsdetailUrlGetRequest struct {
    /*
        快递公司编码     */
    CpCode  *string `json:"cp_code" required:"true" `
    /*
        电子面单单号     */
    WaybillCode  *string `json:"waybill_code" required:"true" `
}

func (s *CainiaoWaybillIiLogisticsdetailUrlGetRequest) SetCpCode(v string) *CainiaoWaybillIiLogisticsdetailUrlGetRequest {
    s.CpCode = &v
    return s
}
func (s *CainiaoWaybillIiLogisticsdetailUrlGetRequest) SetWaybillCode(v string) *CainiaoWaybillIiLogisticsdetailUrlGetRequest {
    s.WaybillCode = &v
    return s
}

func (req *CainiaoWaybillIiLogisticsdetailUrlGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CpCode != nil) {
        paramMap["cp_code"] = *req.CpCode
    }
    if(req.WaybillCode != nil) {
        paramMap["waybill_code"] = *req.WaybillCode
    }
    return paramMap
}

func (req *CainiaoWaybillIiLogisticsdetailUrlGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}