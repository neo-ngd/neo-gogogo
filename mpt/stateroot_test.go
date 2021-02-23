package mpt

import (
	"encoding/json"
	"testing"

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

func TestStateRootSerialize(t *testing.T) {
	var sr StateRoot
	data := []byte(`{"version":0,"index":409,"prehash":"0x5094cbeb0642557262cb3ea61d57a1b47119eafce9725ec4ac79398c30b3c7a0","stateroot":"0x0000000000000000000000000000000000000000000000000000000000000000","witness":{"invocation":"40cd8d09b0f8bec0c342ebfc1024d0805297b439fdff3a2eccad01c1bd1ab2e8d4671bb1780bb5c8199092e75dfde567db19117f87494012491751a2cfaca42179409c54a8a1d7ae58b8090e2526c548294d37c909c671e64b2a9c5acf83a4d0605242e6f6ffcc3fe240b94c5446cc588256daa1f0bf065d27b881a43404c9d0be8840c9ba349daf13dae1ff77adf5ca7d3a0a913ba534ce3fc5e645a950fb628b61e96c470e3efba66fb4f3e7fc5aef7396949eda98bd080cff14dec42f7cb791226a","verification":"5321030f59a5482a4e42a2e5a848608dac4e84a698e567e2860e0ca5f23fc9e818d37c21032e78261370d4d62cf4c13584ca90f46c5565117b5b97544312f2e7b7c36b9eba21026e271722c21c482f0ac74dd932e61cdc2a2dd889633a2c5d8ecef43f2769f51e2103d55bfbcd493d06ab49c09cde0cea5d9ba890d81331a2fcd6f68d329932d0398f54ae"}}`)
	err := json.Unmarshal(data, &sr)
	if err != nil {
		t.Error(err)
	}
	data, _ = io.ToArray(&sr)
	if helper.BytesToHex(data) != "0099010000a0c7b3308c3979acc45e72e9fcea1971b4a1571da63ecb6272554206ebcb9450000000000000000000000000000000000000000000000000000000000000000001c340cd8d09b0f8bec0c342ebfc1024d0805297b439fdff3a2eccad01c1bd1ab2e8d4671bb1780bb5c8199092e75dfde567db19117f87494012491751a2cfaca42179409c54a8a1d7ae58b8090e2526c548294d37c909c671e64b2a9c5acf83a4d0605242e6f6ffcc3fe240b94c5446cc588256daa1f0bf065d27b881a43404c9d0be8840c9ba349daf13dae1ff77adf5ca7d3a0a913ba534ce3fc5e645a950fb628b61e96c470e3efba66fb4f3e7fc5aef7396949eda98bd080cff14dec42f7cb791226a8b5321030f59a5482a4e42a2e5a848608dac4e84a698e567e2860e0ca5f23fc9e818d37c21032e78261370d4d62cf4c13584ca90f46c5565117b5b97544312f2e7b7c36b9eba21026e271722c21c482f0ac74dd932e61cdc2a2dd889633a2c5d8ecef43f2769f51e2103d55bfbcd493d06ab49c09cde0cea5d9ba890d81331a2fcd6f68d329932d0398f54ae" {
		t.Error("serialize failed")
	}
	io.AsSerializable(&sr, data)
	data, _ = io.ToArray(&sr)
	if helper.BytesToHex(data) != "0099010000a0c7b3308c3979acc45e72e9fcea1971b4a1571da63ecb6272554206ebcb9450000000000000000000000000000000000000000000000000000000000000000001c340cd8d09b0f8bec0c342ebfc1024d0805297b439fdff3a2eccad01c1bd1ab2e8d4671bb1780bb5c8199092e75dfde567db19117f87494012491751a2cfaca42179409c54a8a1d7ae58b8090e2526c548294d37c909c671e64b2a9c5acf83a4d0605242e6f6ffcc3fe240b94c5446cc588256daa1f0bf065d27b881a43404c9d0be8840c9ba349daf13dae1ff77adf5ca7d3a0a913ba534ce3fc5e645a950fb628b61e96c470e3efba66fb4f3e7fc5aef7396949eda98bd080cff14dec42f7cb791226a8b5321030f59a5482a4e42a2e5a848608dac4e84a698e567e2860e0ca5f23fc9e818d37c21032e78261370d4d62cf4c13584ca90f46c5565117b5b97544312f2e7b7c36b9eba21026e271722c21c482f0ac74dd932e61cdc2a2dd889633a2c5d8ecef43f2769f51e2103d55bfbcd493d06ab49c09cde0cea5d9ba890d81331a2fcd6f68d329932d0398f54ae" {
		t.Error("serialize failed")
	}
}
