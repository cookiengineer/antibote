package gpg

import "os"
import "os/exec"
import "strings"

func ToKeyID(value string) string {

	var result string

	err1 := os.WriteFile("/tmp/key.sig", []byte(value), 0666)

	if err1 == nil {

		cmd := exec.Command("gpg", "--list-packets", "/tmp/key.sig")
		buffer, err2 := cmd.Output()

		if err2 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

				line := lines[l]

				if strings.HasPrefix(line, ":signature packet:") && strings.Contains(line, " keyid ") {
					result = strings.TrimSpace(line[strings.Index(line, " keyid ")+7:])
				}

			}

		}

	}

	return result

}
