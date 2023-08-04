/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Count struct {
	Key
	Attributes CountAttributes `json:"attributes"`
}
type CountResponse struct {
	Data     Count    `json:"data"`
	Included Included `json:"included"`
}

type CountListResponse struct {
	Data     []Count  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustCount - returns Count from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCount(key Key) *Count {
	var count Count
	if c.tryFindEntry(key, &count) {
		return &count
	}
	return nil
}
