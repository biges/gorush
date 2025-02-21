package boltdb

import (
	"log"

	"github.com/biges/gorush/config"
	"github.com/biges/gorush/storage"

	"github.com/asdine/storm"
)

// New func implements the storage interface for gorush (https://github.com/biges/gorush)
func New(config config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config config.ConfYaml
}

// Init client storage.
func (s *Storage) Init() error {
	return nil
}

// Reset Client storage.
func (s *Storage) Reset() {
	s.setBoltDB(storage.TotalCountKey, 0)
	s.setBoltDB(storage.IosSuccessKey, 0)
	s.setBoltDB(storage.IosErrorKey, 0)
	s.setBoltDB(storage.AndroidSuccessKey, 0)
	s.setBoltDB(storage.AndroidErrorKey, 0)
}

func (s *Storage) setBoltDB(key string, count int64) {
	db, _ := storm.Open(s.config.Stat.BoltDB.Path)
	err := db.Set(s.config.Stat.BoltDB.Bucket, key, count)
	if err != nil {
		log.Println("BoltDB set error:", err.Error())
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BoltDB error:", err.Error())
		}
	}()
}

func (s *Storage) getBoltDB(key string, count *int64) {
	db, _ := storm.Open(s.config.Stat.BoltDB.Path)
	err := db.Get(s.config.Stat.BoltDB.Bucket, key, count)
	if err != nil {
		log.Println("BoltDB get error:", err.Error())
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BoltDB error:", err.Error())
		}
	}()
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	s.setBoltDB(storage.TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	s.setBoltDB(storage.IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	s.setBoltDB(storage.IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	s.setBoltDB(storage.AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	s.setBoltDB(storage.AndroidErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getBoltDB(storage.TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getBoltDB(storage.IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	s.getBoltDB(storage.IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getBoltDB(storage.AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getBoltDB(storage.AndroidErrorKey, &count)

	return count
}
