package models

type Response struct {
	Products []Product `json:"products"`
}

type Product struct {
    ID                int64       `json:"id"`
    Title             string      `json:"title"`
    BodyHTML          string      `json:"body_html"`
    Vendor            string      `json:"vendor"`
    ProductType       string      `json:"product_type"`
    CreatedAt         string      `json:"created_at"`
    Handle            string      `json:"handle"`
    UpdatedAt         string      `json:"updated_at"`
    PublishedAt       string      `json:"published_at"`
    TemplateSuffix    string      `json:"template_suffix"`
    PublishedScope    string      `json:"published_scope"`
    Tags              string      `json:"tags"`
    Status            string      `json:"status"`
    AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
    Variants          []Variant   `json:"variants"`
    Options           []Option    `json:"options"`
    Images            []Image     `json:"images"`
    Image             string		 `json:"image"`
}

type Variant struct {
    ID                   int64       `json:"id"`
    ProductID            int64       `json:"product_id"`
    Title                string      `json:"title"`
    Price                string      `json:"price"`
    SKU                  string      `json:"sku"`
    Position             int         `json:"position"`
    InventoryPolicy      string      `json:"inventory_policy"`
    CompareAtPrice       string      `json:"compare_at_price"`
    FulfillmentService   string      `json:"fulfillment_service"`
    InventoryManagement  string      `json:"inventory_management"`
    Option1              string      `json:"option1"`
    Option2              string      `json:"option2"`
    Option3              string      `json:"option3"`
    CreatedAt            string      `json:"created_at"`
    UpdatedAt            string      `json:"updated_at"`
    Taxable              bool        `json:"taxable"`
    Barcode              string      `json:"barcode"`
    Grams                int         `json:"grams"`
    ImageID              string		 `json:"image_id"`
    Weight               float64     `json:"weight"`
    WeightUnit           string      `json:"weight_unit"`
    InventoryItemID      int64       `json:"inventory_item_id"`
    InventoryQuantity    int         `json:"inventory_quantity"`
    OldInventoryQuantity int         `json:"old_inventory_quantity"`
    RequiresShipping     bool        `json:"requires_shipping"`
    AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
}

type Option struct {
    ID        int64    `json:"id"`
    ProductID int64    `json:"product_id"`
    Name      string   `json:"name"`
    Position  int      `json:"position"`
    Values    []string `json:"values"`
}

type Image struct {
    Url string `json:"url"`
}

