package request


type TaobaoWirelessPictureCheckRequest struct {
    /*
        图片的URL,URL必须为淘系安全域名地址。图片格式支持png,jpg,webp     */
    Url  *string `json:"url" required:"true" `
}

func (s *TaobaoWirelessPictureCheckRequest) SetUrl(v string) *TaobaoWirelessPictureCheckRequest {
    s.Url = &v
    return s
}

func (req *TaobaoWirelessPictureCheckRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Url != nil) {
        paramMap["url"] = *req.Url
    }
    return paramMap
}

func (req *TaobaoWirelessPictureCheckRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}