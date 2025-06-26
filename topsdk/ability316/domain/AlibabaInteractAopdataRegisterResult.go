package domain


type AlibabaInteractAopdataRegisterResult struct {
    /*
        xx     */
    Data  *string `json:"data,omitempty" `

    /*
        xxx失败     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        xxx失败     */
    ErrMsg  *string `json:"err_msg,omitempty" `

    /*
        接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        跟踪请求使用     */
    TraceId  *string `json:"trace_id,omitempty" `

}

func (s *AlibabaInteractAopdataRegisterResult) SetData(v string) *AlibabaInteractAopdataRegisterResult {
    s.Data = &v
    return s
}
func (s *AlibabaInteractAopdataRegisterResult) SetErrCode(v string) *AlibabaInteractAopdataRegisterResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaInteractAopdataRegisterResult) SetErrMsg(v string) *AlibabaInteractAopdataRegisterResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaInteractAopdataRegisterResult) SetSuccess(v bool) *AlibabaInteractAopdataRegisterResult {
    s.Success = &v
    return s
}
func (s *AlibabaInteractAopdataRegisterResult) SetTraceId(v string) *AlibabaInteractAopdataRegisterResult {
    s.TraceId = &v
    return s
}
