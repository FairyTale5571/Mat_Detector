package Mat_Detector

import (
	"regexp"
	"strings"
)

var badPatterns = []string{
	"^(о|а)н(о|а)нист.*",
	"^лошар.*",
	"^к(а|о)злина$",
	"^к(о|а)зел$",
	"^сволоч(ь|ъ|и|уга|ам|ами).*",
	"^лох[уеыаоэяию].*",
	".*урод(ы|у|ам|ина|ины).*",
	".*бля(т|д).*", ".*гандо.*",
	"^м(а|о)нд(а|о).*",
	".*сперма.*",
	".*[уеыаоэяию]еб$",
	"^сучк(а|у|и|е|ой|ай).*",
	"^придур(ок|ки).*",
	"^д(е|и)би(л|лы).*",
	"^сос(ать|и|ешь|у)$",
	"^залуп.*",
	"^муд(е|ил|о|а|я|еб).*",
	".*шалав(а|ы|ам|е|ами).*",
	".*пр(а|о)ст(и|е)т(у|е)тк(а|и|ам|е|ами).*",
	".*шлюх(а|и|ам|е|ами).*",
	".*ху(й|и|я|е|л(и|е)).*",
	".*п(и|е|ы)зд.*",
	"^бл(я|т|д).*",
	"(с|сц)ук(а|о|и|у).*",
	"^еб.*",
	".*(д(о|а)лб(о|а)|разъ|разь|за|вы|по)ебы*.*",
	".*пид(а|о|е)р.*",
	".*хер.*",
}
var goodPatterns = []string{
	".*психу.*",
	".*к(о|а)манд.*",
	".*истр(е|и)блять.*",
	".*л(о|а)х(о|а)трон.*",
	".*(о|а)ск(о|а)рблять.*",
	"хул(е|и)ган",
	".*м(а|о)нд(а|о)рин.*",
	".*р(а|о)ссл(а|о)блять.*",
	".*п(о|а)тр(е|и)блять.*",
	".*@.*\\.(ру|сом|нет)$",
}

var goodWords = []string{
	"дезмонда",
	"застрахуйте",
	"одномандатный",
	"подстрахуй",
	"психуй",
}

var letters = map[string]string{
	"a": "а",
	"b": "в",
	"c": "с",
	"e": "е",
	"f": "ф",
	"g": "д",
	"h": "н",
	"i": "и",
	"k": "к",
	"l": "л",
	"m": "м",
	"n": "н",
	"o": "о",
	"p": "р",
	"r": "р",
	"s": "с",
	"t": "т",
	"u": "у",
	"v": "в",
	"x": "х",
	"y": "у",
	"w": "ш",
	"z": "з",
	"ё": "е",
	"6": "б",
	"9": "д",
}

func ContainsMat(s string) bool {

	s = cleanBadSymbols(s)
	words := strings.Split(s, " ")

	for _, elem := range words {
		var word = convertEngToRus(elem)
		if isInGoodWords(word) && isInGoodPatterns(word) {
			continue
		}
		if isInBadPatterns(word) {
			return true
		}
	}
	if containsMatInSpaceWords(words) {
		return true
	}

	return false
}


func convertEngToRus(s string) string {

	slice := strings.Split(s,"")
	for i := 0; i < len(slice); i++{
		if key, ok := letters[slice[i]]; ok {
			if slice[i] == key {
				slice[i] = key
			}
		}
	}
	return strings.Join(slice,"")
}

func cleanBadSymbols(s string) string {
	s = strings.ToLower(s)
	reg := regexp.MustCompile("/[^a-zA-Zа-яА-Яё0-9\\s]/g")
	res := reg.ReplaceAllString(s,"")
	return res
}

func isInGoodWords(s string) bool {
	return false
}

func isInGoodPatterns(s string) bool {
	return false
}

func isInBadPatterns(s string) bool {
	return false
}

func containsMatInSpaceWords(s []string) bool {
	return false
}

func findSpaceWords() []string {
	var out []string
	return out
}

func addBadPattern() {

}

func addGoodPattern() {

}

func addGoodWord() {

}


func main() {

}
