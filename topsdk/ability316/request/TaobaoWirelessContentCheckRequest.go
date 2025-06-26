package request


type TaobaoWirelessContentCheckRequest struct {
    /*
        待检查的文本     */
    Text  *string `json:"text" required:"true" `
}

func (s *TaobaoWirelessContentCheckRequest) SetText(v string) *TaobaoWirelessContentCheckRequest {
    s.Text = &v
    return s
}

func (req *TaobaoWirelessContentCheckRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Text != nil) {
        paramMap["text"] = *req.Text
    }
    return paramMap
}

func (req *TaobaoWirelessContentCheckRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}