package bigbinary_test

import (
	"testing"
	"time"
	"os"
	"fmt"
	"math/rand"
	"strconv"
	"encoding/binary"
	"bytes"
	"github.com/lagarciag/bigbinary"
)

func TestMain(t *testing.M) {
	//Create a fully random seed:
	seed := time.Now().UTC().UnixNano()
	//Take seed fomr env id env var is set
	envSeed := os.Getenv("GO_TEST_RAND_SEED")
	envSeedInt, _ := strconv.ParseUint(envSeed, 10, 64)
	if envSeedInt != 0 {
		seed = int64(envSeedInt)
	}
	//Set Seed
	rand.Seed(seed)
	fmt.Println("Using seed: ", seed)

	v := t.Run()
	os.Exit(v)

}


func TestDefaultBinary(t *testing.T) {
	t.Skip()
	var pi uint8
	b := []byte{0xA, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		t.Log("binary.Read failed:", err)
	}
	t.Log(pi)
}

func TestRead(t *testing.T) {
	b := []byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1}
	//buf := bytes.NewReader(b)

	const offset int = 9
	const length int = 2
	ret, err := bigbinary.Read(b,offset,length)

	if (err != nil) {
		t.Error(err)
	}

	t.Log(ret)

}


