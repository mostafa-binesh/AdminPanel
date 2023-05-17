package models

type Post struct {
	// ID            uint64  `json:"id"`
	// FavoriteColor string  `json:"favoriteColor" admin:"type:text,name:name Karbari,table:true"` // ! table, yani tooye items haye table biad ya na
	// UserID        uint64  `json:"userID" admin:"rel:users"`                                     // if type is int, so relationship is belongsTo
	// Laws          []M.Law `json:"books" admin:"rel:laws"`
	ID     uint   `gorm:"primaryKey"`
	Slug   string `json:"slug"`
	Body   string `json:"body" admin:"type:text,name:Badane,table:false"` //! table, yani tooye items haye table biad
	Tags   []Tag  `json:"tags" admin:"rel:Tags"`                          // relationship type: hasMany
	// UserID int64  `json:"user_id" admin:"rel:User"`                      // if type is int, so relationship is belongsTo
	UserID int64  `json:"user_id"`                      // if type is int, so relationship is belongsTo
	User User `json:"user" admin:"rel:User"`
}
type Tag struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Post Post `json:"post" admin:"rel:Post"`
	PostID uint `json:"post_id"`
}
