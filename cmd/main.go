package main

import "github.com/arifai/zenith-cli/cmd/command"

//const (
//	Reset  = "\033[0m"
//	Red    = "\033[31m"
//	Green  = "\033[32m"
//	Yellow = "\033[33m"
//)
//
//const version = "1.0.0"
//
//type TemplateData struct {
//	Prefix string
//}
//
//// Function to convert input to lowercase and snake_case
//func toSnakeCase(str string) string {
//	var result strings.Builder
//	for i, r := range str {
//		if r == ' ' {
//			result.WriteRune('_')
//		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
//			continue
//		} else if unicode.IsUpper(r) {
//			if i > 0 && result.Len() > 0 && result.String()[result.Len()-1] != '_' {
//				result.WriteRune('_')
//			}
//			result.WriteRune(unicode.ToLower(r))
//		} else {
//			result.WriteRune(r)
//		}
//	}
//	return result.String()
//}
//
//// Function to convert input to camelCase without special characters
//func toCamelCase(str string) string {
//	var result strings.Builder
//	upperNext := true
//	for _, r := range str {
//		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
//			continue
//		}
//		if r == '_' || r == ' ' {
//			upperNext = true
//			continue
//		}
//		if upperNext {
//			result.WriteRune(unicode.ToUpper(r))
//			upperNext = false
//		} else {
//			result.WriteRune(unicode.ToLower(r))
//		}
//	}
//	return result.String()
//}
//
//// Function to write default content to generated files
//func writeFileContent(filePath, prefix string) error {
//	fileContent := `package main
//
//func {{.Prefix | toCamelCase}}Repository() {
//    return
//}
//`
//	tmpl, err := template.New("fileContent").Funcs(template.FuncMap{
//		"toCamelCase": toCamelCase,
//	}).Parse(fileContent)
//	if err != nil {
//		return fmt.Errorf("failed to parse file content template: %w", err)
//	}
//
//	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//	if err != nil {
//		return fmt.Errorf("failed to open file %s: %w", filePath, err)
//	}
//	defer file.Close()
//
//	data := TemplateData{Prefix: prefix}
//	if err := tmpl.Execute(file, data); err != nil {
//		return fmt.Errorf("failed to execute template for file %s: %w", filePath, err)
//	}
//	return nil
//}
//
//// Function to create a directory structure and generate files with default content
//func createDirStructure(basePath, prefix string) error {
//	internalDirPath := filepath.Join(basePath, "internal")
//
//	if _, err := os.Stat(internalDirPath); os.IsNotExist(err) {
//		if err := os.MkdirAll(internalDirPath, os.ModePerm); err != nil {
//			return fmt.Errorf("failed to create directory %s: %w", internalDirPath, err)
//		}
//		fmt.Printf("%sCreated directory: %s%s\n", Green, internalDirPath, Reset)
//	} else {
//		fmt.Printf("%sDirectory already exists: %s%s\n", Yellow, internalDirPath, Reset)
//	}
//
//	snakeCasePrefix := toSnakeCase(prefix)
//
//	structure := map[string][]string{
//		"api/handler": {
//			"handler.go",
//		},
//		"api/router": {
//			"router.go",
//		},
//		"api/types": {
//			"request.go",
//		},
//		"domain/model": {
//			snakeCasePrefix + ".go",
//			snakeCasePrefix + "_migration.go",
//		},
//		"domain/repository": {
//			snakeCasePrefix + "_repository.go",
//		},
//		"domain/service": {
//			snakeCasePrefix + "_auth_service.go",
//			snakeCasePrefix + "_service.go",
//		},
//	}
//
//	mainDirPath := filepath.Join(internalDirPath, snakeCasePrefix)
//	if _, err := os.Stat(mainDirPath); os.IsNotExist(err) {
//		if err := os.MkdirAll(mainDirPath, os.ModePerm); err != nil {
//			return fmt.Errorf("failed to create directory %s: %w", mainDirPath, err)
//		}
//		fmt.Printf("%sCreated directory: %s%s\n", Green, mainDirPath, Reset)
//	} else {
//		fmt.Printf("%sDirectory already exists: %s%s\n", Yellow, mainDirPath, Reset)
//	}
//
//	for dir, files := range structure {
//		dirPath := filepath.Join(mainDirPath, dir)
//		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
//			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
//				return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
//			}
//			fmt.Printf("%sCreated directory: %s%s\n", Green, dirPath, Reset)
//		} else {
//			fmt.Printf("%sDirectory already exists: %s%s\n", Yellow, dirPath, Reset)
//		}
//
//		for _, fileName := range files {
//			filePath := filepath.Join(dirPath, fileName)
//
//			if _, err := os.Stat(filePath); os.IsNotExist(err) {
//				if err := writeFileContent(filePath, prefix); err != nil {
//					return err
//				}
//				fmt.Printf("%sCreated file: %s%s\n", Green, filePath, Reset)
//			} else {
//				fmt.Printf("%sFile already exists: %s%s\n", Yellow, filePath, Reset)
//			}
//		}
//	}
//	return nil
//}
//
//// Function to display usage and version info
//func displayUsage() {
//	fmt.Printf("%sZen CLI Tool %s%s\n", Green, version, Reset)
//	fmt.Println("Usage:")
//	fmt.Printf("  %szen make NAMA_PREFIX%s  - Create directory structure with NAMA_PREFIX\n", Yellow, Reset)
//	fmt.Println("  zen version           - Display current version")
//}

func main() {
	command.Execute()
	//// Check if no arguments are provided
	//if len(os.Args) == 1 {
	//	displayUsage()
	//	return
	//}
	//
	//// Check if "version" is called
	//if len(os.Args) == 2 && os.Args[1] == "version" {
	//	fmt.Printf("%sZen CLI Tool version: %s%s\n", Green, version, Reset)
	//	return
	//}
	//
	//// Check if the number of arguments is correct for "make"
	//if len(os.Args) != 3 {
	//	fmt.Printf("%sUsage: zen make NAMA_PREFIX%s\n", Yellow, Reset)
	//	return
	//}
	//
	//// Check if the command is "make"
	//command := os.Args[1]
	//if command != "make" {
	//	fmt.Printf("%sError: Unknown command '%s'. Use 'make'.%s\n", Red, command, Reset)
	//	return
	//}
	//
	//// Get the prefix argument and convert to snake_case
	//prefix := toSnakeCase(os.Args[2])
	//
	//// Check if go.mod exists in the root directory
	//if !checkGoModFileExists() {
	//	fmt.Printf("%sError: go.mod file does not exist in the root directory. Please run `go mod init` first.%s\n", Red, Reset)
	//	return
	//}
	//
	//// Create the directory and file structure with templates
	//if err := createDirStructure(".", prefix); err != nil {
	//	fmt.Printf("%sError: %v%s\n", Red, err, Reset)
	//} else {
	//	fmt.Printf("%sTemplate structure created successfully!%s\n", Green, Reset)
	//}
}
