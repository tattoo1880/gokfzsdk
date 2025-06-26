package request


type TaobaoCrmServiceChannelShortlinkCreateRequest struct {
    /*
        淘短链类型：LT_ITEM（商品淘短链）,LT_SHOP（店铺首页淘短链）,LT_ACTIVITY（活动页淘短链）,LT_TRADE（订单详情页淘短链）。     */
    LinkType  *string `json:"link_type" required:"true" `
    /*
        类型为LT_ITEM时必须传入商品ID，类型为LT_SHOP时必须传入空值，类型为LT_ACTIVITY时传入活动页URL（URL必须是taobao.com,tmall.com,jaeapp.com这三个域名下的URL），类型为LT_TRADE时传入订单ID。     */
    ShortLinkData  *string `json:"short_link_data,omitempty" required:"false" `
    /*
        淘短链名称（最多只能16个中文字符，类型为订单链接时传入订单ID）。     */
    ShortLinkName  *string `json:"short_link_name" required:"true" `
}

func (s *TaobaoCrmServiceChannelShortlinkCreateRequest) SetLinkType(v string) *TaobaoCrmServiceChannelShortlinkCreateRequest {
    s.LinkType = &v
    return s
}
func (s *TaobaoCrmServiceChannelShortlinkCreateRequest) SetShortLinkData(v string) *TaobaoCrmServiceChannelShortlinkCreateRequest {
    s.ShortLinkData = &v
    return s
}
func (s *TaobaoCrmServiceChannelShortlinkCreateRequest) SetShortLinkName(v string) *TaobaoCrmServiceChannelShortlinkCreateRequest {
    s.ShortLinkName = &v
    return s
}

func (req *TaobaoCrmServiceChannelShortlinkCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LinkType != nil) {
        paramMap["link_type"] = *req.LinkType
    }
    if(req.ShortLinkData != nil) {
        paramMap["short_link_data"] = *req.ShortLinkData
    }
    if(req.ShortLinkName != nil) {
        paramMap["short_link_name"] = *req.ShortLinkName
    }
    return paramMap
}

func (req *TaobaoCrmServiceChannelShortlinkCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}