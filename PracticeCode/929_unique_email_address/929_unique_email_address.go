/**
Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain '.'s or '+'s.

If you add periods ('.') between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name.  For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.  (Note that this rule does not apply for domain names.)

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails?

**/

// my solution
// Since there is no set(python datastructure) in go,you need to
// implement it with map by your self

import "strings"

func numUniqueEmails(emails []string) int {
	if len(emails) == 0 {
		return 0
	}

	var localname string
	var domainname string

	dict := make(map[string]bool)
	for _, value := range emails {
		s := strings.Split(value, "@")
		localname, domainname = s[0], s[1]

		// deal with "." situation
		if strings.Contains(localname, ".") {
			// for -1,if n<0,then there will be no limit on the number of replacements
			// check documentation:https://golang.org/pkg/strings/#Replace
			localname = strings.Replace(localname, ".", "", -1)
		}

		if strings.Contains(localname, "+") {
			localname = strings.Split(localname, "+")[0]
		}

		dict[localname+domainname] = true

	}

	return len(dict)

}