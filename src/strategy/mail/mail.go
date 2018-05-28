package mail

const (
	SMTP = "smtp.qq.com"
	PORT = 465
)

func NewEmail(config *Config) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", config.To)
	m.SetHeader("Subject", config.Subject)
	m.SetBody("text/html", config.Text)
	d := gomail.NewDialer(SMTP, PORT, config.From, config.PassWrod)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

type Config struct {
	From     string
	To       string
	Subject  string
	Text     string
	PassWrod string
}
