package request

import (
        "topsdk/ability316/domain"
        "topsdk/util"
    )

type AlibabaInteractAopdataRegisterRequest struct {
    /*
        入参     */
    ParamTopIsvDecorateParam  *domain.AlibabaInteractAopdataRegisterTopIsvDecorateParam `json:"param_top_isv_decorate_param" required:"true" `
}

func (s *AlibabaInteractAopdataRegisterRequest) SetParamTopIsvDecorateParam(v domain.AlibabaInteractAopdataRegisterTopIsvDecorateParam) *AlibabaInteractAopdataRegisterRequest {
    s.ParamTopIsvDecorateParam = &v
    return s
}

func (req *AlibabaInteractAopdataRegisterRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ParamTopIsvDecorateParam != nil) {
        paramMap["param_top_isv_decorate_param"] = util.ConvertStruct(*req.ParamTopIsvDecorateParam)
    }
    return paramMap
}

func (req *AlibabaInteractAopdataRegisterRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}