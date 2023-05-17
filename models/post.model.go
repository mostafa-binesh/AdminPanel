package models

type Post struct {
	// ID            uint64  `json:"id"`
	// FavoriteColor string  `json:"favoriteColor" admin:"type:text,name:name Karbari,table:true"` // ! table, yani tooye items haye table biad ya na
	// UserID        uint64  `json:"userID" admin:"rel:users"`                                     // if type is int, so relationship is belongsTo
	// Laws          []M.Law `json:"books" admin:"rel:laws"`
	ID     uint   `gorm:"primaryKey"`
	Slug   string `json:"slug"`
	Body   string `json:"body" admin:"type:text,name:Badane,table:false"` //! table, yani tooye items haye table biad
	// Tags   []Tag  `json:"tags" admin:"rel:tags"`                          // relationship type: hasMany
	UserID int64  `json:"user_id" admin:"rel:users"`                      // if type is int, so relationship is belongsTo
}
type Tag struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}
