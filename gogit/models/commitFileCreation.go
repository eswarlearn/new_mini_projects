package models

func commitFileCreation() {

}


// entries, err := os.ReadDir(".gogit/commit")
// if err != nil {
//     return fmt.Errorf("failed to read commit directory: %w", err)
// }

// lastFolder := 0

// for _, entry := range entries {
//     if !entry.IsDir() {
//         continue
//     }

//     n, err := strconv.Atoi(entry.Name())
//     if err == nil && n > lastFolder {
//         lastFolder = n
//     }
// }

// newFolder := filepath.Join(".gogit", "commit", strconv.Itoa(lastFolder+1))

// if err := os.MkdirAll(newFolder, 0755); err != nil {
//     return fmt.Errorf("failed to create commit folder: %w", err)
// }