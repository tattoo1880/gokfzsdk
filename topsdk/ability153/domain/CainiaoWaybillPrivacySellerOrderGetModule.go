package domain


type CainiaoWaybillPrivacySellerOrderGetModule struct {
    /*
        商家ID     */
    SellerId  *string `json:"seller_id,omitempty" `

    /*
        隐私次数     */
    PrivacyCount  *int64 `json:"privacy_count,omitempty" `

    /*
        日期     */
    OrderDate  *string `json:"order_date,omitempty" `

    /*
        订单渠道     */
    OrderChannel  *string `json:"order_channel,omitempty" `

    /*
        店铺id     */
    ShopId  *string `json:"shop_id,omitempty" `

}

func (s *CainiaoWaybillPrivacySellerOrderGetModule) SetSellerId(v string) *CainiaoWaybillPrivacySellerOrderGetModule {
    s.SellerId = &v
    return s
}
func (s *CainiaoWaybillPrivacySellerOrderGetModule) SetPrivacyCount(v int64) *CainiaoWaybillPrivacySellerOrderGetModule {
    s.PrivacyCount = &v
    return s
}
func (s *CainiaoWaybillPrivacySellerOrderGetModule) SetOrderDate(v string) *CainiaoWaybillPrivacySellerOrderGetModule {
    s.OrderDate = &v
    return s
}
func (s *CainiaoWaybillPrivacySellerOrderGetModule) SetOrderChannel(v string) *CainiaoWaybillPrivacySellerOrderGetModule {
    s.OrderChannel = &v
    return s
}
func (s *CainiaoWaybillPrivacySellerOrderGetModule) SetShopId(v string) *CainiaoWaybillPrivacySellerOrderGetModule {
    s.ShopId = &v
    return s
}
