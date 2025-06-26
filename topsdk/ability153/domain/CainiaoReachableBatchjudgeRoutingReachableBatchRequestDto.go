package domain


type CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto struct {
    /*
        快递公司code     */
    CpCode  *string `json:"cp_code,omitempty" `

    /*
        收发地址和服务列表     */
    AddressAndServiceList  *[]CainiaoReachableBatchjudgeReachableAddressAndServiceDto `json:"address_and_service_list,omitempty" `

    /*
        发件揽收网点     */
    SendBranchCode  *string `json:"send_branch_code,omitempty" `

}

func (s *CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto) SetCpCode(v string) *CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto {
    s.CpCode = &v
    return s
}
func (s *CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto) SetAddressAndServiceList(v []CainiaoReachableBatchjudgeReachableAddressAndServiceDto) *CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto {
    s.AddressAndServiceList = &v
    return s
}
func (s *CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto) SetSendBranchCode(v string) *CainiaoReachableBatchjudgeRoutingReachableBatchRequestDto {
    s.SendBranchCode = &v
    return s
}
