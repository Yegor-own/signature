package config

import "testing"

func TestloadEnv(t *testing.T) {
	_, err := loadEnv()
	if err != nil {
		t.Error(err)
	}
}
