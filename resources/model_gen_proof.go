/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GenProof struct {
	Key
	Attributes GenProofAttributes `json:"attributes"`
}
type GenProofRequest struct {
	Data     GenProof `json:"data"`
	Included Included `json:"included"`
}

type GenProofListRequest struct {
	Data     []GenProof `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustGenProof - returns GenProof from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGenProof(key Key) *GenProof {
	var genProof GenProof
	if c.tryFindEntry(key, &genProof) {
		return &genProof
	}
	return nil
}
