/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Course struct {
	Key
	Attributes CourseAttributes `json:"attributes"`
}
type CourseResponse struct {
	Data     Course   `json:"data"`
	Included Included `json:"included"`
}

type CourseListResponse struct {
	Data     []Course `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustCourse - returns Course from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCourse(key Key) *Course {
	var course Course
	if c.tryFindEntry(key, &course) {
		return &course
	}
	return nil
}
