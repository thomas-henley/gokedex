package main

import "testing"

func Test_commandMap(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		{
			"basic",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := commandMap(nil)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("commandMap() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("commandMap() succeeded unexpectedly")
			}
		})
	}
}
