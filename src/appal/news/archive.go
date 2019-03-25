package news

type Topic struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	Description string `json:"description"`
	UrlToImage  string `json:"urltoimage"`
	PublishedAt string `json:"publishedat"`
}

type Archive map[string][]Topic

func New() Archive {
	return make(map[string][]Topic, 0)
}

// method collect для Archive
func (a Archive) collect(category string) {
	sources := getSources(category)
	topics := getTopics(sources)
	a[category] = topics
}

// method result для Archive
func (a Archive) result(category string) []Topic {
	return a[category]
}
