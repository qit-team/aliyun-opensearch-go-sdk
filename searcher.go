package aliyun_opensearch_go_sdk

type OpenSearcher interface {
	Name() string
	// todo: to supply params
	Search()
	Update()
	Delete()
}

type AliOpenSearcher struct {
	name   string
	client OpenSearchClient
}

func NewAliOpenSearcher(name string, client OpenSearchClient) OpenSearcher {
	if name == "" || len(name) == 0 {
		panic(PanicInfoPrefix + "table name could not be empty")
	}

	queue := new(AliOpenSearcher)
	queue.client = client
	queue.name = name

	return queue
}

func (p *AliOpenSearcher) Name() string {
	return p.name
}

func (p *AliOpenSearcher) Search() {
	tempHeader := map[string]string{
		"test": "1",
	}
	p.client.Send(GET, tempHeader, "test", "test")
}

func (p *AliOpenSearcher) Update() {
	tempHeader := map[string]string{
		"test": "1",
	}
	p.client.Send(GET, tempHeader, "test", "test")
}

func (p *AliOpenSearcher) Delete() {
	tempHeader := map[string]string{
		"test": "1",
	}
	p.client.Send(GET, tempHeader, "test", "test")
}
