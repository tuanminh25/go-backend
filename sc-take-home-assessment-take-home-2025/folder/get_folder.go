package folder

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders
	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Check if orgID exists
	foldersInOrg := f.GetFoldersByOrgID(orgID)
	if len(foldersInOrg) == 0 {
		fmt.Println("Error: orgID does not exist")
		return []Folder{}
	}

	// Find the parent folder
	var parentFolder Folder
	parentFound := false
	for _, folder := range foldersInOrg {
		if folder.Name == name {
			parentFolder = folder
			parentFound = true
			break
		}
	}

	if !parentFound {
		fmt.Println("Error: Folder does not exist")
		return []Folder{}
	}

	// Find all child folders
	childFolders := []Folder{}
	parentPath := parentFolder.Paths + "."
	for _, folder := range foldersInOrg {
		if folder.Paths != parentFolder.Paths && strings.HasPrefix(folder.Paths, parentPath) {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders
}
