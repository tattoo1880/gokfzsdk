package domain


type CainiaoWaybillPrivacySubscriptionGetResult struct {
    /*
        错误code列表     */
    ErrorCodeList  *[]string `json:"error_code_list,omitempty" `

    /*
        是否失败     */
    Failure  *bool `json:"failure,omitempty" `

    /*
        第一个错误     */
    OneErrorInfo  *string `json:"one_error_info,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        商家是否订购     */
    Subscription  *bool `json:"subscription,omitempty" `

    /*
        系统自动生成     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误列表     */
    ErrorInfoList  *[]string `json:"error_info_list,omitempty" `

    /*
        系统信息     */
    ObjectId  *string `json:"object_id,omitempty" `

}

func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetErrorCodeList(v []string) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.ErrorCodeList = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetFailure(v bool) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.Failure = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetOneErrorInfo(v string) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.OneErrorInfo = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetSuccess(v bool) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.Success = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetSubscription(v bool) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.Subscription = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetErrorMessage(v string) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.ErrorMessage = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetErrorCode(v string) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.ErrorCode = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetErrorInfoList(v []string) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.ErrorInfoList = &v
    return s
}
func (s *CainiaoWaybillPrivacySubscriptionGetResult) SetObjectId(v string) *CainiaoWaybillPrivacySubscriptionGetResult {
    s.ObjectId = &v
    return s
}
