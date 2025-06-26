package domain

import (
        "topsdk/util"
    )

type AlibabaInteractLotterydrawDodrawLotteryAwardDto struct {
    /*
        totalResCount     */
    TotalResCount  *int64 `json:"total_res_count,omitempty" `

    /*
        awardDetailUrl     */
    AwardDetailUrl  *string `json:"award_detail_url,omitempty" `

    /*
        useEndDate     */
    UseEndDate  *util.LocalTime `json:"use_end_date,omitempty" `

    /*
        useStartDate     */
    UseStartDate  *util.LocalTime `json:"use_start_date,omitempty" `

    /*
        endDate     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" `

    /*
        startDate     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" `

    /*
        币种     */
    Currency  *string `json:"currency,omitempty" `

    /*
        awardPrice     */
    AwardPrice  *int64 `json:"award_price,omitempty" `

    /*
        startFee     */
    StartFee  *int64 `json:"start_fee,omitempty" `

    /*
        pictUrl     */
    PictUrl  *string `json:"pict_url,omitempty" `

    /*
        displayName     */
    DisplayName  *string `json:"display_name,omitempty" `

    /*
        name     */
    Name  *string `json:"name,omitempty" `

    /*
        awardInstId     */
    AwardInstId  *string `json:"award_inst_id,omitempty" `

    /*
        awardTypeName     */
    AwardTypeName  *string `json:"award_type_name,omitempty" `

    /*
        awardType     */
    AwardType  *int64 `json:"award_type,omitempty" `

    /*
        canUseResCount     */
    CanUseResCount  *int64 `json:"can_use_res_count,omitempty" `

    /*
        condition     */
    Condition  *string `json:"condition,omitempty" `

    /*
        币种符号     */
    CurrencySign  *string `json:"currency_sign,omitempty" `

    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        lotteryActivityId     */
    LotteryActivityId  *int64 `json:"lottery_activity_id,omitempty" `

}

func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetTotalResCount(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.TotalResCount = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetAwardDetailUrl(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.AwardDetailUrl = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetUseEndDate(v util.LocalTime) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.UseEndDate = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetUseStartDate(v util.LocalTime) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.UseStartDate = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetEndDate(v util.LocalTime) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.EndDate = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetStartDate(v util.LocalTime) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.StartDate = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetCurrency(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.Currency = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetAwardPrice(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.AwardPrice = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetStartFee(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.StartFee = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetPictUrl(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.PictUrl = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetDisplayName(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.DisplayName = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetName(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.Name = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetAwardInstId(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.AwardInstId = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetAwardTypeName(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.AwardTypeName = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetAwardType(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.AwardType = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetCanUseResCount(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.CanUseResCount = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetCondition(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.Condition = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetCurrencySign(v string) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.CurrencySign = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetId(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.Id = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryAwardDto) SetLotteryActivityId(v int64) *AlibabaInteractLotterydrawDodrawLotteryAwardDto {
    s.LotteryActivityId = &v
    return s
}
