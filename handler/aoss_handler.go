package handler

import (
	"encoding/json"
	"net/http"
)

type Version struct {
	Distribution                     string `json:"distribution"`
	Number                           string `json:"number"`
	BuildType                        string `json:"build_type"`
	BuildHash                        string `json:"build_hash"`
	BuildDate                        string `json:"build_date"`
	BuildSnapshot                    bool   `json:"build_snapshot"`
	LuceneVersion                    string `json:"lucene_version"`
	MinimumWireCompatibilityVersion  string `json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
}

type Info struct {
	Name        string  `json:"name"`
	ClusterName string  `json:"cluster_name"`
	ClusterUuid string  `json:"cluster_uuid"`
	Version     Version `json:"version"`
	TagLine     string  `json:"tag_line"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		version := Version{
			Distribution:                     "aoss",
			Number:                           "2.3.0",
			BuildType:                        "serverless",
			BuildHash:                        "unknown",
			BuildDate:                        "2023-",
			BuildSnapshot:                    false,
			LuceneVersion:                    "8.10.1",
			MinimumWireCompatibilityVersion:  "6.8.0",
			MinimumIndexCompatibilityVersion: "6.0.0-beta1",
		}
		info := Info{
			Name:        "serverless",
			ClusterName: "serverless",
			ClusterUuid: "0",
			Version:     version,
			TagLine:     "The OpenSearch Project: https://opensearch.org/",
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			println("Could not encode info details")
		}
	}
}

type Os struct {
	Name                string `json:"name"`
	Version             string `json:"version"`
	AvailableProcessors string `json:"available_processors"`
}

type Jvm struct {
	VmVendor string `json:"vm_vendor"`
	Version  string `json:"version"`
}

type NodeInfo struct {
	Name string `json:"name"`
	Os   Os     `json:"os"`
	Jvm  Jvm    `json:"jvm"`
}

type NodesInfo struct {
	ClusterName string     `json:"cluster_name"`
	Nodes       []NodeInfo `json:"nodes"`
}

func GetNodesInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		os := Os{
			Name:                "Linux",
			Version:             "1.0.0",
			AvailableProcessors: "1",
		}

		jvm := Jvm{
			VmVendor: "Amazon",
			Version:  "11.0",
		}

		nodeInfo := NodeInfo{
			Name: "node",
			Os:   os,
			Jvm:  jvm,
		}

		nodesInfo := NodesInfo{
			ClusterName: "serverless",
			Nodes:       []NodeInfo{nodeInfo},
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(nodesInfo)
		if err != nil {
			println("Could not encode info details")
		}
	}
}
