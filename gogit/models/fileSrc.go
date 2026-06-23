package models
import(
	"os"
)

func getFilesName() ([]string, error) {
	entries, err := os.ReadDir(".gogit/staging")
	if err != nil {
		return nil, err
	}
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}