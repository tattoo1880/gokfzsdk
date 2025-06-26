package request


type AlibabaInteractOpenIsattentionRequest struct {
    /*
        1 defalutValue��1    */
    Param0  *int64 `json:"param0,omitempty" required:"false" `
}

func (s *AlibabaInteractOpenIsattentionRequest) SetParam0(v int64) *AlibabaInteractOpenIsattentionRequest {
    s.Param0 = &v
    return s
}

func (req *AlibabaInteractOpenIsattentionRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Param0 != nil) {
        paramMap["param0"] = *req.Param0
    }
    return paramMap
}

func (req *AlibabaInteractOpenIsattentionRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}