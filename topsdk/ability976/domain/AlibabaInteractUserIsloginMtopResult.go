package domain


type AlibabaInteractUserIsloginMtopResult struct {
    /*
        model     */
    Model  *string `json:"model,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaInteractUserIsloginMtopResult) SetModel(v string) *AlibabaInteractUserIsloginMtopResult {
    s.Model = &v
    return s
}
func (s *AlibabaInteractUserIsloginMtopResult) SetMsgCode(v string) *AlibabaInteractUserIsloginMtopResult {
    s.MsgCode = &v
    return s
}
func (s *AlibabaInteractUserIsloginMtopResult) SetMsgInfo(v string) *AlibabaInteractUserIsloginMtopResult {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaInteractUserIsloginMtopResult) SetSuccess(v bool) *AlibabaInteractUserIsloginMtopResult {
    s.Success = &v
    return s
}
