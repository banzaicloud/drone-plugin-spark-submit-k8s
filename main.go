package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version string = ""
)

func main() {

	app := cli.NewApp()
	app.Name = "spark-k8s plugin"
	app.Usage = "spark-k8s plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s", version)

	app.Flags = []cli.Flag{

		//
		// plugin args
		//

		cli.StringFlag{
			Name:   "plugin.spark.deploy.mode",
			Usage:  "Spark Deploy Mode",
			EnvVar: "PLUGIN_SPARK_DEPLOY_MODE",
			Value:  "cluster",
		},
		cli.StringFlag{
			Name:   "plugin.spark.class",
			Usage:  "Spark Class",
			EnvVar: "PLUGIN_SPARK_CLASS",
		},
		cli.StringFlag{
			Name:   "plugin.spark.kubernetes.local.deploy",
			Usage:  "Use local kubernetes cluster (true|false)",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_LOCAL_DEPLOY",
			Value:  "true",
		},
		cli.StringFlag{
			Name:   "plugin.spark.kubernetes.local.url",
			Usage:  "Kubernetes API Local Deploy URL",
			EnvVar: "KUBERNETES_PORT_443_TCP_ADDR",
		},
		cli.StringFlag{
			Name:   "plugin.spark.kubernetes.local.port",
			Usage:  "Kubernetes API Local Deploy Port",
			EnvVar: "KUBERNETES_SERVICE_PORT_HTTPS",
		},
		cli.StringFlag{
			Name:   "plugin.spark.master",
			Usage:  "Spark K8S cluster URL",
			EnvVar: "PLUGIN_SPARK_MASTER",
		},
		cli.StringFlag{
			Name:   "plugin.spark.kubernetes.namespace",
			Usage:  "Spark K8S Namespace",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_NAMESPACE",
			Value:  "default",
		},
		cli.StringFlag{
			Name:   "plugin.spark.app.name",
			Usage:  "Spark App Name",
			EnvVar: "PLUGIN_SPARK_APP_NAME",
		},
		cli.StringFlag{
			Name:   "plugin.spark.local.dir",
			Usage:  "Spark Local Directory",
			EnvVar: "PLUGIN_SPARK_LOCAL_DIR",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.driver.docker.image",
			Usage:  "Spark K8S Driver Image",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_DRIVER_DOCKER_IMAGE",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.executor.docker.image",
			Usage:  "Spark K8S Executor Image",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_EXECUTOR_DOCKER_IMAGE",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.initcontainer.docker.image",
			Usage:  "Spark K8S Initcontainer Image",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_INITCONTAINER_DOCKER_IMAGE",
		},
		cli.StringFlag{
			Name:   "plugin.spark.dynamicAllocation.enabled",
			Usage:  "Dynamic Allocation",
			EnvVar: "PLUGIN_SPARK_DYNAMIC_ALLOCATION",
			Value:  "true",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.resourceStagingServer.uri",
			Usage:  "Spark K8S Resource Staging Server URL",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_RESOURCESTAGINGSERVER_URI",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.resourceStagingServer.internal.uri",
			Usage:  "Spark K8S Resource Staging Server Internal URL",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_RESOURCESTAGINGSERVER_INTERNAL_URI",
		},
		cli.StringFlag{
			Name:   "plugin.spark.shuffle.service.enabled",
			Usage:  "Spark Shuffle Service Enabled",
			EnvVar: "PLUGIN_SPARK_SHUFFLE_SERVICE_ENABLED",
			Value:  "true",
		},
		cli.StringFlag{
			Name:   "plugin.spark.kubernetes.shuffle.namespace",
			Usage:  "Spark K8S Shuffle Namespace",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_SHUFFLE_NAMESPACE",
			Value:  "default",
		},
		cli.StringFlag{
			Name:   "plugin.spark.kubernetes.shuffle.labels",
			Usage:  "Spark K8S Shuffle Labels",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_SHUFFLE_LABELS",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.authenticate.driver.serviceAccountName",
			Usage:  "Spark Driver K8s service account name.",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_AUTHENTICATE_DRIVER_SERVICEACCOUNT_NAME",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.authenticate.submission.caCertFile",
			Usage:  "Spark K8S Auth CA CertFile",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_AUTHENTICATE_SUBMISSION_CACERTFILE",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.authenticate.submission.clientCertFile",
			Usage:  "Spark K8S Auth Client CertFile",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_AUTHENTICATE_SUBMISSION_CLIENTCERTFILE",
		},
		cli.StringFlag{
			Name:   "plugin.kubernetes.authenticate.submission.clientKeyFile",
			Usage:  "Spark K8S Auth Client KeyFile",
			EnvVar: "PLUGIN_SPARK_KUBERNETES_AUTHENTICATE_SUBMISSION_CLIENTKEYFILE",
		},
		cli.StringFlag{
			Name:   "plugin.spark.packages",
			Usage:  "Spark Packages",
			EnvVar: "PLUGIN_SPARK_PACKAGES",
		},
		cli.StringFlag{
			Name:   "plugin.spark.exclude-packages",
			Usage:  "Spark Exclude Packages",
			EnvVar: "PLUGIN_SPARK_EXCLUDE-PACKAGES",
		},
		cli.StringFlag{
			Name:   "plugin.spark.app.source",
			Usage:  "Spark App source",
			EnvVar: "PLUGIN_SPARK_APP_SOURCE",
		},
		cli.StringFlag{
			Name:   "plugin.spark.app.args",
			Usage:  "Spark App Args",
			EnvVar: "PLUGIN_SPARK_APP_ARGS",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Repo: Repo{
			Owner:   c.String("repo.owner"),
			Name:    c.String("repo.name"),
			Link:    c.String("repo.link"),
			Avatar:  c.String("repo.avatar"),
			Branch:  c.String("repo.branch"),
			Private: c.Bool("repo.private"),
			Trusted: c.Bool("repo.trusted"),
		},
		Config: Config{
			SparkDeployMode:                                c.String("plugin.spark.deploy.mode"),
			SparkClass:                                     c.String("plugin.spark.class"),
			SparkKubernetesLocalDeploy:                     c.String("plugin.spark.kubernetes.local.deploy"),
			SparkKubernetesLocalUrl:                        c.String("plugin.spark.kubernetes.local.url"),
			SparkKubernetesLocalPort:                       c.String("plugin.spark.kubernetes.local.port"),
			SparkMaster:                                    c.String("plugin.spark.master"),
			SparkKubernetesNamespace:                       c.String("plugin.spark.kubernetes.namespace"),
			SparkAppName:                                   c.String("plugin.spark.app.name"),
			SparkLocalDir:                                  c.String("plugin.spark.local.dir"),
			KubernetesDriverDockerImage:                    c.String("plugin.kubernetes.driver.docker.image"),
			KubernetesExecutorDockerImage:                  c.String("plugin.kubernetes.executor.docker.image"),
			KubernetesInitContainerDockerImage:             c.String("plugin.kubernetes.initcontainer.docker.image"),
			SparkDynamicAllocationEnabled:                  c.String("plugin.spark.dynamicAllocation.enabled"),
			KubernetesResourceStagingServerUri:             c.String("plugin.kubernetes.resourceStagingServer.uri"),
			KubernetesResourceStagingServerInternalUri:     c.String("plugin.kubernetes.resourceStagingServer.internal.uri"),
			SparkShuffleServiceEnabled:                     c.String("plugin.spark.shuffle.service.enabled"),
			SparkKubernetesShuffleNamespace:                c.String("plugin.spark.kubernetes.shuffle.namespace"),
			SparkKubernetesShuffleLabels:                   c.String("plugin.spark.kubernetes.shuffle.labels"),
			KubernetesAuthenticateDriverServiceAccountName: c.String("plugin.kubernetes.authenticate.driver.serviceAccountName"),
			KubernetesAuthenticateSubmissionCaCertFile:     c.String("plugin.kubernetes.authenticate.submission.caCertFile"),
			KubernetesAuthenticateSubmissionClientCertFile: c.String("plugin.kubernetes.authenticate.submission.clientCertFile"),
			KubernetesAuthenticateSubmissionClientKeyFile:  c.String("plugin.kubernetes.authenticate.submission.clientKeyFile"),
			SparkPackages:                                  c.String("plugin.spark.packages"),
			SparkExcludePackages:                           c.String("plugin.spark.exclude-packages"),
			SparkAppSource:                                 c.String("plugin.spark.app.source"),
			SparkAppArgs:                                   c.String("plugin.spark.app.args"),
		},
	}

	if err := plugin.Exec(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
