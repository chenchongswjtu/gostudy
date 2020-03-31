// /opt/gopath/src/github.com/cetc/xledger/orderer/consensus/rbft/rbft_test.go
package main

import (
	"log"
	"testing"
	"time"

	"github.com/cetc/xledger/bccsp/factory"
	"github.com/cetc/xledger/common/channelconfig"
	"github.com/cetc/xledger/common/crypto"
	"github.com/cetc/xledger/common/localmsp"
	"github.com/cetc/xledger/msp"
	"github.com/cetc/xledger/msp/mgmt"
	"github.com/cetc/xledger/orderer/common/blockcutter"
	"github.com/cetc/xledger/orderer/common/localconfig"
	"github.com/cetc/xledger/orderer/common/msgprocessor"
	"github.com/cetc/xledger/orderer/common/multichannel"
	"github.com/cetc/xledger/orderer/consensus"
	cb "github.com/cetc/xledger/protos/common"
	m "github.com/cetc/xledger/protos/msp"
	mspo "github.com/cetc/xledger/protos/msp"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var r = &rbft{
	ordererConfig: localconfig.TopLevel{
		General: localconfig.General{
			LocalMSPDir: "/opt/gopath/src/github.com/cetc/xledger/examples/e2e_cli/crypto-config/ordererOrganizations/example.com/orderers/orderer1.example.com/msp",
			LocalMSPID:  "OrdererMSP",
		},
	},
}

func TestComputeHash(t *testing.T) {
	err := factory.InitFactories(nil)
	if err != nil {
		t.Error("InitFactories err:", err)
	}

	msg := []byte("message")
	var r = &rbft{}
	hash1, err := r.computeHash(msg)
	if err != nil {
		t.Error("computeHash err:", err)
		return
	}

	hash2, err := r.computeHash(msg)
	if err != nil {
		t.Error("computeHash err:", err)
		return
	}

	assert.Equal(t, hash1, hash2)
}

func TestGetCert(t *testing.T) {
	err := factory.InitFactories(nil)
	if err != nil {
		t.Error("InitFactories err:", err)
	}
	id1 := r.getCert()

	cert := `-----BEGIN CERTIFICATE-----
MIICDDCCAbOgAwIBAgIQDRwbRpquYXYqoKM7r6zkETAKBggqgRzPVQGDdTBpMQsw
CQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy
YW5jaXNjbzEUMBIGA1UEChMLZXhhbXBsZS5jb20xFzAVBgNVBAMTDmNhLmV4YW1w
bGUuY29tMB4XDTE5MTIxNzA4NDEzNVoXDTI5MTIxNDA4NDEzNVowWTELMAkGA1UE
BhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lz
Y28xHTAbBgNVBAMTFG9yZGVyZXIxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYI
KoEcz1UBgi0DQgAEBdGXIoeugepxyoVeHXgFFT1YcE4ds84+S/fy3IjJNGPsWus6
uw1xz39UxQ22GekY9TuE1yUw1zIuCxm3QgsZlaNNMEswDgYDVR0PAQH/BAQDAgeA
MAwGA1UdEwEB/wQCMAAwKwYDVR0jBCQwIoAgAyCX2DHaKLPekAHXWRbq9kVugAZc
2zCgFTkh44+Adj4wCgYIKoEcz1UBg3UDRwAwRAIganZ8ZpoL0Bq1OthxmpCuxIms
zHKlFJVPU7FpLeutuPgCIAUDNxit1unKK47x0HgIyqBb9HSx9C+fdJ1gvWx4xVaG
-----END CERTIFICATE-----
`
	sId := &m.SerializedIdentity{
		Mspid:   "OrdererMSP",
		IdBytes: []byte(cert),
	}
	id2, err := proto.Marshal(sId)
	if err != nil {
		t.Error("proto.Marshal err:", err)
	}

	assert.Equal(t, id1, id2)
}

func TestSignBytes(t *testing.T) {
	r.support = support
	msg := []byte("test")
	sig := r.signBytes(msg)

	assert.Equal(t, []byte("111"), sig)

	r.support = nil
	msg = []byte("test")
	sig = r.signBytes(msg)

	// 产生panic
	assert.Panics(t, func() {
		log.Println("panic")
	}, sig)
}

func TestVerifyBytes(t *testing.T) {
	if err := factory.InitFactories(nil); err != nil {
		t.Error("InitFactories err:", err)
	}
	r.support = cs

	msg := []byte("test")
	sig := r.signBytes(msg)

	log.Println(sig)

	if !r.verifyBytes(msg, sig, r.getCert()) {
		t.Error("verify failed")
	}
}

// mock
type mockCS struct {
	//signer crypto.LocalSigner
	si msp.SigningIdentity
}

type mockSI struct {
	signer crypto.Signer
	msp.Identity
}

func (m *mockSI) Serialize() ([]byte, error) {
	panic("implement me")
}

func (m *mockSI) SatisfiesPrincipal(principal *mspo.MSPPrincipal) error {
	panic("implement me")
}

func (m *mockSI) Sign(msg []byte) ([]byte, error) {
	return m.signer.Sign(msg)
}

var _ consensus.ConsenterSupport = new(mockCS)
var _ msp.SigningIdentity = new(mockSI)

func (r *rbft) getID() msp.Identity {
	mspInst := mgmt.GetLocalMSP()
	if mspInst == nil {
		return nil
	}

	id, err := mspInst.DeserializeIdentity(r.getCert())
	if err != nil {
		return nil
	}

	return id
}

var cs = &multichannel.ChainSupport{
	LocalSigner: localmsp.NewSigner(),
}

var si = &mockSI{
	signer: localmsp.NewSigner(),
	//r.getID()
}

var support = &mockCS{
	si: si,
}

func (m *mockCS) NewSignatureHeader() (*cb.SignatureHeader, error) {
	return nil, nil
}

func (m *mockCS) Sign(message []byte) ([]byte, error) {
	return []byte("111"), nil
	//return m.si.Sign(message)
}

func (m *mockCS) ClassifyMsg(chdr *cb.ChannelHeader) msgprocessor.Classification {
	return 0
}

func (m *mockCS) ProcessNormalMsg(env *cb.Envelope) (configSeq uint64, err error) {
	return 0, nil
}

func (m *mockCS) ProcessConfigUpdateMsg(env *cb.Envelope) (config *cb.Envelope, configSeq uint64, err error) {
	return nil, 0, nil
}

func (m *mockCS) ProcessConfigMsg(env *cb.Envelope) (*cb.Envelope, uint64, error) {
	return nil, 0, nil
}

func (m *mockCS) VerifyBlockSignature(arg1 []*cb.SignedData, arg2 *cb.ConfigEnvelope) error {
	return nil
}

func (m *mockCS) BlockCutter() blockcutter.Receiver {
	return nil
}

func (m *mockCS) SharedConfig() channelconfig.Orderer {
	return nil
}

func (m *mockCS) Block(number uint64) *cb.Block {
	return nil
}

func (m *mockCS) WriteBlock(block *cb.Block, encodedMetadataValue []byte) {

}

func (m *mockCS) WriteConfigBlock(block *cb.Block, encodedMetadataValue []byte) {

}

func (m *mockCS) ChainID() string {
	return ""
}

func (m *mockCS) Height() uint64 {
	return 0
}

func (m *mockCS) Sequence() uint64 {
	return 0
}

func (m *mockCS) CreateNextBlock(messages []*cb.Envelope, height uint64) *cb.Block {
	return nil
}

func (m *mockSI) GetPublicVersion() msp.Identity {
	return nil
}

func (m *mockSI) ExpiresAt() time.Time {
	return time.Now()
}

func (m *mockSI) GetIdentifier() *msp.IdentityIdentifier {
	return nil
}

func (m *mockSI) GetMSPIdentifier() string {
	return ""
}

func (m *mockSI) Validate() error {
	return nil
}

func (m *mockSI) GetOrganizationalUnits() []*msp.OUIdentifier {
	return nil
}

func (m *mockSI) Anonymous() bool {
	return false
}

func (m *mockSI) Verify(msg []byte, sig []byte) error {
	return nil
}

func (m *mockCS) Serialize() ([]byte, error) {
	return nil, nil
}

func (m *mockCS) SatisfiesPrincipal(principal *mspo.MSPPrincipal) error {
	return nil
}
