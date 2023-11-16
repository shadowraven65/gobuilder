package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

// Target represents a build target in the config.
type Target struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

// Config represents the JSON structure of the config file.
type Config struct {
	Targets []Target `json:"targets"`
}

func main() {
    // Read the config file.
    data, err := ioutil.ReadFile("buildconfig.json")
    if err != nil {
        fmt.Printf("Error reading config file: %v\n", err)
        return
    }

    // Parse the JSON data.
    var config Config
    err = json.Unmarshal(data, &config)
    if err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        return
    }

    goFile := flag.String("file", "", "specific Go file to build (optional)")
    flag.Parse()

    var scriptName string

    if *goFile == "" {
        // Find the .go file in the current directory if no specific file is provided.
        scriptName, err = findGoFile() // Removed 'var' keyword
        if err != nil {
            fmt.Println(err)
            return
        }
    } else {
        // Use the provided file name, removing the .go extension.
        scriptName = strings.TrimSuffix(filepath.Base(*goFile), ".go")
    }

    // Iterate over the targets and build for each.
    for _, target := range config.Targets {
        fmt.Printf("Building for OS: %s, Arch: %s", target.OS, target.Arch)
        err = buildTarget(scriptName, target) // Removed 'var' keyword
        if err != nil {
            fmt.Printf("Error building for target %v: %v", target, err)
        }
    }
}

func findGoFile() (string, error) {
    files, err := ioutil.ReadDir(".")
    if err != nil {
        return "", err
    }
    for _, file := range files {
        if filepath.Ext(file.Name()) == ".go" {
            return strings.TrimSuffix(file.Name(), ".go"), nil
        }
    }
    return "", fmt.Errorf("no .go file found in the current directory")
}

// buildTarget runs the go build command for the given target.
func buildTarget(scriptName string, target Target) error {
    // Define the output binary name format: scriptname_os_arch
    outputBinary := fmt.Sprintf("%s_%s_%s", scriptName, target.OS, target.Arch)
    // Building the command with -trimpath and -o flags
    cmd := exec.Command("go", "build", "-trimpath", "-o", outputBinary)
    cmd.Env = append(os.Environ(), "GOOS="+target.OS, "GOARCH="+target.Arch)
    // Executing the command
    output, err := cmd.CombinedOutput()
    fmt.Println(string(output))
    return err
}


