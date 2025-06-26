package domain


type TaobaoRegionPriceQueryRegionalPriceDto struct {
    /*
        市     */
    City  *string `json:"city,omitempty" `

    /*
        区县，特殊可选     */
    District  *string `json:"district,omitempty" `

    /*
        省     */
    Province  *string `json:"province,omitempty" `

    /*
        街道，特殊可选     */
    Street  *string `json:"street,omitempty" `

}

func (s *TaobaoRegionPriceQueryRegionalPriceDto) SetCity(v string) *TaobaoRegionPriceQueryRegionalPriceDto {
    s.City = &v
    return s
}
func (s *TaobaoRegionPriceQueryRegionalPriceDto) SetDistrict(v string) *TaobaoRegionPriceQueryRegionalPriceDto {
    s.District = &v
    return s
}
func (s *TaobaoRegionPriceQueryRegionalPriceDto) SetProvince(v string) *TaobaoRegionPriceQueryRegionalPriceDto {
    s.Province = &v
    return s
}
func (s *TaobaoRegionPriceQueryRegionalPriceDto) SetStreet(v string) *TaobaoRegionPriceQueryRegionalPriceDto {
    s.Street = &v
    return s
}
