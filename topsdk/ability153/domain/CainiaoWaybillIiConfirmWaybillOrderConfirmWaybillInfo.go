package domain


type CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo struct {
    /*
        包裹高，单位厘米     */
    Height  *int64 `json:"height,omitempty" `

    /*
        包裹长，单位厘米     */
    Length  *int64 `json:"length,omitempty" `

    /*
        体积, 单位 ml     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        面单号     */
    WaybillCode  *string `json:"waybill_code,omitempty" `

    /*
        重量,单位 g     */
    Weight  *int64 `json:"weight,omitempty" `

    /*
        包裹宽，单位厘米     */
    Width  *int64 `json:"width,omitempty" `

}

func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) SetHeight(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo {
    s.Height = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) SetLength(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo {
    s.Length = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) SetVolume(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo {
    s.Volume = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) SetWaybillCode(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo {
    s.WaybillCode = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) SetWeight(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo {
    s.Weight = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) SetWidth(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo {
    s.Width = &v
    return s
}
