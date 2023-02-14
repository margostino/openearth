package utils

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, expected string, result string) {
	if result != expected {
		t.Fatalf(fmt.Sprintf("expected %s but it was %s", expected, result))
	}
}

func TestNameNormalizer(t *testing.T) {
	var original, expected, result string

	original = "Top Net Personal Wealth Shares – WID (2018)"
	expected = "top_net_personal_wealth_shares_wid_2018"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "Trade – Giovanni and Tena-Junguito (2016)"
	expected = "trade_giovanni_and_tena_junguito_2016"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "È,É,Ê,Ë,Û,Ù,Ï,Î,À,Â,Ô,è,é,ê,ë,û,ù,ï,î,à,â,ô,Ç,ç,Ã,ã,Õ,õ"
	expected = "e_e_e_e_u_u_i_i_a_a_o_e_e_e_e_u_u_i_i_a_a_o_c_c_a_a_o_o"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "çÇáéíóúýÁÉÍÓÚÝàèìòùÀÈÌÒÙãõñäëïöüÿÄËÏÖÜÃÕÑâêîôûÂÊÎÔÛ"
	expected = "ccaeiouyaeiouyaeiouaeiouaonaeiouyaeiouaonaeiouaeiou"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "Gisele Bündchen da Conceição e Silva foi batizada assim em homenagem à sua conterrânea de Horizontina, RS."
	expected = "gisele_bundchen_da_conceicao_e_silva_foi_batizada_assim_em_homenagem_a_sua_conterranea_de_horizontina_rs"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "Share of drinkers who \"binged\" on heaviest day of drinking in last week"
	expected = "share_of_drinkers_who_binged_on_heaviest_day_of_drinking_in_last_week"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "Percentage of children (2–14) who experience any violent discipline (UNICEF Global Databases (2016))"
	expected = "percentage_of_children_2_14_who_experience_any_violent_discipline_unicef_global_databases_2016"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "PRIO–UCDP"
	expected = "prio_ucdp"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "PM₂.₅"
	expected = "pm25"
	result = NormalizeName(original)
	assert(t, expected, result)

	original = "PM₁₀ (Index)"
	expected = "pm10_index"
	result = NormalizeName(original)
	assert(t, expected, result)

	//TODO
}

func TestToSnakeCase(t *testing.T) {

	original := "TopIncomeSharesWorldWealthAndIncomeDatabase2018"
	expected := "top_income_shares_world_wealth_and_income_database_2018"
	result := ToSnakeCase(original)
	assert(t, expected, result)

}
