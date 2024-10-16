package folder

import (
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Find the source and destination folders
	var sourceFolder, destFolder Folder
	var destIndex int
	sourceFound, destFound := false, false

	for i, folder := range f.folders {
		if folder.Name == name {
			sourceFolder = folder
			sourceFound = true
		}
		if folder.Name == dst {
			destFolder = folder
			destIndex = i
			destFound = true
		}
		if sourceFound && destFound {
			break
		}
	}

	// Error handling
	if !sourceFound {
		return nil, fmt.Errorf("Error: Source folder does not exist")
	}
	if !destFound {
		return nil, fmt.Errorf("Error: Destination folder does not exist")
	}
	if sourceFolder.OrgId != destFolder.OrgId {
		return nil, fmt.Errorf("Error: Cannot move a folder to a different organization")
	}
	if name == dst {
		return nil, fmt.Errorf("Error: Cannot move a folder to itself")
	}
	if strings.HasPrefix(destFolder.Paths, sourceFolder.Paths) {
		return nil, fmt.Errorf("Error: Cannot move a folder to a child of itself")
	}

	// Perform the move
	oldPath := sourceFolder.Paths
	newPath := destFolder.Paths + "." + sourceFolder.Name

	// Create a new slice to store the updated folders
	updatedFolders := make([]Folder, 0, len(f.folders))

	// First, add all folders that are not being moved
	for _, folder := range f.folders {
		if !strings.HasPrefix(folder.Paths, oldPath) {
			updatedFolders = append(updatedFolders, folder)
		}
	}

	// Find the index of the destination folder in the updated slice
	destIndex = -1
	for i, folder := range updatedFolders {
		if folder.Name == dst {
			destIndex = i
			break
		}
	}

	if destIndex == -1 {
		return nil, fmt.Errorf("Error: Destination folder not found in updated list")
	}

	// Insert the moved folders after the destination folder
	movedFolders := make([]Folder, 0)
	for _, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, oldPath) {
			updatedFolder := folder
			updatedFolder.Paths = strings.Replace(folder.Paths, oldPath, newPath, 1)
			movedFolders = append(movedFolders, updatedFolder)
		}
	}

	// Insert moved folders after the destination folder
	updatedFolders = append(updatedFolders[:destIndex+1], append(movedFolders, updatedFolders[destIndex+1:]...)...)

	f.folders = updatedFolders
	return f.folders, nil
}
