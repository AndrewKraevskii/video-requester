package auth

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {

	// correct Token
	info, err := Validate("Bearer", "m7zd25h4pkv9iwss81t4m8hyh24moo")
	correctInfo := TokenValidationResponse{
		ClientID: "gp762nuuoqcoxypju8c569th9wz7q5",
		Login:    "andrewkraevskii",
		Scopes: []string{
			"channel:manage:broadcast",
		},
		UserID:    "278168317",
		ExpiresIn: 0,
	}
	if err != nil {
		t.Error("Expected true, got", err)
	}
	if !reflect.DeepEqual(correctInfo, info) {
		t.Error("Expected", correctInfo, "got", info)
	}

	// wrong Token
	_, err = Validate("Bearer", "m7zd25h4pkv9iwss81t4m8hyh24mo1")
	if err != nil {
		t.Error("Expected false, got", err)
	}

}

func TestGetToken(t *testing.T) {
	const client_id = "to0d2ggvuyjpadj2cdxxoxf0kw2flj"
	scopes := []string{
		"channel:manage:redemptions",
		"channel:read:redemptions",
		"chat:edit",
		"chat:read",
	}
	result, err := GetToken(client_id, scopes)
	t.Log(result)
	if err != nil {
		t.Error("Token getting failed")
	}
	info, err := Validate(result.TokenType, result.AccessToken)
	if err != nil {
		t.Error("Get invalid token")
		// f.Error(info)
	}
	if !reflect.DeepEqual(info.Scopes, scopes) {
		t.Error("scopes don't match", info.Scopes, scopes)
	}
}

func TestRevokeToken(t *testing.T) {
	const client_id = "to0d2ggvuyjpadj2cdxxoxf0kw2flj"

	result, _ := GetToken(client_id, []string{})
	info, _ := Validate(result.TokenType, result.AccessToken)

	err := RevokeToken(info.ClientID, result.AccessToken)
	if err != nil {
		t.Error("Fail deleting token")
	}
	_, err = Validate(result.TokenType, result.AccessToken)
	if err != nil {
		t.Error("Token wasn't deleted")
	}
}
