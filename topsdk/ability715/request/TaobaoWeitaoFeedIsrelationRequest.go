package request


type TaobaoWeitaoFeedIsrelationRequest struct {
    /*
        要查询的公共账号的淘宝昵称     */
    SellerNick  *string `json:"seller_nick" required:"true" `
    /*
        要查询的粉丝的淘宝昵称     */
    FansNick  *string `json:"fans_nick" required:"true" `
}

func (s *TaobaoWeitaoFeedIsrelationRequest) SetSellerNick(v string) *TaobaoWeitaoFeedIsrelationRequest {
    s.SellerNick = &v
    return s
}
func (s *TaobaoWeitaoFeedIsrelationRequest) SetFansNick(v string) *TaobaoWeitaoFeedIsrelationRequest {
    s.FansNick = &v
    return s
}

func (req *TaobaoWeitaoFeedIsrelationRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SellerNick != nil) {
        paramMap["seller_nick"] = *req.SellerNick
    }
    if(req.FansNick != nil) {
        paramMap["fans_nick"] = *req.FansNick
    }
    return paramMap
}

func (req *TaobaoWeitaoFeedIsrelationRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}