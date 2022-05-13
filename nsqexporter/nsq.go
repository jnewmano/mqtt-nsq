package nsqexporter

/*
type Handler interface {
	HandleMessage(message *nsq.Message) error
}

func NewConsumer(h Handler, topic string, channel string, lookupdHTTPAddrs []string, cfg *nsq.Config) (*nsq.Consumer, error) {

	if cfg == nil {
		cfg = nsq.NewConfig()
		cfg.MaxAttempts = 5
		cfg.MaxBackoffDuration = time.Second * 10
		cfg.MaxInFlight = 20
	}

	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return nil, err
	}

	err = c.AddConcurrentHandlers(handler, concurrency)
	if err != nil {
		return nil, err
	}

	err = c.ConnectToNSQLookupds(lookupdHTTPAddrs)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewProducer(nsqdTCPAddr string, cfg *nsq.Config) (*nsq.Producer, error) {

	if cfg == nil {
		cfg = nsq.NewConfig()
	}

	p := nsq.NewProducer(nsqdTCPAddr, config)

	err := p.Ping()
	if err != nil {
		return err
	}
}
*/
