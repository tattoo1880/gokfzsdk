package request

import (
        "topsdk/ability153/domain"
        "topsdk/util"
    )

type CainiaoWaybillIiConfirmRequest struct {
    /*
        订单确认信息     */
    ParamWaybillOrderConfirmRequest  *domain.CainiaoWaybillIiConfirmWaybillOrderConfirmRequest `json:"param_waybill_order_confirm_request,omitempty" required:"false" `
}

func (s *CainiaoWaybillIiConfirmRequest) SetParamWaybillOrderConfirmRequest(v domain.CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) *CainiaoWaybillIiConfirmRequest {
    s.ParamWaybillOrderConfirmRequest = &v
    return s
}

func (req *CainiaoWaybillIiConfirmRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ParamWaybillOrderConfirmRequest != nil) {
        paramMap["param_waybill_order_confirm_request"] = util.ConvertStruct(*req.ParamWaybillOrderConfirmRequest)
    }
    return paramMap
}

func (req *CainiaoWaybillIiConfirmRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}