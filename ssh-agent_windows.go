package sshagent

import (
	"golang.org/x/crypto/ssh"
	"github.com/davidmz/go-pageant"
)

func New() func() ([]ssh.Signer, error) {
	aconn := pageant.New()

	return aconn.Signers
}
