package auth

import (
	"reflect"
	"testing"
)

func TestIsTokenValid(t *testing.T) {

	// correct Token
	info, ok := IsTokenValid("Bearer", "m7zd25h4pkv9iwss81t4m8hyh24moo")
	correctInfo := TokenValidationResponse{
		ClientID: "gp762nuuoqcoxypju8c569th9wz7q5",
		Login:    "andrewkraevskii",
		Scopes: []string{
			"channel:manage:broadcast",
		},
		UserID:    "278168317",
		ExpiresIn: 0,
	}
	if !ok {
		t.Error("Expected true, got", ok)
	}
	if !reflect.DeepEqual(correctInfo, info) {
		t.Error("Expected", correctInfo, "got", info)
	}

	// wrong Token
	_, ok = IsTokenValid("Bearer", "m7zd25h4pkv9iwss81t4m8hyh24mo1")
	if ok {
		t.Error("Expected false, got", ok)
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
	result, ok := GetToken(client_id, scopes)
	t.Log(result)
	if !ok {
		t.Error("Token getting failed")
	}
	info, ok := IsTokenValid(result.TokenType, result.AccessToken)
	if !ok {
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
	info, _ := IsTokenValid(result.TokenType, result.AccessToken)
	
	ok := RevokeToken(info.ClientID, result.AccessToken)
	if !ok {
		t.Error("Fail deleting token")
	}
	_, ok = IsTokenValid(result.TokenType, result.AccessToken)
	if ok {
		t.Error("Token wasn't deleted")
	}
}
