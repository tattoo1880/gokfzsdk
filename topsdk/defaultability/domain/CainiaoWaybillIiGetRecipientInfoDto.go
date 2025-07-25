package domain


type CainiaoWaybillIiGetRecipientInfoDto struct {
    /*
        地址     */
    Address  *CainiaoWaybillIiGetAddressDto `json:"address,omitempty" `

    /*
        手机号码     */
    Mobile  *string `json:"mobile,omitempty" `

    /*
        姓名     */
    Name  *string `json:"name,omitempty" `

    /*
        固定电话     */
    Phone  *string `json:"phone,omitempty" `

    /*
        淘宝订单收件人ID (Open Addressee ID)，长度不超过128个字符，淘宝订单加密情况用于解密。     */
    Oaid  *string `json:"oaid,omitempty" `

    /*
        电商平台真实交易订单号，针对电商平台订单隐私加密场景使用，非必填，如果填写则必须是电商平台真实的交易订单ID     */
    Tid  *string `json:"tid,omitempty" `

    /*
        菜鸟解密地址ID，用于电商平台收件人信息加密的场景使用，非订单加密场景请勿使用。     */
    Caid  *string `json:"caid,omitempty" `

}

func (s *CainiaoWaybillIiGetRecipientInfoDto) SetAddress(v CainiaoWaybillIiGetAddressDto) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Address = &v
    return s
}
func (s *CainiaoWaybillIiGetRecipientInfoDto) SetMobile(v string) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Mobile = &v
    return s
}
func (s *CainiaoWaybillIiGetRecipientInfoDto) SetName(v string) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Name = &v
    return s
}
func (s *CainiaoWaybillIiGetRecipientInfoDto) SetPhone(v string) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Phone = &v
    return s
}
func (s *CainiaoWaybillIiGetRecipientInfoDto) SetOaid(v string) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Oaid = &v
    return s
}
func (s *CainiaoWaybillIiGetRecipientInfoDto) SetTid(v string) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Tid = &v
    return s
}
func (s *CainiaoWaybillIiGetRecipientInfoDto) SetCaid(v string) *CainiaoWaybillIiGetRecipientInfoDto {
    s.Caid = &v
    return s
}
