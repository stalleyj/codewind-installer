/*******************************************************************************
 * Copyright (c) 2019 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/eclipse/codewind-installer/errors"
	"github.com/eclipse/codewind-installer/utils"
	"github.com/urfave/cli"
)

type (
	// ProjectType represents the information Codewind requires to build a project.
	ProjectType struct {
		Language  string `json:"language"`
		BuildType string `json:"buildType"`
	}

	// ValidationResponse represents the response to validating a project on the users filesystem.
	ValidationResponse struct {
		Status string      `json:"status"`
		Path   string      `json:"path"`
		Result ProjectType `json:"result"`
	}
)

// DownloadTemplate using the url/link provided
func DownloadTemplate(c *cli.Context) {
	destination := c.Args().Get(0)

	if destination == "" {
		log.Fatal("destination not set")
	}

	url := c.String("u")

	err := utils.DownloadFromURLThenExtract(url, destination)
	if err != nil {
		log.Fatal(err)
	}
}

// ValidateProject returns the language and buildType for a project at given filesystem path,
// and writes a default .cw-settings file to that project
func ValidateProject(c *cli.Context) {
	projectPath := c.Args().Get(0)
	utils.CheckProjectPath(projectPath)

	language, buildType := utils.DetermineProjectInfo(projectPath)
	response := ValidationResponse{
		Status: "success",
		Result: ProjectType{language, buildType},
		Path:   projectPath,
	}
	projectInfo, err := json.Marshal(response)

	errors.CheckErr(err, 203, "")
	writeCwSettingsIfNotInProject(projectPath, buildType)
	fmt.Println(string(projectInfo))
}

func writeCwSettingsIfNotInProject(projectPath string, BuildType string) {
	pathToCwSettings := path.Join(projectPath, ".cw-settings")
	pathToLegacySettings := path.Join(projectPath, ".mc-settings")

	if _, err := os.Stat(pathToLegacySettings); os.IsExist(err) {
		utils.RenameLegacySettings(pathToLegacySettings, pathToCwSettings)
	} else if _, err := os.Stat(pathToCwSettings); os.IsNotExist(err) {
		utils.WriteNewCwSettings(pathToCwSettings, BuildType)
	}
}
