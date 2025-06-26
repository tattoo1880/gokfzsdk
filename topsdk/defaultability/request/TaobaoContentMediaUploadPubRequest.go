package request


type TaobaoContentMediaUploadPubRequest struct {
    /*
        视频上传时对应的渠道业务编码,发布关注(订阅):out_tao_dingyue,   发布视频空间:out_tao_mv     */
    BizCode  *string `json:"biz_code,omitempty" required:"false" `
    /*
        视频上传oss时的文件名（文件名在上传oss时需要uuid生成，做到不重复）     */
    UuidFileName  *string `json:"uuid_file_name,omitempty" required:"false" `
    /*
        文件的类型     */
    MimeType  *string `json:"mime_type,omitempty" required:"false" `
}

func (s *TaobaoContentMediaUploadPubRequest) SetBizCode(v string) *TaobaoContentMediaUploadPubRequest {
    s.BizCode = &v
    return s
}
func (s *TaobaoContentMediaUploadPubRequest) SetUuidFileName(v string) *TaobaoContentMediaUploadPubRequest {
    s.UuidFileName = &v
    return s
}
func (s *TaobaoContentMediaUploadPubRequest) SetMimeType(v string) *TaobaoContentMediaUploadPubRequest {
    s.MimeType = &v
    return s
}

func (req *TaobaoContentMediaUploadPubRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizCode != nil) {
        paramMap["biz_code"] = *req.BizCode
    }
    if(req.UuidFileName != nil) {
        paramMap["uuid_file_name"] = *req.UuidFileName
    }
    if(req.MimeType != nil) {
        paramMap["mime_type"] = *req.MimeType
    }
    return paramMap
}

func (req *TaobaoContentMediaUploadPubRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}