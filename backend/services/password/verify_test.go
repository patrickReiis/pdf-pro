package password

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	type testVerifyCase struct {
		plainPwd  string
		hashedPwd string
		isEqual   bool
	}

	testCases := []testVerifyCase{
		{"plain", "$argon2id$v=20$m=16,t=2,p=1$TlQ4VlFUSTVVOXdYdWRSeg$DOPiC1dzuCWQ8fp20U7hKQ", false},
		{"plain", "$argon2id$v=19$m=16,t=4,p=1$cHhMQzNramZKMmhud1VSVA$mkFEIXUkAagVsHNZ2jgJvjwEHOuLsKk6", true},
		{"very long password bla bla bla", "$argon2id$v=19$m=40000,t=14,p=3$a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", true},
		{"very long password bla bla bla", "argon2d$v=19$m=40000,t=14,p=3$a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", false},
		{"very long password bla bla bla", "argon2d$v=19$m=40000,t=not number,=3$a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", false},
		{"very long password bla bla bla", "$argon2id$v=19$m=40000,t=14,p=3$ç~´´´a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", false},
		{"very long password bla bla bla", "argon2d$v=19$m=40000,t=14,p=3$a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFç]]ç~KGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", false},
		{"very long password bla bla bla", "$argon2id$v=19$m=40000,t=14,p=3$a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Yb", false},
		{"very long password bla bla bla", "$argon2i$v=19$m=40000,t=14,p=3$a0ZvNFdnZnpOZDEwTGNtZA$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", false},
		{"plain", "$argon2id$v=19$f=false$cHhMQzNramZKMmhud1VSVA$mkFEIXUkAagVsHNZ2jgJvjwEHOuLsKk6", false},
		{"very long password bla bla bla", "$argon2id$v=19$m=40000,t=14,p=3$a0ZvNFdnZnpOZDEwTG$nWMoW4vrZFKGvAI5z6TB4GXOqIDptyos1yTW1Ybykm4", false},
		{"plain", "$argon2id$v=19$m=16,t=2,p=1$R0xKclFSQ3VmRUhUTlNNag$cRR+6shII+S2fRR1Tp8XqA", false},
	}

	for _, e := range testCases {
		t.Run(fmt.Sprintf("%s -> %s is %t", e.plainPwd, e.hashedPwd, e.isEqual), func(t *testing.T) {
			isEqual := Verify(e.hashedPwd, e.plainPwd)
			if isEqual != e.isEqual {
				t.Errorf("Expected %t, got %t instead", e.isEqual, isEqual)
			}
		})
	}

}
