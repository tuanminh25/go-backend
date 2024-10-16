package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	tests := []struct {
		name        string
		folders     []folder.Folder
		sourceName  string
		destName    string
		want        []folder.Folder
		expectError bool
	}{
		{
			name: "Valid move",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "charlie", OrgId: orgID1},
			},
			sourceName: "bravo",
			destName:   "charlie",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "charlie", Paths: "charlie", OrgId: orgID1},
				{Name: "bravo", Paths: "charlie.bravo", OrgId: orgID1},
			},
			expectError: false,
		},
		{
			name: "Move to non-existent destination",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
			},
			sourceName:  "bravo",
			destName:    "charlie",
			want:        nil,
			expectError: true,
		},
		{
			name: "Move non-existent source",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
			},
			sourceName:  "charlie",
			destName:    "alpha",
			want:        nil,
			expectError: true,
		},
		{
			name: "Move to different organization",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "bravo", OrgId: orgID2},
			},
			sourceName:  "alpha",
			destName:    "bravo",
			want:        nil,
			expectError: true,
		},
		{
			name: "Move to itself",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
			},
			sourceName:  "alpha",
			destName:    "alpha",
			want:        nil,
			expectError: true,
		},
		{
			name: "Move to child of itself",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
			},
			sourceName:  "alpha",
			destName:    "bravo",
			want:        nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.sourceName, tt.destName)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
