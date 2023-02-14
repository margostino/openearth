package utils

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strconv"
	"strings"
	"unicode"
)

type NameNormalizer struct {
	Value string
}

func NormalizeName(value string) string {
	normalizer := NameNormalizer{Value: value}
	return normalizer.Normalize()
}

// Normalize TODO: implement a more generic normalizer to be able to cover ALL cases.
// TODO: implement shorter naming for long variable names.
func (normalizer *NameNormalizer) Normalize() string {
	return normalizer.toLowercase().
		cleanAccents().
		replace("\t", "").
		replace("=", "").
		replace("(", "").
		replace(")", "").
		replace("[", "").
		replace("]", "").
		replace("\"", "").
		replace("$", "money").
		replace(" - ", "_").
		replace(" -", "_").
		replace("- ", "_").
		replace("-", "_").
		replace(`-`, "_").
		replace(`—`, "_").
		replace(`–`, "_").
		replace(",", "_").
		replace("&", "and").
		replace("%", "perc").
		replace(".", "").
		replace(";", "").
		replace(",", "").
		replace("+", "").
		replace("!", "").
		replace(":", "").
		replace("/", "").
		replace("?", "").
		replace("--", "").
		replace("'s", "s").
		replace("’s", "s").
		replace("'", "").
		replace("[", "_").
		replace("]", "_").
		replace(" ", "_").
		replace("  ", "_").
		replace("    ", "_").
		replace("     ", "_").
		replace("__", "_").
		replace("___", "_").
		replace("_–_", "_").
		replace("_—_", "_").
		replace("__", "_").
		replace("___", "_").
		replace("<", "less").
		replace(">", "greater").
		replace("≥", "greater_or_equal").
		replace("≥", "greater_or_equal").
		replace("₀", "0").
		replace("₁", "1").
		replace("₂", "2").
		replace("₃", "3").
		replace("₄", "4").
		replace("₅", "5").
		replace("₆", "6").
		replace("₇", "7").
		replace("₈", "8").
		replace("₉", "9").
		replace("ö", "o").
		replace("ü", "u").
		replace("è̀", "e").
		replace("̀\t̀", "").
		sanitizeHead().
		sanitizeTail().
		Value
}

func (normalizer *NameNormalizer) cleanAccents() *NameNormalizer {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, normalizer.Value)
	if e != nil {
		panic(e)
	}
	normalizer.Value = output
	return normalizer
}

func (normalizer *NameNormalizer) extractUntilStop(stopFlag string) *NameNormalizer {
	parts := strings.SplitN(normalizer.Value, stopFlag, -1)
	if len(parts) > 0 {
		normalizer.Value = parts[0]
	}
	return normalizer
}

func (normalizer *NameNormalizer) toLowercase() *NameNormalizer {
	normalizer.Value = strings.ToLower(normalizer.Value)
	return normalizer
}

func (normalizer *NameNormalizer) sanitizeHead() *NameNormalizer {
	if len(normalizer.Value) > 0 {
		if _, err := strconv.Atoi(normalizer.Value[0:1]); err == nil {
			normalizer.Value = fmt.Sprintf("o%s", normalizer.Value)
		}
	}
	return normalizer
}

func (normalizer *NameNormalizer) sanitizeTail() *NameNormalizer {
	if strings.HasSuffix(normalizer.Value, "_") {
		normalizer.Value = normalizer.Value[:len(normalizer.Value)-len("_")]
	}
	return normalizer
}

func (normalizer *NameNormalizer) replace(old string, character string) *NameNormalizer {
	normalizer.Value = strings.ReplaceAll(normalizer.Value, old, character)
	return normalizer
}
