package uri2env

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func Parse(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	scheme := u.Scheme
	host, port, _ := net.SplitHostPort(u.Host)
	path := u.Path
	query := u.RawQuery
	queryPs, _ := url.ParseQuery(query)

	for k, m := range queryPs {
		for _, c := range m {

			fmt.Printf("%s=%s\n", k, c)
		}
	}
	fragment := u.Fragment
	user := u.User.Username()
	password, _ := u.User.Password()
	return fmt.Sprintf("SCHEME=%s\nHOST=%s\nPORT=%s\nUSER=%s\nPASSWORD=%s\nPath=%s\nQuery=%s\nFragment=%s", scheme, host, port, user, password, path, query, fragment), nil
}

func AppendIfSet(input *[]string, value string, envVariable string) {
	if value != "" {
		*input = append(*input, fmt.Sprintf("%s=%s", envVariable, value))
	}
}

func ConvertUri(input string, prefix string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	if prefix != "" {
		prefix = prefix + "_"
	}
	var results []string
	AppendIfSet(&results, u.Scheme, prefix+"SCHEME")
	host, port, _ := net.SplitHostPort(u.Host)
	AppendIfSet(&results, host, prefix+"HOST")
	AppendIfSet(&results, port, prefix+"PORT")
	AppendIfSet(&results, u.User.Username(), prefix+"USER")
	password, _ := u.User.Password()
	AppendIfSet(&results, password, prefix+"PASSWORD")
	AppendIfSet(&results, u.Path, prefix+"PATH")
	// AppendIfSet(&results, u.RawQuery, prefix+"QUERY")
	query := u.RawQuery
	queryPs, _ := url.ParseQuery(query)

	for k, m := range queryPs {
		for _, c := range m {

			AppendIfSet(&results, c, strings.ToUpper(prefix+k))
		}
	}
	AppendIfSet(&results, u.Fragment, prefix+"FRAGMENT")
	return strings.Join(results, "\n"), nil
}
