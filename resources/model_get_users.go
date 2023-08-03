/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetUsers struct {
	Key
	Attributes GetUsersAttributes `json:"attributes"`
}
type GetUsersRequest struct {
	Data     GetUsers `json:"data"`
	Included Included `json:"included"`
}

type GetUsersListRequest struct {
	Data     []GetUsers `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustGetUsers - returns GetUsers from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGetUsers(key Key) *GetUsers {
	var getUsers GetUsers
	if c.tryFindEntry(key, &getUsers) {
		return &getUsers
	}
	return nil
}
