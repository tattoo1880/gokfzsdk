package domain


type CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto struct {
    /*
        单个结果是否异常     */
    ErrorInfo  *CainiaoReachableBatchjudgeErrorInfo `json:"error_info,omitempty" `

    /*
        单个结果是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        可达信息     */
    Module  *CainiaoReachableBatchjudgeReachableDto `json:"module,omitempty" `

    /*
        与入参地址列表中单项objectId对应     */
    ObjectId  *string `json:"object_id,omitempty" `

}

func (s *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto) SetErrorInfo(v CainiaoReachableBatchjudgeErrorInfo) *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto {
    s.ErrorInfo = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto) SetSuccess(v bool) *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto {
    s.Success = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto) SetModule(v CainiaoReachableBatchjudgeReachableDto) *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto {
    s.Module = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto) SetObjectId(v string) *CainiaoReachableBatchjudgeReachableServiceWaybillResponseDto {
    s.ObjectId = &v
    return s
}
