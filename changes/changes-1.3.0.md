Changes by Version
==================
Release Notes.

1.3.0
------------------
#### Features

* Support `SHA256` and `SHA512` in verify case
* Support non-fail-fast execution of test cases
* Support `hasPrefix` and `hasSuffix` string verifier
* support verify cases concurrently
* Add .exe suffix to windows build artifact
* Add a field kubeconfig to support  e2e test on an existing kubernetes cluster
* Export the kubeconfig path during executing the following steps
* Automatically pull images before loading into KinD
* Support outputting the result of 'verify' in YAML format and only outputting the summary of the result of 'verify'
* Make e2e test itself in github action
* Support outputting the summary of 'verify' in YAML format
* Make e2e output summary with numeric information
* Add 'subtractor' function

#### Improvements

* Bump up GHA to avoid too many warnings
* Leverage the built-in cache in setup-go@v4
* Add `batchOutput` config to reduce outputs
* Disable batch mode by default, add it to GHA and enable by default
* Improve GitHub Actions usability and speed by using composite actions' new feature
* Migrate deprecated GitHub Actions command to recommended ones
* Bump up kind to v0.14.0
* Optimization of the output information of  verification
* verifier: notEmpty should be able to handle nil
* Remove invalid configuration in GitHub Actions

#### Bug Fixes

* Fix deprecation warnings
* Ignore cancel error when copying container logs
* Fix kind load docker-image 

#### Documentation

* Add a doc to introduce how to use e2e to test itself

#### Issues and PR
- All issues are [here](https://github.com/apache/skywalking/milestone/148?closed=1)
- All pull requests are [here](https://github.com/apache/skywalking-infra-e2e/milestone/4?closed=1)
