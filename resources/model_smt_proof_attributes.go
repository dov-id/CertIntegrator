/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SmtProofAttributes struct {
	// key hex string for node for which proof was generated
	NodeKey string `json:"node_key"`
	// value hex string for node for which proof was generated
	NodeValue string `json:"node_value"`
	// proof itself generated for specified key
	Proof []string `json:"proof"`
}
