package domain


type AlibabaInteractLotterydrawDodrawResultDto struct {
    /*
        result     */
    LotteryDrawResult  *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto `json:"lottery_draw_result,omitempty" `

    /*
        code     */
    Code  *int64 `json:"code,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        msg     */
    Msg  *string `json:"msg,omitempty" `

}

func (s *AlibabaInteractLotterydrawDodrawResultDto) SetLotteryDrawResult(v AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) *AlibabaInteractLotterydrawDodrawResultDto {
    s.LotteryDrawResult = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawResultDto) SetCode(v int64) *AlibabaInteractLotterydrawDodrawResultDto {
    s.Code = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawResultDto) SetSuccess(v bool) *AlibabaInteractLotterydrawDodrawResultDto {
    s.Success = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawResultDto) SetMsg(v string) *AlibabaInteractLotterydrawDodrawResultDto {
    s.Msg = &v
    return s
}
