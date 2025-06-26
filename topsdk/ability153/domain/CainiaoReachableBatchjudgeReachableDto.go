package domain


type CainiaoReachableBatchjudgeReachableDto struct {
    /*
        阻断原因     */
    InterruptReason  *string `json:"interrupt_reason,omitempty" `

    /*
        是否阻断 true:阻断  false:可达     */
    InterruptApplyWaybillCode  *bool `json:"interrupt_apply_waybill_code,omitempty" `

}

func (s *CainiaoReachableBatchjudgeReachableDto) SetInterruptReason(v string) *CainiaoReachableBatchjudgeReachableDto {
    s.InterruptReason = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReachableDto) SetInterruptApplyWaybillCode(v bool) *CainiaoReachableBatchjudgeReachableDto {
    s.InterruptApplyWaybillCode = &v
    return s
}
