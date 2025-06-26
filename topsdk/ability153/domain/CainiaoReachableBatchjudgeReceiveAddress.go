package domain


type CainiaoReachableBatchjudgeReceiveAddress struct {
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

func (s *CainiaoReachableBatchjudgeReceiveAddress) SetTownName(v string) *CainiaoReachableBatchjudgeReceiveAddress {
    s.TownName = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReceiveAddress) SetAddressDetail(v string) *CainiaoReachableBatchjudgeReceiveAddress {
    s.AddressDetail = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReceiveAddress) SetCityName(v string) *CainiaoReachableBatchjudgeReceiveAddress {
    s.CityName = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReceiveAddress) SetAreaName(v string) *CainiaoReachableBatchjudgeReceiveAddress {
    s.AreaName = &v
    return s
}
func (s *CainiaoReachableBatchjudgeReceiveAddress) SetProvinceName(v string) *CainiaoReachableBatchjudgeReceiveAddress {
    s.ProvinceName = &v
    return s
}
