package main

import (
	"fmt"
	"os"

	"github.com/progrium/go-shell"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

type (
	Repo struct {
		Owner   string
		Name    string
		Link    string
		Avatar  string
		Branch  string
		Private bool
		Trusted bool
	}

	Build struct {
		Number   int
		Event    string
		Status   string
		Deploy   string
		Created  int64
		Started  int64
		Finished int64
		Link     string
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Commit struct {
		Remote  string
		Sha     string
		Ref     string
		Link    string
		Branch  string
		Message string
		Author  Author
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
	}

	Config struct {
		SparkDeployMode                                string
		SparkKubernetesLocalDeploy                     string
		SparkKubernetesLocalUrl                        string
		SparkKubernetesLocalPort                       string
		SparkClass                                     string
		SparkMaster                                    string
		SparkKubernetesNamespace                       string
		SparkAppName                                   string
		SparkLocalDir                                  string
		KubernetesDriverDockerImage                    string
		KubernetesExecutorDockerImage                  string
		KubernetesInitContainerDockerImage             string
		SparkDynamicAllocationEnabled                  string
		KubernetesResourceStagingServerUri             string
		KubernetesResourceStagingServerInternalUri     string
		SparkShuffleServiceEnabled                     string
		SparkKubernetesShuffleNamespace                string
		SparkKubernetesShuffleLabels                   string
		KubernetesAuthenticateDriverServiceAccountName string
		KubernetesAuthenticateSubmissionCaCertFile     string
		KubernetesAuthenticateSubmissionClientCertFile string
		KubernetesAuthenticateSubmissionClientKeyFile  string
		SparkMetricsConf                               string
		SparkPackages                                  string
		SparkExcludePackages                           string
		SparkAppSource                                 string
		SparkAppArgs                                   string
	}
)

var validate *validator.Validate

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	shell.Trace = true
	shell.Shell = []string{"/bin/bash", "-c"}

}

func (p *Plugin) Exec() error {
	validate = validator.New()
	err := validate.Struct(p)

	//&p.Config.SparkMaster
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Printf("[%s] field validation error (%+v)\n", v.Field(), v)
		}
		return nil
	}
	var clintCertAuth string
	if p.Config.SparkKubernetesLocalDeploy == "true" {
		if p.Config.SparkKubernetesLocalUrl == "" || p.Config.SparkKubernetesLocalPort == "" {
			log.Panicf("Kubernetes api endpoints is missing! URL: %q, Port: %q",
				p.Config.SparkKubernetesLocalUrl,
				p.Config.SparkKubernetesLocalPort)
		}
		p.Config.SparkMaster = fmt.Sprintf("k8s://https://%s:%s",
			p.Config.SparkKubernetesLocalUrl,
			p.Config.SparkKubernetesLocalPort,
		)
	} else {
		clintCertAuth = fmt.Sprintf(" --conf spark.kubernetes.authenticate.submission.caCertFile=%s "+
			"--conf spark.kubernetes.authenticate.submission.clientCertFile=%s "+
			"--conf spark.kubernetes.authenticate.submission.clientKeyFile=%s ",
			p.Config.KubernetesAuthenticateSubmissionCaCertFile,
			p.Config.KubernetesAuthenticateSubmissionClientCertFile,
			p.Config.KubernetesAuthenticateSubmissionClientKeyFile,
		)

	}

	if p.Config.SparkPackages != "" {
		p.Config.SparkPackages = fmt.Sprintf("--packages %s", p.Config.SparkPackages)
	}

	if p.Config.SparkExcludePackages != "" {
		p.Config.SparkExcludePackages = fmt.Sprintf("--exclude-packages %s", p.Config.SparkExcludePackages)
	}

	sparkRunCmd := fmt.Sprintf("/opt/spark/bin/spark-submit --verbose "+"--deploy-mode cluster "+
		"--class %s "+
		"--master %s "+
		"--kubernetes-namespace %s "+
		"--conf spark.app.name=%s "+
		"--conf spark.local.dir=%s "+
		"--conf spark.kubernetes.driver.docker.image=%s "+
		"--conf spark.kubernetes.executor.docker.image=%s "+
		"--conf spark.kubernetes.initcontainer.docker.image=%s "+
		"--conf spark.dynamicAllocation.enabled=%s "+
		"--conf spark.kubernetes.resourceStagingServer.uri=%s "+
		"--conf spark.kubernetes.resourceStagingServer.internal.uri=%s "+
		"--conf spark.shuffle.service.enabled=%s "+
		"--conf spark.kubernetes.shuffle.namespace=%s "+
		"--conf spark.kubernetes.shuffle.labels='%s' "+
		"--conf spark.kubernetes.authenticate.driver.serviceAccountName='%s' "+
		"--conf spark.metrics.conf='%s' "+
		"%s "+
		"%s "+
		"%s "+
		"%s "+
		"%s ",
		p.Config.SparkClass,
		p.Config.SparkMaster,
		p.Config.SparkKubernetesNamespace,
		p.Config.SparkAppName,
		p.Config.SparkLocalDir,
		p.Config.KubernetesDriverDockerImage,
		p.Config.KubernetesExecutorDockerImage,
		p.Config.KubernetesInitContainerDockerImage,
		p.Config.SparkDynamicAllocationEnabled,
		p.Config.KubernetesResourceStagingServerUri,
		p.Config.KubernetesResourceStagingServerInternalUri,
		p.Config.SparkShuffleServiceEnabled,
		p.Config.SparkKubernetesShuffleNamespace,
		p.Config.SparkKubernetesShuffleLabels,
		p.Config.KubernetesAuthenticateDriverServiceAccountName,
		p.Config.SparkMetricsConf,
		clintCertAuth,
		p.Config.SparkPackages,
		p.Config.SparkExcludePackages,
		p.Config.SparkAppSource,
		p.Config.SparkAppArgs,
	)

	log.Debugf("Spark Command: %s", sparkRunCmd)
	sparkRunResult := shell.Run(sparkRunCmd)
	log.Infof("Exit code: %d", sparkRunResult.ExitStatus)
	log.Debugf("Stdout: %s", sparkRunResult.Stdout)
	log.Debugf("Stderr: %s", sparkRunResult.Stderr)

	if sparkRunResult.ExitStatus != 0 {
		os.Exit(sparkRunResult.ExitStatus)
	}

	return nil
}
