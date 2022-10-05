package node

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var node string

func Get() string {
	return node
}

func init() {
	node = nodeIdentifier()

	fmt.Println("Node:", node)
}

func nodeIdentifier() string {
	rand.Seed(time.Now().UnixNano())

	randInt := rand.Intn(10000)
	randIntToString := strconv.Itoa(randInt)

	currTime := time.Now().UTC()
	str := []byte(currTime.String() + randIntToString)

	sum := md5.Sum(str)
	hash := hex.EncodeToString(sum[:])

	return hash
}
