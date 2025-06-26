package domain


type CainiaoReachableBatchjudgeAddressDto struct {
    /*
        镇/街道     */
    TownName  *string `json:"town_name,omitempty" `

    /*
        详细地址     */
    AddressDetail  *string `json:"address_detail,omitempty" `

    /*
        市     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        区     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        省     */
    ProvinceName  *string `json:"province_name,omitempty" `

}

func (s *CainiaoReachableBatchjudgeAddressDto) SetTownName(v string) *CainiaoReachableBatchjudgeAddressDto {
    s.TownName = &v
    return s
}
func (s *CainiaoReachableBatchjudgeAddressDto) SetAddressDetail(v string) *CainiaoReachableBatchjudgeAddressDto {
    s.AddressDetail = &v
    return s
}
func (s *CainiaoReachableBatchjudgeAddressDto) SetCityName(v string) *CainiaoReachableBatchjudgeAddressDto {
    s.CityName = &v
    return s
}
func (s *CainiaoReachableBatchjudgeAddressDto) SetAreaName(v string) *CainiaoReachableBatchjudgeAddressDto {
    s.AreaName = &v
    return s
}
func (s *CainiaoReachableBatchjudgeAddressDto) SetProvinceName(v string) *CainiaoReachableBatchjudgeAddressDto {
    s.ProvinceName = &v
    return s
}
