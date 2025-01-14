package main

import (
	"fmt"
	"github.com/jadeydi/jubjub/pkg/blake2s"
	"github.com/jadeydi/jubjub/pkg/pedersenhash"
)

func main() {

	// START pedersen hash test vectors
	hasher, _ := pedersenhash.NewPedersenHasher()

	// length 300 test vector
	b := []bool{false,false,false,false,true,false,true,false,false,false,true,false,false,true,true,false,false,false,true,true,true,true,true,true,true,true,true,true,true,false,false,true,false,false,true,false,false,false,false,false,false,false,false,true,false,true,false,true,true,true,true,false,true,false,false,false,false,true,true,false,true,false,false,false,false,false,true,true,true,true,true,false,false,false,true,true,true,true,true,false,false,true,false,false,true,false,true,true,false,false,false,true,false,true,false,false,false,false,false,false,false,false,false,true,false,true,true,true,true,true,true,false,true,false,true,false,false,false,true,false,false,true,true,false,false,false,false,true,true,false,false,false,false,false,false,true,true,false,false,true,false,false,false,false,false,false,false,true,true,false,false,false,true,true,false,false,false,true,true,false,false,true,true,true,true,true,true,false,true,true,true,false,true,true,true,false,false,false,false,true,false,true,true,false,false,true,true,true,true,false,true,true,true,true,false,false,false,false,false,false,true,false,true,true,false,false,false,true,false,false,true,false,false,false,true,false,true,true,false,true,false,true,true,false,false,true,false,false,true,false,true,true,true,false,true,true,true,false,false,true,true,true,true,false,true,false,false,true,true,false,true,true,false,true,false,true,true,true,true,true,true,true,false,false,false,true,false,false,false,true,false,false,true,true,true,false,false,true,true,true,true,false,true,false,false,true,false,false,true,true,true,false,true,true,true,false,false,true,false,false,}
	p, err := hasher.PedersenHashForBits([]bool{true, false, false, false, false, false}, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", p.Text(10))

	// length 301 test vector
	b = []bool{true,false,false,false,true,false,false,true,true,false,true,false,true,true,false,true,true,true,true,true,false,false,false,true,false,false,true,true,false,true,false,false,false,false,false,true,true,false,true,false,true,false,false,true,true,false,false,false,true,false,false,true,true,true,false,true,true,false,false,true,false,true,false,true,true,false,false,false,true,false,false,false,true,true,true,false,true,false,false,true,true,true,false,true,false,false,true,true,true,false,true,false,true,false,true,false,true,false,true,false,false,false,true,false,true,true,false,true,true,true,true,false,true,false,true,true,false,false,false,true,false,false,false,true,true,true,false,true,false,true,true,false,true,true,true,true,true,true,false,true,true,false,true,true,false,false,false,true,false,true,false,false,false,true,false,true,false,true,true,true,true,false,false,true,false,false,true,false,false,true,true,true,false,true,false,false,true,true,false,true,false,true,true,true,true,true,true,false,false,false,true,true,true,false,false,false,true,true,true,false,false,false,true,false,false,true,true,false,true,true,true,true,true,false,false,false,true,true,true,true,true,true,false,false,false,false,true,true,true,false,true,true,true,true,false,false,false,false,true,true,true,true,true,false,true,true,false,false,true,true,true,true,false,false,true,true,true,true,true,true,false,true,false,true,true,false,false,false,true,false,true,false,false,false,true,true,true,false,false,true,true,true,false,false,false,false,true,true,true,false,false,false,false,true,false,true,true,false,false,false,true,}
	p, err = hasher.PedersenHashForBits([]bool{true, false, false, false, false, false}, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", p.Text(10))
	// END pedersen hash test vectors

	// START blake2 test vectors
	msg := []byte{16,156,4,150,203,88,177,214,232,134,127,32,0,213,89,122,10,95,67,130,197,43,139,137,137,0,25,223,83,32,212,134,160,103,243,142,108,104,98,75,86,150,66,87,224,214,193,30,7,65,23,179,97,114,255,243,235,209,166,236,24,250,121,85,94,88,223,90,157,136,157,203,145,221,68,87,188,73,198,22,123,68,117,204,120,99,52,28,85,86,209,198,8,53,209,92,141,64,48,193,214,133,204,175,86,105,136,236,208,49,206,160,71,9,198,189,192,24,40,2,77,145,200,21,201,51,57,37,53,134,139,109,194,195,29,1,135,116,1,206,97,216,129,177,73,30,216,41,241,3,2,29,51,73,83,242,194,117,127,153,144,134,67,210,65,45,179,73,187,29,34,73,214,208,204,83,135,199,220,110,158,138,209,39,62,197,47,215,220,143,17,10,36,86,107,79,140,96,199,169,140,20,32,22,207,0,217,202,137,1,16,60,103,78,237,100,217,40,187,181,233,98,112,61,151,239,107,209,167,47,233,168,230,35,61,34,90,44,208,32,206,137,233,18,45,148,242,66,141,4,12,240,10,47,131,161,12,50,34,179,161,4,215,62,7,136,78,173,168,74,56,120,35,211,196,14,254,91,6,220,145,25,205,82,159,160,54,158,7,188,56,98,121,172,176,3,178,28,124,147,190,163,63,76,172,186,92,206,151,126,81,158,216,50,111,135,156,1,117,100,43,141,201,109,161,184,84,208,152,187,253,110,210,95,153,27,214,137,222,127,231,143,44,72,229,14,68,156,197,176,150,94,180,222,157,158,164,180,119,20,123,243,242,117,112,180,215,86,226,170,227,228,119,40,219,40,115,149,107,189,217,76,222,106,36,133,36,99,225,90,182,94,216,176,212,138,181,78,146,73,96,78,169,77,190,61,33,253,21,103,206,101,34,129,208,99,181,160,247,244,214,6,1,255,64,11,11,57,22,55,132,27,237,29,160,243,139,200,195,207,108,254,172,90,192,167,29,72,38,246,149,188,157,196,114,103,118,137,101,212,176,169,139,70,97,95,164,4,68,136,166,133,213,14,49,45,106,137,28,156,3,22,204,255,73,217,12,76,95,134,86,167,18,159,82,115,6,60,127,74,72,78,64,23,61,207,164,234,60,81,44,30,115,20,156,94,45,176,102,49,133,201,135,57,180,17,97,242,242,125,175,116,225,207,194,133,104,147,207,165,59,199,119,73,103,46,44,3,93,124,40,37,80,63,109,4,145,179,35,127,107,156,251,14,29,83,254,150,106,223,76,100,195,168,251,104,98,254,101,179,67,121,247,93,10,46,220,65,245,149,110,101,175,239,128,5,47,75,149,172,192,111,131,28,126,181,27,18,64,58,38,219,7,117,79,115,206,186,184,243,188,199,14,161,248,162,172,234,235,2,195,33,105,13,155,26,235,191,211,38,54,61,12,210,181,136,22,202,80,171,147,212,255,27,39,66,226,108,21,159,25,152,103,130,222,140,74,216,165,52,177,235,22,0,60,162,75,188,72,232,38,117,220,191,7,181,111,100,88,210,40,252,225,75,11,41,109,98,178,3,95,242,113,97,184,79,22,192,131,244,227,21,53,254,183,231,155,63,232,63,172,201,8,13,161,223,164,160,225,125,147,230,191,29,98,109,152,196,86,9,52,29,103,170,96,204,237,0,82,21,230,193,89,118,240,16,84,249,115,79,139,128,242,172,208,226,231,123,159,77,178,107,246,189,162,127,60,190,235,150,133,77,240,46,236,71,204,113,142,88,135,23,214,191,172,77,18,13,150,81,71,190,238,52,200,193,78,224,225,104,140,49,238,141,113,224,77,159,166,124,194,40,22,147,229,25,188,84,238,197,246,25,206,124,60,91,54,42,137,208,117,149,169,205,131,214,203,59,195,191,20,187,88,159,167,90,88,224,8,20,152,135,134,254,245,25,76,66,175,32,126,73,100,83,121,59,201,243,240,140,161,210,176,78,6,162,62,83,234,225,26,11,51,191,111,80,63,62,90,234,114,8,206,89,48,235,208,5,167,162,93,79,24,184,78,235,16,234,109,134,37,232,250,209,188,210,229,106,26,108,255,140,58,98,151,158,62,69,68,0,79,16,253,218,103,145,238,69,79,206,149,207,9,10,149,112,230,77,92,100,17,219,174,129,195,182,68,217,250,50,182,175,12,85,50,131,221,90,163,0,11,250,110,134,246,30,153,199,220,164,251,180,14,231,122,123,50,190,69,207,212,159,215,100,237,146,70,107,87,97,211,163,2,225,97,245,218,177,172,210,65,185,253,146,}
	blake, err := blake2s.New256WithPersonalization(nil, []byte("12345678"))
	if err != nil {
		panic(err)
	}

	_, err = blake.Write(msg)
	if err != nil {
		panic(err)
	}

	blakeHashBytes := blake.Sum(nil)
	fmt.Printf("hash: %x\n", blakeHashBytes)

	msg = []byte{64,69,62,222,158,105,237,77,245,84,184,99,13,114,225,207,109,67,40,123,123,150,208,86,248,53,177,135,162,36,100,122,173,82,21,8,223,239,68,73,75,6,169,141,125,28,63,51,121,168,199,83,216,9,17,155,188,190,220,221,70,68,190,232,79,63,20,128,37,68,10,83,53,151,60,209,142,156,26,82,231,116,136,183,202,28,57,54,113,20,102,86,215,103,87,60,35,77,41,225,203,243,122,27,12,185,133,48,25,139,112,205,137,94,184,147,201,118,60,106,186,219,224,108,49,248,74,71,99,252,137,70,175,40,139,119,78,223,83,51,166,150,138,158,126,179,214,0,100,70,47,56,104,194,233,82,95,151,7,157,81,73,5,190,160,244,126,10,187,123,12,116,66,37,176,177,167,91,233,220,184,149,136,178,41,216,66,188,85,147,77,119,93,2,118,17,231,152,61,45,185,172,90,123,192,183,111,6,162,208,89,202,173,36,243,60,5,49,197,22,179,38,220,222,44,130,16,235,210,168,196,44,202,119,38,147,205,203,111,214,158,88,110,8,106,163,40,31,129,9,223,190,210,252,183,156,221,186,194,69,148,253,169,222,164,3,25,214,13,98,19,254,215,7,4,73,120,184,143,7,182,173,92,132,57,148,25,225,241,123,63,181,79,100,134,186,162,193,249,131,125,5,249,249,19,55,171,197,125,245,40,24,116,213,191,126,203,150,98,135,40,106,236,2,26,70,253,187,198,209,167,36,62,67,189,105,66,109,35,45,80,24,29,179,58,231,91,42,73,166,198,43,93,239,20,42,168,208,250,240,167,63,129,152,215,127,177,16,175,104,178,65,8,53,235,189,225,72,22,248,152,43,183,201,110,47,222,131,53,59,56,7,135,104,203,62,131,0,60,116,172,44,216,209,5,7,76,213,38,193,248,132,13,207,35,191,127,124,182,73,16,119,144,128,67,194,0,59,203,132,121,68,13,231,78,79,3,35,67,121,177,8,198,179,161,251,208,160,86,239,172,223,81,223,55,92,119,120,71,170,229,80,177,182,39,199,201,160,110,193,154,168,216,74,252,112,229,16,41,239,240,86,191,139,233,150,173,26,61,170,127,24,214,185,107,166,37,244,66,181,129,247,34,211,126,58,224,225,29,184,45,196,8,232,157,38,151,112,229,255,106,77,195,76,175,29,155,13,41,187,5,79,22,121,72,55,1,39,204,96,143,112,123,200,4,1,33,212,135,247,64,14,79,190,210,67,126,199,35,64,154,252,126,84,121,172,16,34,106,160,219,79,217,16,43,77,65,223,25,144,50,138,207,47,201,131,131,244,67,21,47,48,212,214,84,146,143,32,114,47,233,241,81,62,16,145,136,168,186,6,94,34,25,122,70,119,42,122,241,211,112,2,125,90,108,232,245,105,54,103,166,94,102,208,30,76,146,226,199,71,32,59,173,39,33,175,31,150,46,206,85,177,182,145,249,189,124,158,250,144,43,57,162,214,173,58,56,194,129,25,57,175,50,248,83,91,182,149,48,225,131,52,72,56,86,69,169,69,167,174,236,53,120,128,2,49,247,43,17,210,15,93,147,130,213,165,46,208,234,197,91,194,2,155,86,195,178,156,91,214,66,164,67,76,150,170,94,229,234,120,171,2,22,41,86,103,1,49,201,144,246,85,142,144,30,218,128,17,41,206,92,153,200,182,231,175,151,179,44,160,94,144,96,106,39,32,162,216,251,225,73,75,192,77,142,95,47,13,150,184,95,130,194,108,133,21,70,55,20,135,52,173,0,148,129,205,134,77,25,212,150,35,32,85,22,245,241,241,232,161,204,185,58,182,238,48,14,129,39,234,144,1,157,87,245,240,8,58,130,209,106,34,108,55,37,174,109,82,78,200,57,226,90,167,66,246,45,11,143,88,71,138,106,161,34,188,15,238,219,116,31,141,28,149,233,128,144,8,105,65,135,86,194,70,232,93,197,81,212,206,92,245,8,220,82,205,95,24,94,174,144,142,58,138,115,242,128,107,24,85,81,76,254,127,68,144,105,251,20,148,140,236,113,190,51,69,69,71,169,229,252,21,83,61,145,62,123,157,170,198,240,174,201,95,86,114,101,43,13,23,238,234,99,175,230,119,106,198,249,130,47,46,132,74,212,213,201,1,201,76,145,6,146,63,97,58,12,187,33,18,74,186,5,5,26,127,250,184,114,146,224,148,12,19,206,151,144,210,119,7,187,24,168,236,85,174,108,175,43,13,86,161,76,108,101,15,3,120,1,246,62,243,183,5,19,68,126,70,250,122,198,175,77,238,}
	blake, err = blake2s.New256WithPersonalization(nil, []byte("12345678"))
	if err != nil {
		panic(err)
	}

	_, err = blake.Write(msg)
	if err != nil {
		panic(err)
	}

	blakeHashBytes = blake.Sum(nil)
	fmt.Printf("hash: %x\n", blakeHashBytes)
}
