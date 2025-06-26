package domain


type CainiaoWaybillIiConfirmWaybillOrderConfirmRequest struct {
    /*
        预约上门收件     */
    CallDoorPickUp  *bool `json:"call_door_pick_up,omitempty" `

    /*
        cpCode     */
    CpCode  *string `json:"cp_code,omitempty" `

    /*
        预约上门截止时间     */
    DoorPickUpEndTime  *string `json:"door_pick_up_end_time,omitempty" `

    /*
        预约上门时间     */
    DoorPickUpTime  *string `json:"door_pick_up_time,omitempty" `

    /*
        扩展信息，json String     */
    ExtraInfo  *string `json:"extra_info,omitempty" `

    /*
        物流服务， json String     */
    LogisticsServices  *string `json:"logistics_services,omitempty" `

    /*
        快递产品编码     */
    ProductCode  *string `json:"product_code,omitempty" `

    /*
        客户订单货物总高，单位厘米     */
    TotalHeight  *int64 `json:"total_height,omitempty" `

    /*
        订单货物总长,单位厘米     */
    TotalLength  *int64 `json:"total_length,omitempty" `

    /*
        订单货物总宽，单位厘米     */
    TotalWidth  *int64 `json:"total_width,omitempty" `

    /*
        物流单号信息     */
    WaybillInfo  *[]CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo `json:"waybill_info,omitempty" `

    /*
        货物总体积，单位立方厘米     */
    TotalVolume  *int64 `json:"total_volume,omitempty" `

    /*
        货物总重量，单位g     */
    TotalWeight  *int64 `json:"total_weight,omitempty" `

}

func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetCallDoorPickUp(v bool) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.CallDoorPickUp = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetCpCode(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.CpCode = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetDoorPickUpEndTime(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.DoorPickUpEndTime = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetDoorPickUpTime(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.DoorPickUpTime = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetExtraInfo(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.ExtraInfo = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetLogisticsServices(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.LogisticsServices = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetProductCode(v string) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.ProductCode = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetTotalHeight(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.TotalHeight = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetTotalLength(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.TotalLength = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetTotalWidth(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.TotalWidth = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetWaybillInfo(v []CainiaoWaybillIiConfirmWaybillOrderConfirmWaybillInfo) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.WaybillInfo = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetTotalVolume(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.TotalVolume = &v
    return s
}
func (s *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest) SetTotalWeight(v int64) *CainiaoWaybillIiConfirmWaybillOrderConfirmRequest {
    s.TotalWeight = &v
    return s
}
