package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	tests := []struct {
		name       string
		orgID      uuid.UUID
		folderName string
		folders    []folder.Folder
		want       []folder.Folder
	}{
		{
			name:       "Valid folder with children",
			orgID:      orgID1,
			folderName: "alpha",
			folders:    getSampleFolders(orgID1, orgID2),
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
			},
		},
		{
			name:       "Valid folder without children",
			orgID:      orgID1,
			folderName: "echo",
			folders:    getSampleFolders(orgID1, orgID2),
			want: []folder.Folder{},
		},
		{
			name:       "Invalid organization ID",
			orgID:      uuid.Must(uuid.NewV4()),
			folderName: "alpha",
			folders:    getSampleFolders(orgID1, orgID2),
			want:       []folder.Folder{},
		},
		{
			name:       "Invalid folder name",
			orgID:      orgID1,
			folderName: "invalid_folder",
			folders:    getSampleFolders(orgID1, orgID2),
			want:       []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got := f.GetAllChildFolders(tt.orgID, tt.folderName)
			// Sort both the expected and actual results
			sortFolders(tt.want)
			sortFolders(got)
			assert.Equal(t, tt.want, got)
		})
	}
}
