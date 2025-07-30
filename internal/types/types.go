package types

// IPResponse - структура для ответа IP API
type IPResponse struct {
	IP        string `json:"ip"`
	UserAgent string `json:"userAgent"`
}

// PageData - структура данных страницы
type PageData struct {
	Title     string
	Header    string
	Subheader string
	Content   string
	IPResponse // Встроенная структура
}

