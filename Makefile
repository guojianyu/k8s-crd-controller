# Copyright 2018 Oracle and/or its affiliates. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
Author  := guojianyu

BINARIES := lktest
DOCKER_NAME := lktest
DOCKER_TAG  := latest


.PHONY: build
build:
	go build -o bin/$(BINARIES)  cmd/main.go

.PHONY: docker
docker:
	docker build -t  $(DOCKER_NAME):$(DOCKER_TAG) -f docker/Dockerfile  .
