package domain


type CainiaoWaybillIiConfirmWaybillOrderConfirmResponse struct {
    /*
        面单号     */
    WaybillCode  *string `json:"waybill_code,omitempty" `

}

func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmResponse) SetWaybillCode(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmResponse {
    s.WaybillCode = &v
    return s
}
