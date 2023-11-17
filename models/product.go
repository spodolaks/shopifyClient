package models

type Response struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID                int64           `json:"id"`
	Title             string          `json:"title"`
	BodyHTML          string          `json:"body_html"`
	Vendor            string          `json:"vendor"`
	ProductType       string          `json:"product_type"`
	CreatedAt         string          `json:"created_at"`
	Handle            string          `json:"handle"`
	UpdatedAt         string          `json:"updated_at"`
	PublishedAt       string          `json:"published_at"`
	TemplateSuffix    interface{}     `json:"template_suffix"`
	PublishedScope    string          `json:"published_scope"`
	Tags              string          `json:"tags"`
	Status            string          `json:"status"`
	AdminGraphqlAPIID string          `json:"admin_graphql_api_id"`
	Variants          []Variant       `json:"variants"`
	Options           []Option        `json:"options"`
	Images            []ProductImage  `json:"images"`
	Image             *ProductImage   `json:"image"`
}

type Variant struct {
	ID                   int64            `json:"id"`
	ProductID            int64            `json:"product_id"`
	Title                string           `json:"title"`
	Price                string           `json:"price"`
	Sku                  string           `json:"sku"`
	Position             int              `json:"position"`
	InventoryPolicy      string           `json:"inventory_policy"`
	CompareAtPrice       interface{}      `json:"compare_at_price"`
	FulfillmentService   string           `json:"fulfillment_service"`
	InventoryManagement  string           `json:"inventory_management"`
	Option1              string           `json:"option1"`
	Option2              interface{}      `json:"option2"`
	Option3              interface{}      `json:"option3"`
	CreatedAt            string           `json:"created_at"`
	UpdatedAt            string           `json:"updated_at"`
	Taxable              bool             `json:"taxable"`
	Barcode              string           `json:"barcode"`
	Grams                int              `json:"grams"`
	ImageID              interface{}      `json:"image_id"`
	Weight               float64          `json:"weight"`
	WeightUnit           string           `json:"weight_unit"`
	InventoryItemID      int64            `json:"inventory_item_id"`
	InventoryQuantity    int              `json:"inventory_quantity"`
	OldInventoryQuantity int              `json:"old_inventory_quantity"`
	PresentmentPrices    []PresentmentPrice `json:"presentment_prices"`
	RequiresShipping     bool             `json:"requires_shipping"`
	AdminGraphqlAPIID    string           `json:"admin_graphql_api_id"`
}

type PresentmentPrice struct {
	Price            PriceStruct `json:"price"`
	CompareAtPrice   interface{} `json:"compare_at_price"`
}

type PriceStruct struct {
	Amount         string `json:"amount"`
	CurrencyCode   string `json:"currency_code"`
}

type Option struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}

type ProductImage struct {
	ID                int64      `json:"id"`
	ProductID         int64      `json:"product_id"`
	Position          int        `json:"position"`
	CreatedAt         string     `json:"created_at"`
	UpdatedAt         string     `json:"updated_at"`
	Alt               interface{} `json:"alt"`
	Width             int        `json:"width"`
	Height            int        `json:"height"`
	Src               string     `json:"src"`
	VariantIDs        []int64    `json:"variant_ids"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id"`
}