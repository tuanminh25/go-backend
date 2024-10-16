package folder

import (
	"fmt"
	"strings"
)


func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Find the source and destination folders
	var sourceFolder, destFolder Folder
	sourceFound, destFound := false, false

	for _, folder := range f.folders {
		if folder.Name == name {
			sourceFolder = folder
			sourceFound = true
		}
		if folder.Name == dst {
			destFolder = folder
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

	// Update and add the moved folders
	movedFolders := make([]Folder, 0)
	for _, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, oldPath) {
			updatedFolder := folder
			updatedFolder.Paths = strings.Replace(folder.Paths, oldPath, newPath, 1)
			movedFolders = append(movedFolders, updatedFolder)
		}
	}

	// Append moved folders to the end of updatedFolders
	updatedFolders = append(updatedFolders, movedFolders...)

	f.folders = updatedFolders
	return f.folders, nil
}
