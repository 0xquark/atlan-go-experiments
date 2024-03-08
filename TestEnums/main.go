package main

import (
	"fmt"
	"github.com/orsinium-labs/enum"
)

type AtlanTypeCategory enum.Member[string]

// Constants representing type categories
var (
	AtlanTypeCategoryEnum           = AtlanTypeCategory{"ENUM"}
	AtlanTypeCategoryStruct         = AtlanTypeCategory{"STRUCT"}
	AtlanTypeCategoryClassification = AtlanTypeCategory{"CLASSIFICATION"}
	AtlanTypeCategoryEntity         = AtlanTypeCategory{"ENTITY"}
	AtlanTypeCategoryRelationship   = AtlanTypeCategory{"RELATIONSHIP"}
	AtlanTypeCategoryCustomMetadata = AtlanTypeCategory{"BUSINESS_METADATA"}
)

type CustomMetadataDef struct {
	Category *AtlanTypeCategory `json:"category"`
}

func (c *CustomMetadataDef) GetCategory(category AtlanTypeCategory) {
	fmt.Println("Category:", c.Category.Value)
}

func main() {
	atlanTypeCategory := AtlanTypeCategoryCustomMetadata
	ctx := &CustomMetadataDef{
		Category: &atlanTypeCategory,
	}
	ctx.GetCategory(atlanTypeCategory)
}
