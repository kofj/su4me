package main

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

func uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func bytesToUint64(bs []byte) uint64 {
	return binary.BigEndian.Uint64(bs)
}

func createBuckets(name []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(name)
		return err
	})
}
