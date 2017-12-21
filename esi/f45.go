/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.3
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetCharactersCharacterIdBlueprints200Ok. */
//easyjson:json
type GetCharactersCharacterIdBlueprints200OkList []GetCharactersCharacterIdBlueprints200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdBlueprints200Ok struct {
	ItemId             int64  `json:"item_id,omitempty"`             /* Unique ID for this item. */
	TypeId             int32  `json:"type_id,omitempty"`             /* type_id integer */
	LocationId         int64  `json:"location_id,omitempty"`         /* References a solar system, station or item_id if this blueprint is located within a container. If the return value is an item_id, then the Character AssetList API must be queried to find the container using the given item_id to determine the correct location of the Blueprint. */
	LocationFlag       string `json:"location_flag,omitempty"`       /* Type of the location_id */
	Quantity           int32  `json:"quantity,omitempty"`            /* A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet). */
	TimeEfficiency     int32  `json:"time_efficiency,omitempty"`     /* Time Efficiency Level of the blueprint. */
	MaterialEfficiency int32  `json:"material_efficiency,omitempty"` /* Material Efficiency Level of the blueprint. */
	Runs               int32  `json:"runs,omitempty"`                /* Number of runs remaining if the blueprint is a copy, -1 if it is an original. */
}
