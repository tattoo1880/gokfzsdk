package domain


type TaobaoWirelessPictureCheckCheckpoints struct {
    /*
        检查的场景。porn：涉黄；terrorism：暴恐涉政，     */
    Scene  *string `json:"scene,omitempty" `

    /*
        检查的场景的标签。在porn场景，值为normal：正常图片，无色情内容；sexy：性感图片；porn：色情图片。 在terrorism场景，为normal：正常图片；bloody：血腥；explosion：爆炸烟光；outfit：特殊装束；logo：特殊标识；weapon：武器；politics：涉政；violence ： 打斗；crowd：聚众；parade：游行；carcrash：车祸现场；flag：旗帜；location：地标     */
    Label  *string `json:"label,omitempty" `

    /*
        结果建议     */
    Suggestion  *string `json:"suggestion,omitempty" `

    /*
        结果准确度     */
    Rate  *string `json:"rate,omitempty" `

}

func (s *TaobaoWirelessPictureCheckCheckpoints) SetScene(v string) *TaobaoWirelessPictureCheckCheckpoints {
    s.Scene = &v
    return s
}
func (s *TaobaoWirelessPictureCheckCheckpoints) SetLabel(v string) *TaobaoWirelessPictureCheckCheckpoints {
    s.Label = &v
    return s
}
func (s *TaobaoWirelessPictureCheckCheckpoints) SetSuggestion(v string) *TaobaoWirelessPictureCheckCheckpoints {
    s.Suggestion = &v
    return s
}
func (s *TaobaoWirelessPictureCheckCheckpoints) SetRate(v string) *TaobaoWirelessPictureCheckCheckpoints {
    s.Rate = &v
    return s
}
