package domain


type AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto struct {
    /*
        抽奖CODE     */
    RaffleCode  *string `json:"raffle_code,omitempty" `

    /*
        前台appKey     */
    RaffleIdentity  *string `json:"raffle_identity,omitempty" `

    /*
        方案ID     */
    SchemaId  *int64 `json:"schema_id,omitempty" `

    /*
        关联ID     */
    RelationId  *string `json:"relation_id,omitempty" `

    /*
        抽奖参数     */
    Wua  *string `json:"wua,omitempty" `

}

func (s *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto) SetRaffleCode(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto {
    s.RaffleCode = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto) SetRaffleIdentity(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto {
    s.RaffleIdentity = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto) SetSchemaId(v int64) *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto {
    s.SchemaId = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto) SetRelationId(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto {
    s.RelationId = &v
    return s
}
func (s *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto) SetWua(v string) *AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto {
    s.Wua = &v
    return s
}
