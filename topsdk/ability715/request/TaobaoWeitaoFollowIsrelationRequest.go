package request


type TaobaoWeitaoFollowIsrelationRequest struct {
    /*
        要查询的粉丝的淘宝昵称     */
    FansNick  *string `json:"fans_nick,omitempty" required:"false" `
    /*
        ouid     */
    Ouid  *string `json:"ouid,omitempty" required:"false" `
    /*
        openid     */
    Openid  *string `json:"openid,omitempty" required:"false" `
}

func (s *TaobaoWeitaoFollowIsrelationRequest) SetFansNick(v string) *TaobaoWeitaoFollowIsrelationRequest {
    s.FansNick = &v
    return s
}
func (s *TaobaoWeitaoFollowIsrelationRequest) SetOuid(v string) *TaobaoWeitaoFollowIsrelationRequest {
    s.Ouid = &v
    return s
}
func (s *TaobaoWeitaoFollowIsrelationRequest) SetOpenid(v string) *TaobaoWeitaoFollowIsrelationRequest {
    s.Openid = &v
    return s
}

func (req *TaobaoWeitaoFollowIsrelationRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.FansNick != nil) {
        paramMap["fans_nick"] = *req.FansNick
    }
    if(req.Ouid != nil) {
        paramMap["ouid"] = *req.Ouid
    }
    if(req.Openid != nil) {
        paramMap["openid"] = *req.Openid
    }
    return paramMap
}

func (req *TaobaoWeitaoFollowIsrelationRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}