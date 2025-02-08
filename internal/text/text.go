package text

import "strings"

func Lines(text string, maxLength int) []string {
	result := make([]string, 0)

	textLines := strings.Split(strings.ReplaceAll(text, "\t\n", "\n"), "\n")
	for _, textLine := range textLines {
		if len(textLine) <= maxLength {
			textLine = textLine + "\n"
			result = append(result, textLine)
			continue
		}
		sb := strings.Builder{}
		for _, word := range strings.Split(textLine, " ") {
			lengthWithNew := sb.Len() + len(word)
			if lengthWithNew > maxLength {
				// TODO: review edge cases (string exactly max, one below and one less)
				if lengthWithNew != maxLength {
					sb.WriteRune(' ')
				}
				result = append(result, sb.String())
				sb.Reset()
			}
			if sb.Len() > 0 {
				sb.WriteRune(' ')
			}
			sb.WriteString(word)
		}
		if sb.Len() > 0 {
			sb.WriteRune('\n')
			result = append(result, sb.String())
			sb.Reset()
		}
	}

	return result
}
