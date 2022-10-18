package main

import "testing"

func TestGetEnvVar(t *testing.T) {
	t.Run("check not defined env var", func(t *testing.T) {
		got := getEnvVar("random_text_value_please_do_not_set_env")
		want := defaultEnvValue

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("check app name", func(t *testing.T) {
		got := getEnvVar("DEVOPSTESTE_NAME")
		want := "Amo Promo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
