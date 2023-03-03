//  Copyright (c) 2022 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"go.fd.io/govpp/binapigen"
	"go.fd.io/govpp/binapigen/vppapi"
	"go.fd.io/govpp/version"

	"github.com/calico-vpp/vpplink/pkg/wrappergen"
)

const (
	generateLogFname = "generate.log"
)

var (
	//go:embed templates/*
	templates       embed.FS
	parsedTemplates *wrappergen.Template
)

func init() {
	// Trim off the "templates" prefix from the paths of our templates
	templates, err := fs.Sub(templates, "templates")
	if err != nil {
		logrus.Fatalf("error creating subFS for 'templates'")
	}

	// Parse all the templates
	parsedTemplates, err = wrappergen.ParseFS(templates, "*.tmpl.go")
	if err != nil {
		logrus.Fatalf("failed to ParseFS templates: %s", err)
	}

}

func GenerateAll(gen *binapigen.Generator) []*binapigen.GenFile {
	genOpts := gen.GetOpts()

	logrus.Infof("[WRAPPERGEN] GenerateAll (opts: %+v)", genOpts)

	// We output vpplink one directory higher than the regular bindings
	basePkgName := filepath.Join(genOpts.ImportPrefix, "..")
	outputDir := filepath.Join(genOpts.OutputDir, "..")

	data := wrappergen.NewDataFromFiles(genOpts.ImportPrefix, filepath.Base(basePkgName), gen.Files)

	// Execute all the templates
	err := parsedTemplates.ExecuteAll(outputDir, data, gen)
	if err != nil {
		logrus.Fatalf("failed to execute template: %s", err)
	}

	if vppDir := os.Getenv("VPP_DIR"); vppDir != "" {
		createGenerateLog(vppDir, filepath.Join(outputDir, generateLogFname))
	}

	return nil
}

func createGenerateLog(apiDir string, fname string) {
	vppSrcDir, err := findGitRepoRootDir(apiDir)
	if err != nil {
		return
	}

	vppVersion, err := vppapi.GetVPPVersionRepo(vppSrcDir)
	if err != nil {
		logrus.Fatalf("Unable to get vpp version : %s", err)
	}

	cmd := exec.Command("bash", "-c", "git log --oneline -1 $(git log origin/master..HEAD --oneline | tail -1 | awk '{print $1}')")
	cmd.Dir = vppSrcDir
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		logrus.Fatalf("Unable to get vpp base commit : %s", out)
	}
	lastCommit := strings.TrimSpace(string(out))

	cmd = exec.Command("git", "log", "origin/master..HEAD", "--pretty=%s")
	cmd.Dir = vppSrcDir
	cmd.Stderr = os.Stderr
	out, err = cmd.Output()
	if err != nil {
		logrus.Fatalf("Unable to get vpp own branch commits : %s", out)
	}
	ownCommits := strings.TrimSpace(string(out))

	value := fmt.Sprintf("VPP Version                 : %s\n", vppVersion)
	value += fmt.Sprintf("Binapi-generator version    : %s\n", version.Info())
	value += fmt.Sprintf("VPP Base commit             : %s\n", lastCommit)
	value += fmt.Sprintf("------------------ Cherry picked commits --------------------\n")
	value += fmt.Sprintf("%s\n", ownCommits)
	value += fmt.Sprintf("-------------------------------------------------------------\n")

	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		logrus.Fatalf("Unable to open file %s %s", fname, err)
	}
	n, err := f.Write([]byte(value))
	if err != nil || n < len(value) {
		logrus.Fatalf("Unable to write to file %s %s", fname, err)
	}
	err = f.Close()
	if err != nil {
		logrus.Fatalf("Unable to close file %s %s", fname, err)
	}

}

func findGitRepoRootDir(dir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git command failed: %v\noutput: %s", err, out)
	}
	return strings.TrimSpace(string(out)), nil
}
