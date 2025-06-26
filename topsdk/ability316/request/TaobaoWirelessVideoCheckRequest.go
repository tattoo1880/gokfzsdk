package request


type TaobaoWirelessVideoCheckRequest struct {
    /*
        视频截帧间隔，取值范围为[1, 60]，单位为秒。 截帧最多张数为200张，请根据视频时长，合理设置截帧间隔。     */
    Interval  *int64 `json:"interval" required:"true" `
    /*
        视频的URL，必须为淘系安全域名地址。视频格式支持flv、mp4。     */
    Url  *string `json:"url" required:"true" `
}

func (s *TaobaoWirelessVideoCheckRequest) SetInterval(v int64) *TaobaoWirelessVideoCheckRequest {
    s.Interval = &v
    return s
}
func (s *TaobaoWirelessVideoCheckRequest) SetUrl(v string) *TaobaoWirelessVideoCheckRequest {
    s.Url = &v
    return s
}

func (req *TaobaoWirelessVideoCheckRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Interval != nil) {
        paramMap["interval"] = *req.Interval
    }
    if(req.Url != nil) {
        paramMap["url"] = *req.Url
    }
    return paramMap
}

func (req *TaobaoWirelessVideoCheckRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}