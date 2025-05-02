// getArchive.go
// Creates a compressed file of the selected project.

package git

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	git "github.com/libgit2/git2go/v34"
)

// Creates a compressed file of the selected project at the selected reference.
// Returns the temporary file path.
func GetArchive(repositoryName string, reference string, extension string) (string, error) {
	REPOSITORIES_DIR := os.Getenv("REPOSITORIES_DIR")
	if REPOSITORIES_DIR == "" {
		return "", errors.New("missing environment variable REPOSITORIES_DIR")
	}

	repository, err := git.OpenRepository(REPOSITORIES_DIR + repositoryName + ".git")
	if err != nil {
		return "", err
	}

	// Search for reference hash.
	ref := reference
	branch, err := repository.LookupBranch(reference, git.BranchAll)
	if err == nil {
		ref = branch.Target().String()
	} else {
		tag, err := repository.References.Lookup("refs/tags/" + reference)
		if err == nil {
			ref = tag.Target().String()
		}
	}

	oid, err := git.NewOid(ref)
	if err != nil {
		return "", err
	}

	commit, err := repository.LookupCommit(oid)
	if err != nil {
		return "", err
	}

	tree, err := commit.Tree()
	if err != nil {
		return "", err
	}

	// Temporary compressed file location.
	tempFilePath := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%s.%s", repositoryName, reference, extension))

	// Create temporary compressed file.
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// Compressed file writters.
	gzWriter := gzip.NewWriter(tempFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	zipWriter := zip.NewWriter(tempFile)
	defer zipWriter.Close()

	// Get tree files.
	err = tree.Walk(func(name string, entry *git.TreeEntry) error {
		if entry.Type == git.ObjectBlob {
			// Get the blob for the entry.
			blob, err := repository.LookupBlob(entry.Id)
			if err != nil {
				return fmt.Errorf("failed to get blob: %w", err)
			}

			// Create tar.gz file.
			if extension == "tar.gz" {
				header := &tar.Header{
					Name: filepath.Join(name, entry.Name),
					Mode: 0600,
					Size: int64(len(blob.Contents())),
				}

				err = tarWriter.WriteHeader(header)
				if err != nil {
					return errors.New("while writing on tar file header, " + err.Error())
				}

				_, err = tarWriter.Write(blob.Contents())
				if err != nil {
					return errors.New("while writing tar.gz file, " + err.Error())
				}
			} else {
				// Create zip file.
				fileWriter, err := zipWriter.Create(filepath.Join(name, entry.Name))
				if err != nil {
					return errors.New("while creating the file writer, " + err.Error())
				}

				_, err = fileWriter.Write(blob.Contents())
				if err != nil {
					return errors.New("while writting file, " + err.Error())
				}
			}
		}

		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to walk tree: %w", err)
	}

	return tempFilePath, nil
}
