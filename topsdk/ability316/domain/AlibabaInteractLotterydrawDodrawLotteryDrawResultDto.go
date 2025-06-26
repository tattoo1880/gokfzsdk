package domain


type AlibabaInteractLotterydrawDodrawLotteryDrawResultDto struct {
    /*
        resultType     */
    ResultType  *int64 `json:"result_type,omitempty" `

    /*
        resultMsg     */
    ResultMsg  *string `json:"result_msg,omitempty" `

    /*
        shopTitle     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        shopLogo     */
    ShopLogo  *string `json:"shop_logo,omitempty" `

    /*
        template     */
    Template  *string `json:"template,omitempty" `

    /*
        award     */
    Award  *AlibabaInteractLotterydrawDodrawLotteryAwardDto `json:"award,omitempty" `

}

func (s *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) SetResultType(v int64) *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto {
    s.ResultType = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) SetResultMsg(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto {
    s.ResultMsg = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) SetShopTitle(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto {
    s.ShopTitle = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) SetShopLogo(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto {
    s.ShopLogo = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) SetTemplate(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto {
    s.Template = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto) SetAward(v AlibabaInteractLotterydrawDodrawLotteryAwardDto) *AlibabaInteractLotterydrawDodrawLotteryDrawResultDto {
    s.Award = &v
    return s
}
