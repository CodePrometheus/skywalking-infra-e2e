// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cleanup

import (
	"os"
	"strings"
	"time"

	kind "sigs.k8s.io/kind/cmd/kind/app"
	kindcmd "sigs.k8s.io/kind/pkg/cmd"

	"github.com/apache/skywalking-infra-e2e/internal/config"
	"github.com/apache/skywalking-infra-e2e/internal/constant"
	"github.com/apache/skywalking-infra-e2e/internal/logger"
	"github.com/apache/skywalking-infra-e2e/internal/util"
)

const (
	maxRetry      = 5
	retryInterval = 2 // in seconds
)

func KindCleanUp(e2eConfig *config.E2EConfig) error {
	kindConfigFilePath := e2eConfig.Setup.GetFile()

	logger.Log.Infof("deleting kind cluster...\n")
	if err := cleanKindCluster(kindConfigFilePath); err != nil {
		logger.Log.Error("delete kind cluster failed")
		return err
	}
	logger.Log.Info("delete kind cluster succeeded")

	kubeConfigPath := constant.K8sClusterConfigFilePath
	logger.Log.Infof("deleting k8s cluster config file:%s", kubeConfigPath)
	err := os.Remove(kubeConfigPath)
	if err != nil {
		logger.Log.Infoln("delete k8s cluster config file failed")
	}

	return nil
}

func cleanKindCluster(kindConfigFilePath string) (err error) {
	clusterName, err := util.GetKindClusterName(kindConfigFilePath)
	if err != nil {
		return err
	}

	args := []string{"delete", "cluster", "--name", clusterName}

	logger.Log.Debugf("cluster delete commands: %s %s", constant.KindCommand, strings.Join(args, " "))

	// Sometimes kind delete cluster failed, so we retry it.
	for i := 0; i < maxRetry; i++ {
		if err = kind.Run(kindcmd.NewLogger(), kindcmd.StandardIOStreams(), args); err == nil {
			return nil
		}
		time.Sleep(retryInterval * time.Second)
	}

	return
}
