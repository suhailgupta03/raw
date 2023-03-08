package core

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"hash/fnv"
)

type DUtil struct {
}

func DBUtil() *DUtil {
	return &DUtil{}
}

func (util *DUtil) GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func (util *DUtil) GenerateHashNumber(primaryKeyValue string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(primaryKeyValue))
	return h.Sum32()
}

func (util *DUtil) GeneratePrimaryKeyHash(primaryKeyValue string) string {
	hash := md5.Sum([]byte(primaryKeyValue))
	return hex.EncodeToString(hash[:])
}

// GetPrimaryKeyRoot This functions returns the path
func (util *DUtil) GetPrimaryKeyRoot() string {
	return "PK"
}
