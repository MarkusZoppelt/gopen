package applications

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetInstalledApplications() []string {
	var applications []string

	apps, err := filepath.Glob("/Applications/*.app")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	apps, err = filepath.Glob("/System/Applications/*.app")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	apps, err = filepath.Glob(home + "/Applications/*.app")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	apps, err = filepath.Glob("/bin/*")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	apps, err = filepath.Glob("/usr/bin/*")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	apps, err = filepath.Glob("/usr/local/bin/*")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	apps, err = filepath.Glob("/opt/homebrew/bin/*")
	if err != nil {
		log.Fatal(err)
	}
	applications = append(applications, apps...)

	return applications
}

// openApplication opens the application with the given filepth
// TODO; support Windows
func OpenWithApplication(application string, filepaths []string, verbose bool) ([]byte, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		if strings.HasSuffix(application, ".app") {
			files := strings.Join(filepaths, " ")
			cmd = exec.Command("open", "-a", application, files)
		} else {
			cmd = exec.Command(application, filepaths...)
		}
	case "Linux":
		cmd = exec.Command(application, filepaths...)
	case "Windows":
		log.Printf("gopen is not supported on %s", runtime.GOOS)
		cmd = exec.Command("start", filepaths...)
	}

	if verbose {
		log.Printf("executing command: %s", cmd.String())
	}

	return cmd.Output()
}
