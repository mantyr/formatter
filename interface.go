package formatter

type Formatter interface {
	SetFormat(Format) error
	SetHeader(Header) error
	Write(interface{}) error
	Flush() error
}
