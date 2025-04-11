package notifier

type notifier struct {
	From     string
	Password string
	Host     string
	Port     string
}

func New(from, password, host, port string) Notifier {
	return &notifier{from, password, host, port}
}
