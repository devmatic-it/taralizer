# Copyright 2021 taralizer authors
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
package rules.asvs
import data.rules.technical_asset_by_id
import data.rules.different_trust_boundaries
import data.rules.calc_impact

violation[{
    "id":id,
     "msg": msg, 
     "likelihood":likelihood,
     "impact": impact}] {
    server := input.technical_assets[_]

    # all containers are effected  
    server.technology == "kubernetes-pod"
    
	id := sprintf("container-baseimage-backdooring@%v", [server.id])
    msg := sprintf("asset '%v' has risk of container image backdooring through base images", [server.id])
    likelihood := 1 
    impact := calc_impact(1) # medium
}
