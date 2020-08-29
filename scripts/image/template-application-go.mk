# Copyright 2020 The arhat.dev Authors.
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

# build
image.build.template-application-go.linux.amd64:
	sh scripts/image/build.sh $@

image.build.template-application-go.linux.armv6:
	sh scripts/image/build.sh $@

image.build.template-application-go.linux.armv7:
	sh scripts/image/build.sh $@

image.build.template-application-go.linux.arm64:
	sh scripts/image/build.sh $@

image.build.all: \
	image.build.template-application-go.linux.amd64 \
	image.build.template-application-go.linux.armv6 \
	image.build.template-application-go.linux.armv7 \
	image.build.template-application-go.linux.arm64

# push
image.push.template-application-go.linux.amd64:
	sh scripts/image/push.sh $@

image.push.template-application-go.linux.armv6:
	sh scripts/image/push.sh $@

image.push.template-application-go.linux.armv7:
	sh scripts/image/push.sh $@

image.push.template-application-go.linux.arm64:
	sh scripts/image/push.sh $@

image.push.all: \
	image.push.template-application-go.linux.amd64 \
	image.push.template-application-go.linux.armv6 \
	image.push.template-application-go.linux.armv7 \
	image.push.template-application-go.linux.arm64
