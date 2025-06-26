package request


type AlibabaInteractUserIsloginRequest struct {
    /*
        用户nick     */
    BuyerNick  *string `json:"buyer_nick" required:"true" `
}

func (s *AlibabaInteractUserIsloginRequest) SetBuyerNick(v string) *AlibabaInteractUserIsloginRequest {
    s.BuyerNick = &v
    return s
}

func (req *AlibabaInteractUserIsloginRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BuyerNick != nil) {
        paramMap["buyer_nick"] = *req.BuyerNick
    }
    return paramMap
}

func (req *AlibabaInteractUserIsloginRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}