package domain


type AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult struct {
    /*
        isv活动的id     */
    ActivityId  *int64 `json:"activity_id,omitempty" `

    /*
        isv活动url     */
    H5Url  *string `json:"h5_url,omitempty" `

}

func (s *AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult) SetActivityId(v int64) *AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult {
    s.ActivityId = &v
    return s
}
func (s *AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult) SetH5Url(v string) *AlibabaInteractLotteryactivityRegisterActivityLotteryWriteResult {
    s.H5Url = &v
    return s
}
