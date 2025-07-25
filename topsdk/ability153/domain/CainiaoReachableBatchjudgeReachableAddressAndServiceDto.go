package domain


type CainiaoReachableBatchjudgeReachableAddressAndServiceDto struct {
    /*
        收货地址     */
    ReceiveAddress  *CainiaoReachableBatchjudgeReceiveAddress `json:"receive_address,omitempty" `

    /*
        服务列表,每一项必须为json的string格式,快运必填,快递为空则默认为'标准快递'     */
    ServiceCodeList  *[]string `json:"service_code_list,omitempty" `

    /*
        每条收发地址的key,用户自定义,每次请求多个地址不能重复     */
    ObjectId  *string `json:"object_id,omitempty" `

    /*
        发货地址     */
    SendAddress  *CainiaoReachableBatchjudgeAddressDto `json:"send_address,omitempty" `

    /*
        淘宝开放地址ID     */
    Oaid  *string `json:"oaid,omitempty" `

    /*
        订单id     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        C2M&1688开放地址ID     */
    Caid  *string `json:"caid,omitempty" `

}

func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetReceiveAddress(v CainiaoReachableBatchjudgeReceiveAddress) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.ReceiveAddress = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetServiceCodeList(v []string) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.ServiceCodeList = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetObjectId(v string) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.ObjectId = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetSendAddress(v CainiaoReachableBatchjudgeAddressDto) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.SendAddress = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetOaid(v string) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.Oaid = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetOrderId(v int64) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.OrderId = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableAddressAndServiceDto) SetCaid(v string) *CainiaoReachableBatchjudgeReachableAddressAndServiceDto {
    s.Caid = &v
    return s
}
