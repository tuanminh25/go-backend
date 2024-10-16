package folder_test

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func getSampleFolders(orgID1, orgID2 uuid.UUID) []folder.Folder {
	return []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: orgID1},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
		{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
		{Name: "echo", Paths: "echo", OrgId: orgID1},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
	}
}
