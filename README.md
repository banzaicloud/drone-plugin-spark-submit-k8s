# Pipeline Plugin for spark-submit on Kubernetes

This project contains a plugin that can be used to set up a `spark-submit` step in the Banzai Cloud [Pipeline](https://github.com/banzaicloud/pipeline) workflow.
This plugin can be configured as a building block of the Banzai Pipeline CI/CD workflow and it can be specified/used and configured in the Pipeline CI/CD configuration file.

For better understanding of the Banzai Pipeline CI/CD workflow and PaaS please check [this](https://github.com/banzaicloud/pipeline/README.md) documentation

This plugin implements a fully configurable `spark-submit` step, as described [here](https://spark.apache.org/docs/latest/submitting-applications.html)

## Usage

For using the plugin the plugin configuration step (run) has to be added to the Banzai Pipeline CI/CD configuration as it can be seen in the following examples 
 [here](https://github.com/banzaicloud/spark-pi-example/blob/master/.pipeline.yml) and [here](https://github.com/banzaicloud/spark-pdi-example/blob/master/.pipeline.yml)
