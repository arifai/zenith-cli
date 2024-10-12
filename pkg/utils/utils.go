package utils

import (
	"fmt"
	"github.com/arifai/zenith-cli/pkg/printer"
	"golang.org/x/mod/modfile"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

func CheckGoModFileExists() bool {
	_, err := os.Stat("go.mod")
	return !os.IsNotExist(err)
}

func ConvertCase(str string, capitalizeFirst bool) string {
	var result strings.Builder
	upperNext := capitalizeFirst

	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if r == '_' || r == ' ' {
				upperNext = true
			}
			continue
		}
		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func ToSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		if r == ' ' || r == '_' {
			result.WriteRune('_')
		} else if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(rune(str[i-1])) || str[i-1] == '_') {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func UpdateGoModAndImports(repoUrl, clonePath, moduleName, orgName string) error {
	if err := updateGoMod(clonePath, moduleName, orgName); err != nil {
		return fmt.Errorf("failed to update go.mod: %v", err)
	}

	newModuleName := constructModuleName(moduleName, orgName)

	if err := updateImports(clonePath, repoUrl, newModuleName); err != nil {
		return fmt.Errorf("failed to update import paths: %v", err)
	}

	if err := os.Chdir(clonePath); err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	if err := RunCommand("go", "mod", "tidy"); err != nil {
		return fmt.Errorf("failed to run go mod tidy: %v", err)
	} else {
		printer.Green("✨ Successfully running go mod tidy.")
	}

	return nil
}

func constructModuleName(moduleName, orgName string) string {
	if orgName != "" {
		return fmt.Sprintf("%s/%s", strings.TrimSuffix(orgName, "/"), moduleName)
	}
	return moduleName
}
func updateGoMod(clonePath, moduleName, orgName string) error {
	modFile, goModPath, err := readModFile(clonePath)
	if err != nil {
		return fmt.Errorf("failed to read go.mod: %v", err)
	}

	newModuleName := constructModuleName(moduleName, orgName)
	if err := modFile.AddModuleStmt(newModuleName); err != nil {
		return fmt.Errorf("failed to update module statement: %v", err)
	}

	newGoModContent, err := modFile.Format()
	if err != nil {
		return fmt.Errorf("failed to format go.mod: %v", err)
	}

	if err := os.WriteFile(goModPath, newGoModContent, 0644); err != nil {
		return fmt.Errorf("failed to write updated go.mod: %v", err)
	}

	return nil
}

func updateImports(clonePath, repoUrl, newModuleName string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command",
			"Get-ChildItem -Path '"+clonePath+"' -Filter '*.go' -Recurse | ForEach-Object { (Get-Content $_.FullName) "+
				"-replace '"+repoUrl+"', '"+newModuleName+"' | Set-Content $_.FullName }")
	} else {
		cmd = exec.Command("find", clonePath, "-name", "*.go", "-exec", "sed", "-i", "",
			"s|"+repoUrl+"|"+newModuleName+"|g", "{}", ";")
	}
	return cmd.Run()
}

func RunCommand(name string, args ...string) error {
	return exec.Command(name, args...).Run()
}

func GoFmt() error {
	printer.Yellow("📝 Formatting all Go files...")

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			if err := RunCommand("gofmt", "-w", path); err != nil {
				return fmt.Errorf("failed to format file %s: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	printer.Green("✨ Successfully formatted all Go files.")
	return nil
}

func readModFile(clonePath string) (*modfile.File, string, error) {
	goModPath := filepath.Join(clonePath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return nil, "", err
	}

	modFile, err := modfile.Parse(goModPath, content, nil)
	if err != nil {
		return nil, "", err
	}

	return modFile, goModPath, nil
}

func GetModuleName(clonePath string) (string, error) {
	modFile, _, err := readModFile(clonePath)
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %v", err)
	}

	if modFile.Module == nil {
		return "", fmt.Errorf("no module name found in go.mod")
	}

	return modFile.Module.Mod.Path, nil
}
