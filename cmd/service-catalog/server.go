/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/golang/glog"
	// set up logging the k8s way
	"k8s.io/kubernetes/pkg/util/logs"

	"github.com/kubernetes-incubator/service-catalog/pkg/cmd/server"
	// TODO: may be necessary
	_ "k8s.io/kubernetes/pkg/api/install"
	// install our API groups
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
)

func main() {
	logs.InitLogs()
	// make sure we print all the logs while shutting down.
	defer logs.FlushLogs()

	cmd := server.NewCommandServer(os.Stdout)
	if err := cmd.Execute(); err != nil {
		glog.Errorln(err)
		logs.FlushLogs()
		os.Exit(1)
	}
}
