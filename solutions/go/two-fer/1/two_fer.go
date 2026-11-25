// twofer package: reply with correct wording
package twofer

// ShareWith takes in the name of the person sharing a cookie with
// and returns the whole formatted string.
func ShareWith(name string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
    msg := ""
    if name != "" {
        msg += "One for " + name + ", one for me."
    } else {
        msg += "One for you, one for me."
    }
	return msg
}
