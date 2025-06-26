package domain


type CainiaoWaybillIiDeliveryWaybillOrderConfirmResponse struct {
    /*
        面单号     */
    WaybillCode  *string `json:"waybill_code,omitempty" `

}

func (s *CainiaoWaybillIiDeliveryWaybillOrderConfirmResponse) SetWaybillCode(v string) *CainiaoWaybillIiDeliveryWaybillOrderConfirmResponse {
    s.WaybillCode = &v
    return s
}
