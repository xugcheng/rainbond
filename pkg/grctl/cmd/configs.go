// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/goodrain/rainbond/pkg/grctl/clients"
	"github.com/urfave/cli"
)

//NewCmdConfigs 全局配置相关命令
func NewCmdConfigs() cli.Command {
	c := cli.Command{
		Name:  "configs",
		Usage: "系统任务相关命令，grctl tasks -h",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "get",
				Usage: "get all datacenter configs",
				Action: func(c *cli.Context) error {
					configs, err := clients.NodeClient.Configs().Get()
					if err != nil {
						return err
					}
					taskTable := termtables.CreateTable()
					taskTable.AddHeaders("Name", "CNName", "ValueType", "Value")
					for _, config := range configs.Configs {
						taskTable.AddRow(config.Name, config.CNName, config.ValueType, config.Value)
					}
					fmt.Println(taskTable.Render())
					return nil
				},
			},
		},
	}
	return c
}