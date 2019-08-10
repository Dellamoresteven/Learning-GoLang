package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
    // name = "One for " + name + ", one for me."
    if name == "" {
        name = "you"
    }
	return "One for " + name + ", one for me."
}
