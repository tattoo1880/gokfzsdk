package domain


type TaobaoWirelessContentCheckCheckpoints struct {
    /*
        检查的场景。antispam为黄暴政     */
    Scene  *string `json:"scene,omitempty" `

    /*
        检查的场景标签。normal：正常文本 spam：含垃圾信息；ad：广告；politics：涉政；terrorism：暴恐；abuse：辱骂；porn：色情；flood：灌水；contraband：违禁；meaningless：无意义     */
    Label  *string `json:"label,omitempty" `

    /*
        结果建议     */
    Suggestion  *string `json:"suggestion,omitempty" `

    /*
        结果准确度     */
    Rate  *string `json:"rate,omitempty" `

}

func (s *TaobaoWirelessContentCheckCheckpoints) SetScene(v string) *TaobaoWirelessContentCheckCheckpoints {
    s.Scene = &v
    return s
}
func (s *TaobaoWirelessContentCheckCheckpoints) SetLabel(v string) *TaobaoWirelessContentCheckCheckpoints {
    s.Label = &v
    return s
}
func (s *TaobaoWirelessContentCheckCheckpoints) SetSuggestion(v string) *TaobaoWirelessContentCheckCheckpoints {
    s.Suggestion = &v
    return s
}
func (s *TaobaoWirelessContentCheckCheckpoints) SetRate(v string) *TaobaoWirelessContentCheckCheckpoints {
    s.Rate = &v
    return s
}
