package main

const link = "http://"

func masker(s string) string {
	var res []byte
	linkLen := len(link)
	linkFounnd := false

	for i := 0; i < len(s); i++ {
		if linkFounnd {
			if s[i] != ' ' {
				res = append(res, '*')
				continue
			}
			res = append(res, s[i])
			linkFounnd = false
			continue
		}
		if len(s[i:]) < linkLen {
			res = append(res, s[i:]...)
			break
		}
		if s[i:i+linkLen] == link {
			linkFounnd = true
			res = append(res, []byte(link)...)
			i += linkLen - 1
		}

		res = append(res, s[i])

	}

	return string(res)
}