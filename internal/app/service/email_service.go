// service/email_service.go

package service

type EmailService struct {
    // メール送信に必要な設定やクライアントなどを含む
}

func NewEmailService() *EmailService {
    // 初期化処理...
    return &EmailService{}
}

func (es *EmailService) SendEmail(to, subject, body string) error {
    // メール送信の実装...
}
