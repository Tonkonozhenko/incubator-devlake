/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
	"time"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.MigrationScript = (*addCicdRelease)(nil)

type cicdRelease20240514 struct {
	archived.DomainEntity
	PublishedAt time.Time `json:"publishedAt"`

	CicdScopeId string `gorm:"index;type:varchar(255)"`

	Name         string `gorm:"type:varchar(255)"`
	DisplayTitle string `gorm:"type:varchar(255)"`
	Description  string `json:"description"`
	URL          string `json:"url"`

	IsDraft      bool `json:"isDraft"`
	IsLatest     bool `json:"isLatest"`
	IsPrerelease bool `json:"isPrerelease"`

	AuthorID string `json:"id" gorm:"type:varchar(255)"`

	RepoId string `gorm:"type:varchar(255)"`

	TagName string `json:"tagName"`
}

func (cicdRelease20240514) TableName() string {
	return "cicd_releases"
}

type addCicdRelease struct{}

func (*addCicdRelease) Up(basicRes context.BasicRes) errors.Error {
	return basicRes.GetDal().AutoMigrate(cicdRelease20240514{})
}

func (*addCicdRelease) Version() uint64 {
	return 20240514181200
}

func (*addCicdRelease) Name() string {
	return "add cicd_releases table"
}
