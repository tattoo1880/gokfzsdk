package request


type TaobaoContentMediaUploadSecretHzRequest struct {
    /*
        业务线,需和对应业务对应上 否则平台读取不到文件     */
    BizLine  *string `json:"biz_line,omitempty" required:"false" `
}

func (s *TaobaoContentMediaUploadSecretHzRequest) SetBizLine(v string) *TaobaoContentMediaUploadSecretHzRequest {
    s.BizLine = &v
    return s
}

func (req *TaobaoContentMediaUploadSecretHzRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizLine != nil) {
        paramMap["biz_line"] = *req.BizLine
    }
    return paramMap
}

func (req *TaobaoContentMediaUploadSecretHzRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}