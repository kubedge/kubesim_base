# Copyright 2018 Kubedge.
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

#FROM arm32v7/python
FROM don41382/rpi-python3-with-bluetooth
RUN apt-get update && apt-get install -y \
    bluetooth \
    bluez \
    bluez-tools \
    python-bluez
ADD utilities/detect_bluetooth.py /detect_bluetooth.py
EXPOSE 50051
CMD ["python","detect_bluetooth.py"]
