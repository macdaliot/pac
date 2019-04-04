package ui

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

var defaultDirectory = "c:\\users\\mmcnairy"
var projectDirectory = "c:\\users\\mmcnairy\\go\\src\\testacular"

func start() {
	r := mux.NewRouter()
	r.HandleFunc("/pac/setup", pacSetup)
	r.HandleFunc("/jenkins/rebuild", rebuildJenkins)
	r.HandleFunc("/", handler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home page."))
}

func pacSetup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	os.Chdir(defaultDirectory)
	exec.Command("pac", "setup", "--name", vars["project-name"])
}

func rebuildJenkins(w http.ResponseWriter, r *http.Request) {
	os.Chdir(projectDirectory + "\\terraform")
	exec.Command("terraform", "taint", "aws_ecs_task_defition.jenkins")
	exec.Command("terraform", "apply -auto-approve")

	w.WriteHeader(http.StatusOK)
	fmt.Print("Jenkins is restarting. Give it a minute or two.")
}
