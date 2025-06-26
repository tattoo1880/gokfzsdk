package request


type CainiaoWaybillIiDeliveryRequest struct {
    /*
        物流供应商编码     */
    CpCode  *string `json:"cp_code,omitempty" required:"false" `
    /*
        派送类型，1:通知派送； -1: 通知退回     */
    DeliveryAction  *int64 `json:"delivery_action,omitempty" required:"false" `
    /*
        面单号     */
    WaybillCode  *string `json:"waybill_code,omitempty" required:"false" `
}

func (s *CainiaoWaybillIiDeliveryRequest) SetCpCode(v string) *CainiaoWaybillIiDeliveryRequest {
    s.CpCode = &v
    return s
}
func (s *CainiaoWaybillIiDeliveryRequest) SetDeliveryAction(v int64) *CainiaoWaybillIiDeliveryRequest {
    s.DeliveryAction = &v
    return s
}
func (s *CainiaoWaybillIiDeliveryRequest) SetWaybillCode(v string) *CainiaoWaybillIiDeliveryRequest {
    s.WaybillCode = &v
    return s
}

func (req *CainiaoWaybillIiDeliveryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CpCode != nil) {
        paramMap["cp_code"] = *req.CpCode
    }
    if(req.DeliveryAction != nil) {
        paramMap["delivery_action"] = *req.DeliveryAction
    }
    if(req.WaybillCode != nil) {
        paramMap["waybill_code"] = *req.WaybillCode
    }
    return paramMap
}

func (req *CainiaoWaybillIiDeliveryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}