/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package standard

func changelogTemplate() string {
	var tmpl = "{{.HEADER}}{{range .COMMITS}}{{.}}{{end}}{{.FOOTER}}"
	return tmpl
}
