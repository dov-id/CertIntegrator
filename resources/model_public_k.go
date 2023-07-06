/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PublicK struct {
	Key
	Attributes PublicKAttributes `json:"attributes"`
}
type PublicKResponse struct {
	Data     PublicK  `json:"data"`
	Included Included `json:"included"`
}

type PublicKListResponse struct {
	Data     []PublicK `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustPublicK - returns PublicK from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPublicK(key Key) *PublicK {
	var publicK PublicK
	if c.tryFindEntry(key, &publicK) {
		return &publicK
	}
	return nil
}
