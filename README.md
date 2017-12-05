# Pipeline Plugin for spark-submit on k8s

This project represents a plugin that can be used to set up a spark-submit step in the Banzai Cloud Pipeline workflow.
A plugin can be configured as a building block of the Banzai Pipeline CI/CD workflow, that can be specified and configured in the Pipeline CI/CD configuration file.

For better understanding of the Banzai Pipeline CI/CD workflow please check [this](https://github.com/banzaicloud/pipeline "Banzai Pipeline")

This plugin implements a fully configurable spark-submit step, as described [here](https://spark.apache.org/docs/latest/submitting-applications.html)

## Usage

For using the plugin, one has to add the plugin configuration step (run) to the Banzai Pipeline CI/CD configuration as it can be seen
 [here](https://github.com/banzaicloud/spark-pi-example/blob/master/.drone.yml) and [here](https://github.com/banzaicloud/spark-pdi-example/blob/master/.drone.yml)
