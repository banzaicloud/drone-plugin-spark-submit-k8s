# Spark submit client plugin for Pipeline CI/CD

This repo contains a plugin that can be used to set up a `spark-submit` step in the Banzai Cloud [Pipeline](https://github.com/banzaicloud/pipeline) workflow.
This plugin is a building block of the Pipeline CI/CD workflow and it's configured using the steps below.

>For better understanding of the Banzai Pipeline CI/CD workflow and PaaS please check [this](https://github.com/banzaicloud/pipeline/README.md) documentation

This plugin implements a fully configurable `spark-submit` step, as described [here](https://spark.apache.org/docs/latest/submitting-applications.html).

## Usage

For using the plugin please configure the `.pipeline.yml` properly, and let the magic happen. 

If you need help configuring the `yml` please read the [Readme](https://github.com/banzaicloud/drone-plugin-pipeline-client) of the related plugin, which handles the cluster related operations.

To configure the `spark-submit` related operations configuration step `run` also has to be added to the `.pipeline.yml`, little help and available options are listed below:

## Main options

| Option                       | Description                                    | Default  | Required |
| -------------                | -----------------------                        | --------:| --------:|
| spark_deploy_mode | Mode the spark should run | cluster | No |
| spark_class  | Main class of the spark application | ""       | Yes |
| spark_kubernetes_local_deploy | Use local kubernetes cluster | true | No |
| spark_kubernetes_namespace | Spark K8S Namespace | default | No |
| spark_app_name | Spark App Name | "" | Yes |
| spark.local.dir | Spark Local Directory | "" | Yes |
| spark_kubernetes_driver_docker_image | Spark K8S Driver Image | "" | Yes |
| spark_kubernetes_executor_docker_image | Spark K8S Executor Image | "" | Yes |
| spark_kubernetes_initcontainer_docker_image | Spark K8S Initcontainer Image | "" | Yes |
| spark_dynamic_allocation | Dynamic Allocation | true | No |
| spark_kubernetes_resourcestagingserver_uri | Spark K8S Resource Staging Server URL | "" | Yes |
| spark_kubernetes_resourcestagingserver_internal_uri | Spark K8S Resource Staging Server Internal URL | "" | Yes |
| spark_shuffle_service_enabled | Spark Shuffle Service Enabled | true | No |
| spark_kubernetes_shuffle_namespace | Spark K8S Shuffle Namespace | default | No |
| spark_kubernetes_shuffle_labels | Spark K8S Shuffle Labels | "" | Yes |
| spark_kubernetes_authenticate_driver_serviceaccount_name | Spark Driver K8s service account name | "" | Yes |
| spark_kubernetes_authenticate_submission_caCertFile | Spark K8S Auth CA CertFile | "" | No |
| spark_kubernetes_authenticate_submission_clientCertFile | Spark K8S Auth Client CertFile | "" | No |
| spark_kubernetes_authenticate_submission_clientKeyFile | Spark K8S Auth Client KeyFile | "" | No |
| spark_metrics_conf | Spark Metrics Config | "" | Yes |
| spark_app_source | Spark App source | "" | Yes |
| spark_app_args | Spark App Args | "" | Yes |
| spark_eventLog_enabled | Spark Event Log Enabled | "" | Yes |

## Cloud provider specific options

If the `spark_eventLog_enabled` set to `yes` there are couple of cloud provider specific options.

### Amazon

| Option                       | Description                                    | Default  | Required |
| -------------                | -----------------------                        | --------:| --------:|
| spark_eventLog_dir | Spark Event Log Directory (s3a://spark-k8-logs/eventLog)| "" | No |

### Azure (AKS)

| Option                       | Description                                    | Default  | Required |
| -------------                | -----------------------                        | --------:| --------:|
| spark_eventLog_dir | Spark Event Log Directory (wasb://...) | "" | No |
| azure_storage_account | Azure storage account | "" | No |
| azure_storage_account_access_key | Azure storage account access key | "" | No |



## Examples

### Spark-Pi

```
 run:
    image: banzaicloud/plugin-k8s-proxy:0.2.0
    original_image: banzaicloud/plugin-spark-submit-k8s:0.2.0
    pod_service_account: spark
    pull: true
    spark_deploy_mode: cluster
    spark_class: banzaicloud.SparkPi
    spark_app_name: sparkpi
    spark_local_dir: /tmp/spark-local
    spark_kubernetes_driver_docker_image: banzaicloud/spark-driver:v2.2.0-k8s-1.0.197
    spark_kubernetes_executor_docker_image: banzaicloud/spark-executor:v2.2.0-k8s-1.0.197
    spark_kubernetes_initcontainer_docker_image: banzaicloud/spark-init:v2.2.0-k8s-1.0.197
    spark_kubernetes_resourcestagingserver_uri: http://spark-rss:10000
    spark_kubernetes_resourcestagingserver_internal_uri: http://spark-rss:10000
    spark_kubernetes_shuffle_labels: "app=spark-shuffle-service,spark-version=2.2.0"
    spark_kubernetes_authenticate_driver_serviceaccount_name: "spark"
    spark_metrics_conf: /opt/spark/conf/metrics.properties
    spark_eventLog_enabled: false
    spark_app_source: target/spark-pi-1.0-SNAPSHOT.jar
    spark_app_args: 1000
```
For the full the configuration file please click [here](https://github.com/banzaicloud/spark-pi-example/blob/master/.pipeline.yml).

### Spark-Pdi

```
 run:
    image: banzaicloud/plugin-k8s-proxy:0.2.0
    original_image: banzaicloud/plugin-spark-submit-k8s:0.2.0
    pod_service_account: spark
    pull: true
    spark_deploy_mode: cluster
    spark_class: com.banzaicloud.sfdata.SFPDIncidents
    spark_app_name: SFPDIncidents
    spark_local_dir: /tmp/spark-local
    spark_kubernetes_driver_docker_image: banzaicloud/spark-driver:v2.2.0-k8s-1.0.197
    spark_kubernetes_executor_docker_image: banzaicloud/spark-executor:v2.2.0-k8s-1.0.197
    spark_kubernetes_initcontainer_docker_image: banzaicloud/spark-init:v2.2.0-k8s-1.0.197
    spark_kubernetes_resourcestagingserver_uri: http://spark-rss:10000
    spark_kubernetes_resourcestagingserver_internal_uri: http://spark-rss:10000
    spark_kubernetes_shuffle_labels: "app=spark-shuffle-service,spark-version=2.2.0"
    spark_kubernetes_authenticate_driver_serviceaccount_name: "spark"
    spark_metrics_conf: /opt/spark/conf/metrics.properties
    spark_eventLog_enabled: false
    spark_app_source: target/scala-2.11/sf-police-incidents_2.11-0.1.jar
    spark_app_args: --dataPath s3a://lp-deps-test/data/Police_Department_Incidents.csv
```

For the full the configuration file please click [here](https://github.com/banzaicloud/spark-pdi-example/blob/master/.pipeline.yml).
