package domain


type AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO struct {
    /*
        上游商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        上游商品最高价(分)     */
    ItemHighestPrice  *int64 `json:"item_highest_price,omitempty" `

    /*
        上游商品最低价(分)     */
    ItemLowestPrice  *int64 `json:"item_lowest_price,omitempty" `

    /*
        上游商品店铺ID     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        上游商品来源     */
    Type  *string `json:"type,omitempty" `

    /*
        上游商品URL     */
    ItemUrl  *string `json:"item_url,omitempty" `

}

func (s *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) SetItemId(v string) *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO {
    s.ItemId = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) SetItemHighestPrice(v int64) *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO {
    s.ItemHighestPrice = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) SetItemLowestPrice(v int64) *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO {
    s.ItemLowestPrice = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) SetShopId(v string) *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO {
    s.ShopId = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) SetType(v string) *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO {
    s.Type = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) SetItemUrl(v string) *AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO {
    s.ItemUrl = &v
    return s
}
