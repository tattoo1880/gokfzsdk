package domain


type AlibabaInteractOpenIsattentionResultDo struct {
    /*
        isRetry     */
    IsRetry  *bool `json:"is_retry,omitempty" `

    /*
        data     */
    Data  *int64 `json:"data,omitempty" `

    /*
        code     */
    Code  *string `json:"code,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        msg     */
    Msg  *string `json:"msg,omitempty" `

}

func (s *AlibabaInteractOpenIsattentionResultDo) SetIsRetry(v bool) *AlibabaInteractOpenIsattentionResultDo {
    s.IsRetry = &v
    return s
}
func (s *AlibabaInteractOpenIsattentionResultDo) SetData(v int64) *AlibabaInteractOpenIsattentionResultDo {
    s.Data = &v
    return s
}
func (s *AlibabaInteractOpenIsattentionResultDo) SetCode(v string) *AlibabaInteractOpenIsattentionResultDo {
    s.Code = &v
    return s
}
func (s *AlibabaInteractOpenIsattentionResultDo) SetSuccess(v bool) *AlibabaInteractOpenIsattentionResultDo {
    s.Success = &v
    return s
}
func (s *AlibabaInteractOpenIsattentionResultDo) SetMsg(v string) *AlibabaInteractOpenIsattentionResultDo {
    s.Msg = &v
    return s
}
