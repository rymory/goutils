// Copyright (c) 2017-2026 Onur Yaşar
// Licensed under AGPL v3 + Commercial Exception
// See LICENSE.txt

// https://github.com/rymory/goutils/blob/main/LICENSE
// https://github.com/rymory/goutils/blob/main/README.md
// rymory.org
// onuryasar.org
// onxorg@proton.me 

package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetRoleLevel(roleId int) int {
	roleLevel, err := strconv.Atoi((fmt.Sprintf("%c", fmt.Sprint(roleId)[0])))
	if err != nil {
		log.Fatal(err)
	}
	return roleLevel
}

func GetProjectId(projectInfo string) int {
	projectId, err := strconv.Atoi(strings.Trim(strings.Split(projectInfo, ":")[0], " "))
	if err != nil {
		log.Fatal(err)
	}
	return projectId
}
