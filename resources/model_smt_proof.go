/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SmtProof struct {
	Key
	Attributes SmtProofAttributes `json:"attributes"`
}
type SmtProofResponse struct {
	Data     SmtProof `json:"data"`
	Included Included `json:"included"`
}

type SmtProofListResponse struct {
	Data     []SmtProof `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustSmtProof - returns SmtProof from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSmtProof(key Key) *SmtProof {
	var sMTProof SmtProof
	if c.tryFindEntry(key, &sMTProof) {
		return &sMTProof
	}
	return nil
}
