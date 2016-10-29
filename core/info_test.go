package core

import "testing"

func TestGetVersionDisplay(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Display Version",
			want: ProductName + " version " + Version + "-" + VersionPrerelease + "\n",
		},
	}
	for _, tt := range tests {
		if got := GetVersionDisplay(); got != tt.want {
			t.Errorf("%q. GetVersionDisplay() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getHumanVersion(t *testing.T) {
	GitCommit = "f3b4a47+CHANGES"
	GitDescribe = "f3b4a47"
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Display Git variables defined",
			want: "f3b4a47-dev (f3b4a47+CHANGES)",
		},
	}
	for _, tt := range tests {
		if got := getHumanVersion(); got != tt.want {
			t.Errorf("%q. getHumanVersion() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
