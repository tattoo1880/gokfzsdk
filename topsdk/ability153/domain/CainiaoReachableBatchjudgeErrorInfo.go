package domain


type CainiaoReachableBatchjudgeErrorInfo struct {
    /*
        错误信息     */
    ErrorMessage  *string `json:"error_message,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

}

func (s *CainiaoReachableBatchjudgeErrorInfo) SetErrorMessage(v string) *CainiaoReachableBatchjudgeErrorInfo {
    s.ErrorMessage = &v
    return s
}
func (s *CainiaoReachableBatchjudgeErrorInfo) SetErrorCode(v string) *CainiaoReachableBatchjudgeErrorInfo {
    s.ErrorCode = &v
    return s
}
