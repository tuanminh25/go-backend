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
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName: "bravo",
			destName:   "delta",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
				{Name: "golf", Paths: "golf", OrgId: orgID1},
			},
			expectError: false,
		},
		{
			name: "Valid move 2",
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName: "bravo",
			destName:   "golf",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "golf.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
				{Name: "golf", Paths: "golf", OrgId: orgID1},
			},
			expectError: false,
		},
		{
			name: "Invalid move: move to child of itself",
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName:  "bravo",
			destName:    "charlie",
			want:        nil,
			expectError: true,
		},
		{
			name: "Invalid move: Move to itself",
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName:  "bravo",
			destName:    "bravo",
			want:        nil,
			expectError: true,
		},
		{
			name: "Invalid move: Move to different organization",
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName:  "bravo",
			destName:    "foxtrot",
			want:        nil,
			expectError: true,
		},
		{
			name: "Invalid move: Move non-existent source",
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName:  "invalid_folder",
			destName:    "delta",
			want:        nil,
			expectError: true,
		},
		{
			name: "Move to non-existent destination",
			folders:    getSampleFolders2(orgID1, orgID2),
			sourceName:  "bravo",
			destName:    "invalid_folder",
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
				// Sort both the expected and actual results
				sortFolders(tt.want)
				sortFolders(got)
				
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
