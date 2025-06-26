package domain


type CainiaoReachableBatchjudgeReachableServiceWaybillForTopResponseDto struct {
    /*
        结果列表     */
    ResultList  *[]CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto `json:"result_list,omitempty" `

}

func (s *CainiaoReachableBatchjudgeReachableServiceWaybillForTopResponseDto) SetResultList(v []CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto) *CainiaoReachableBatchjudgeReachableServiceWaybillForTopResponseDto {
    s.ResultList = &v
    return s
}
