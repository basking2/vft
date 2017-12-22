package Client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func (c *Client) configureTLS(certPath string) *tls.Config {
	caPool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		c.Log.Fatal(err)
	}
	caPool.AppendCertsFromPEM(serverCert)
	return &tls.Config{
		RootCAs: caPool,
	}
}
