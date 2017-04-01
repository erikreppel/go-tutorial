package main

import (
	"crypto/sha256"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Lets say we want need to manipulate data then store it.
// We could choose a backend and just go for it, or we could write an interface
// that contains our desired behavior, then implement that interface.

// In this example we want to hash data and then persist that data.
// We also later want to retrieve that data but without any in-memory caching

// Lets define an interface for our persisted storage

// Storage is an interface for storing things
type Storage interface {
	Store(key string, bytes []byte) error
	Retrieve(key string) ([]byte, error)
}

// MapStorage is a map[string][]byte that implements the Storage interface
type MapStorage map[string][]byte

// Store saves an item
func (m MapStorage) Store(key string, bytes []byte) error {
	if _, taken := m[key]; taken {
		return errors.New("key taken")
	}
	m[key] = bytes
	return nil
}

// Retrieve gets data from Storage
func (m MapStorage) Retrieve(key string) ([]byte, error) {
	if _, avail := m[key]; avail {
		return m[key], nil
	}
	return nil, errors.New("key does not exist")
}

// Heres an implementation of storeage that uses disk

// DiskStorage implements storage for a folder on disk
type DiskStorage struct {
	BaseDir string
}

// Store saves an item
func (d *DiskStorage) Store(key string, bytes []byte) error {
	fp := path.Join(d.BaseDir, key)
	if _, err := os.Stat(fp); !os.IsNotExist(err) {
		return err
	}
	return ioutil.WriteFile(fp, bytes, 777)
}

// Retrieve gets data from Storage
func (d *DiskStorage) Retrieve(key string) ([]byte, error) {
	fp := path.Join(d.BaseDir, key)
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// And lets write function that uses the Storage interface

// StoreStringAsHash takes a string, hashes and stores it
func StoreStringAsHash(s string, storage Storage) error {
	hash := sha256Hash(s)
	return storage.Store(s, hash)
}

// helper function for sha256 hashing
func sha256Hash(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

// RetrieveStringHash gets a string hash from storage
func RetrieveStringHash(s string, storage Storage) ([]byte, error) {
	b, err := storage.Retrieve(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {

	log.Println("Using memory")
	mapstore := MapStorage{}
	s := "Hello, World"

	err := StoreStringAsHash(s, mapstore)
	if err != nil {
		log.Panic(err)
	}

	b, err := RetrieveStringHash(s, mapstore)
	if err != nil {
		log.Panic(err)
	}
	log.Println(s, b)

	log.Println("Using disk")

	diskstore := &DiskStorage{BaseDir: "./tmp"}
	err = StoreStringAsHash(s, diskstore)
	if err != nil {
		log.Panic(err)
	}

	b, err = RetrieveStringHash(s, diskstore)
	if err != nil {
		log.Panic(err)
	}
	log.Println(s, b)

}
