package vsphere

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFolderOptions(t *testing.T) {
	tests := []struct {
		name          string
		inventoryPath string
		prefix        string
		want          []string
	}{
		{
			name:          "returns absolute and relative for nested folder",
			inventoryPath: "/dc-1/vm/folder1/subFolder1",
			prefix:        "/dc-1/vm",
			want:          []string{"/dc-1/vm/folder1/subFolder1", "folder1/subFolder1"},
		},
		{
			name:          "returns absolute and relative for direct folder",
			inventoryPath: "/dc-1/vm/folder1",
			prefix:        "/dc-1/vm",
			want:          []string{"/dc-1/vm/folder1", "folder1"},
		},
		{
			name:          "returns nil when path outside vm prefix",
			inventoryPath: "/dc-1/host/cluster-a",
			prefix:        "/dc-1/vm",
			want:          nil,
		},
		{
			name:          "returns nil for base vm folder",
			inventoryPath: "/dc-1/vm",
			prefix:        "/dc-1/vm",
			want:          nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, folderOptions(tt.inventoryPath, tt.prefix))
		})
	}
}
