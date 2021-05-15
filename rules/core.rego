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
package rules

technical_asset_by_id (myid) = asset{
    some k
    input.technical_assets[k].id == myid
    asset:= input.technical_assets[k]
}

data_asset_by_id (myid) = asset{
    some k
    input.data_assets[k].id == myid
    asset:= input.data_assets[k]
}

different_trust_boundaries(id1, id2){
    some i, j
    input.trust_boundaries[i].technical_assets_inside[_] == id1
    input.trust_boundaries[j].technical_assets_inside[_] == id2
    i != j
}

same_trust_boundaries(id1, id2){
    some i, j
    input.trust_boundaries[i].technical_assets_inside[_] == id1
    input.trust_boundaries[j].technical_assets_inside[_] == id2
    i == j
}

direct_connection(id1, id2){
    some i
    input.technical_assets[i].id == id1
    input.technical_assets[i].communication_links[_].target == id2
}
