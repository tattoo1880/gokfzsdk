package domain


type TaobaoContentMediaUploadPubPubOssFileToMultiMediaResponse struct {
    /*
        文件上传token     */
    UploadToken  *string `json:"upload_token,omitempty" `

}

func (s *TaobaoContentMediaUploadPubPubOssFileToMultiMediaResponse) SetUploadToken(v string) *TaobaoContentMediaUploadPubPubOssFileToMultiMediaResponse {
    s.UploadToken = &v
    return s
}
