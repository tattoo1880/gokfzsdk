package domain


type AlibabaInteractLotteryactivityRegisterResult struct {
    /*
        返回的数据     */
    Data  *AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult `json:"data,omitempty" `

    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误信息     */
    ErrMsg  *string `json:"err_msg,omitempty" `

    /*
        xxx     */
    ModuleMap  *AlibabaInteractLotteryactivityRegisterModulemap `json:"module_map,omitempty" `

    /*
        注册抽奖活动失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        方便追踪请求的唯一标识     */
    TraceId  *string `json:"trace_id,omitempty" `

}

func (s *AlibabaInteractLotteryactivityRegisterResult) SetData(v AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult) *AlibabaInteractLotteryactivityRegisterResult {
    s.Data = &v
    return s
}
func (s *AlibabaInteractLotteryactivityRegisterResult) SetErrCode(v string) *AlibabaInteractLotteryactivityRegisterResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaInteractLotteryactivityRegisterResult) SetErrMsg(v string) *AlibabaInteractLotteryactivityRegisterResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaInteractLotteryactivityRegisterResult) SetModuleMap(v AlibabaInteractLotteryactivityRegisterModulemap) *AlibabaInteractLotteryactivityRegisterResult {
    s.ModuleMap = &v
    return s
}
func (s *AlibabaInteractLotteryactivityRegisterResult) SetSuccess(v bool) *AlibabaInteractLotteryactivityRegisterResult {
    s.Success = &v
    return s
}
func (s *AlibabaInteractLotteryactivityRegisterResult) SetTraceId(v string) *AlibabaInteractLotteryactivityRegisterResult {
    s.TraceId = &v
    return s
}
