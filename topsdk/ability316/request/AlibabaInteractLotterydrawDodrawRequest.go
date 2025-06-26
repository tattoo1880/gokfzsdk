package request

import (
        "topsdk/ability316/domain"
        "topsdk/util"
    )

type AlibabaInteractLotterydrawDodrawRequest struct {
    /*
        抽奖请求对象     */
    LotteryDrawQuery  *domain.AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto `json:"lottery_draw_query" required:"true" `
}

func (s *AlibabaInteractLotterydrawDodrawRequest) SetLotteryDrawQuery(v domain.AlibabaInteractLotterydrawDodrawLotteryDrawQueryDto) *AlibabaInteractLotterydrawDodrawRequest {
    s.LotteryDrawQuery = &v
    return s
}

func (req *AlibabaInteractLotterydrawDodrawRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LotteryDrawQuery != nil) {
        paramMap["lottery_draw_query"] = util.ConvertStruct(*req.LotteryDrawQuery)
    }
    return paramMap
}

func (req *AlibabaInteractLotterydrawDodrawRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}