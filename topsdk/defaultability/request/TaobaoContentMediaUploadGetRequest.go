package request


type TaobaoContentMediaUploadGetRequest struct {
    /*
        文件上传token     */
    UploadToken  *string `json:"upload_token,omitempty" required:"false" `
}

func (s *TaobaoContentMediaUploadGetRequest) SetUploadToken(v string) *TaobaoContentMediaUploadGetRequest {
    s.UploadToken = &v
    return s
}

func (req *TaobaoContentMediaUploadGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UploadToken != nil) {
        paramMap["upload_token"] = *req.UploadToken
    }
    return paramMap
}

func (req *TaobaoContentMediaUploadGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}