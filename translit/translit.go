package translit

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rainycape/unidecode"
)

// TranslitUniversal выполняет транслитерацию и делает web-safe:
// - заменяет все символы на ASCII через unidecode
// - нижний регистр
// - заменяет всё, кроме a-z0-9, на _
func TranslitUniversal(filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	// unidecode для всех символов
	safe := unidecode.Unidecode(name)

	// нижний регистр
	safe = strings.ToLower(safe)

	// заменяем всё, кроме a-z0-9, на _
	re := regexp.MustCompile(`[^a-z0-9]+`)
	safe = re.ReplaceAllString(safe, "_")

	// убираем лишние _ в начале и конце
	safe = strings.Trim(safe, "_")

	return safe + ext
}
