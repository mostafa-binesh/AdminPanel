package utils
// ! Relationship
type RelationType int8

const (
	NoRelationship RelationType = iota // 0
	BelongsTo                          // 1
	HasMany                            // 2
)