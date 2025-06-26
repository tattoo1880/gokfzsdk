package domain


type TaobaoRegionPriceManageRegionalPriceDto struct {
    /*
        市     */
    City  *string `json:"city,omitempty" `

    /*
        金额（分）     */
    Price  *int64 `json:"price,omitempty" `

    /*
        省     */
    Province  *string `json:"province,omitempty" `

}

func (s *TaobaoRegionPriceManageRegionalPriceDto) SetCity(v string) *TaobaoRegionPriceManageRegionalPriceDto {
    s.City = &v
    return s
}
func (s *TaobaoRegionPriceManageRegionalPriceDto) SetPrice(v int64) *TaobaoRegionPriceManageRegionalPriceDto {
    s.Price = &v
    return s
}
func (s *TaobaoRegionPriceManageRegionalPriceDto) SetProvince(v string) *TaobaoRegionPriceManageRegionalPriceDto {
    s.Province = &v
    return s
}
