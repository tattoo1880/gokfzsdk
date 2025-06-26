package domain


type CainiaoReachableBatchjudgeClientInfoDto struct {
    /*
        调用时自定义描述信息     */
    Description  *string `json:"description,omitempty" `

}

func (s *CainiaoReachableBatchjudgeClientInfoDto) SetDescription(v string) *CainiaoReachableBatchjudgeClientInfoDto {
    s.Description = &v
    return s
}
