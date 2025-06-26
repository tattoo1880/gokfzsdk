package request

import (
        "topsdk/ability316/domain"
        "topsdk/util"
    )

type AlibabaInteractLotteryactivityRegisterRequest struct {
    /*
        入参     */
    ParamTopUpdateActivityLotteryInfoParam  *domain.AlibabaInteractLotteryactivityRegisterTopUpdateActivityLotteryInfoParam `json:"param_top_update_activity_lottery_info_param" required:"true" `
}

func (s *AlibabaInteractLotteryactivityRegisterRequest) SetParamTopUpdateActivityLotteryInfoParam(v domain.AlibabaInteractLotteryactivityRegisterTopUpdateActivityLotteryInfoParam) *AlibabaInteractLotteryactivityRegisterRequest {
    s.ParamTopUpdateActivityLotteryInfoParam = &v
    return s
}

func (req *AlibabaInteractLotteryactivityRegisterRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ParamTopUpdateActivityLotteryInfoParam != nil) {
        paramMap["param_top_update_activity_lottery_info_param"] = util.ConvertStruct(*req.ParamTopUpdateActivityLotteryInfoParam)
    }
    return paramMap
}

func (req *AlibabaInteractLotteryactivityRegisterRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}