package models

type Category struct {
    CategoryID uint   `json:"category_id" gorm:"primaryKey"`
    Name       string `json:"name"`
    // Add this if you want to navigate from Category to Threads (optional)
    Threads    []Thread `gorm:"many2many:thread_categories;"`
}

