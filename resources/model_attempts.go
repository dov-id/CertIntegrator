/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Attempts struct {
	Key
	Attributes AttemptsAttributes `json:"attributes"`
}
type AttemptsResponse struct {
	Data     Attempts `json:"data"`
	Included Included `json:"included"`
}

type AttemptsListResponse struct {
	Data     []Attempts `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustAttempts - returns Attempts from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAttempts(key Key) *Attempts {
	var attempts Attempts
	if c.tryFindEntry(key, &attempts) {
		return &attempts
	}
	return nil
}
