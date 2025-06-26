package domain


type CainiaoReachableBatchjudgeBaseResultDto struct {
    /*
        请求错误信息     */
    OneErrorInfo  *CainiaoReachableBatchjudgeErrorInfo `json:"one_error_info,omitempty" `

    /*
        请求是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回信息     */
    Module  *CainiaoReachableBatchjudgeReachableServiceWaybillForTopResponseDto `json:"module,omitempty" `

}

func (s *CainiaoReachableBatchjudgeBaseResultDto) SetOneErrorInfo(v CainiaoReachableBatchjudgeErrorInfo) *CainiaoReachableBatchjudgeBaseResultDto {
    s.OneErrorInfo = &v
    return s
}
func (s *CainiaoReachableBatchjudgeBaseResultDto) SetSuccess(v bool) *CainiaoReachableBatchjudgeBaseResultDto {
    s.Success = &v
    return s
}
func (s *CainiaoReachableBatchjudgeBaseResultDto) SetModule(v CainiaoReachableBatchjudgeReachableServiceWaybillForTopResponseDto) *CainiaoReachableBatchjudgeBaseResultDto {
    s.Module = &v
    return s
}
