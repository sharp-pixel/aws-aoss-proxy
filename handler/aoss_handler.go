package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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
	log.Info("Intercepted /")

	version := Version{
		Distribution:                     "aoss",
		Number:                           "2.3.0",
		BuildType:                        "serverless",
		BuildHash:                        "unknown",
		BuildDate:                        "2023-01-21T00:00:00.000000Z",
		BuildSnapshot:                    false,
		LuceneVersion:                    "9.3.0",
		MinimumWireCompatibilityVersion:  "7.10.0",
		MinimumIndexCompatibilityVersion: "7.0.0",
	}
	info := Info{
		Name:        "serverless",
		ClusterName: "serverless",
		ClusterUuid: "0",
		Version:     version,
		TagLine:     "The OpenSearch Project: https://opensearch.org/",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		println("Could not encode info details")
	}

	b, error := json.MarshalIndent(info, "", "  ")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Println(string(b))
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
	log.Info("Intercepted _nodes/stats")

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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(nodesInfo)
	if err != nil {
		println("Could not encode info details")
	}
}

type Health struct {
	ClusterName                 string  `json:"cluster_name"`
	Status                      string  `json:"status"`
	TimedOut                    bool    `json:"timed_out"`
	NumberOfNodes               int     `json:"number_of_nodes"`
	NumberOfDataNodes           int     `json:"number_of_data_nodes"`
	ActivePrimaryShards         int     `json:"active_primary_shards"`
	ActiveShards                int     `json:"active_shards"`
	RelocatingShards            int     `json:"relocating_shards"`
	InitializingShards          int     `json:"initializing_shards"`
	UnassignedShards            int     `json:"unassigned_shards"`
	DelayedUnassignedShards     int     `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        int     `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       int     `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis int     `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercentAsNumber float64 `json:"active_shards_percent_as_number"`
}

func GetHealthInfo(w http.ResponseWriter, r *http.Request) {
	log.Info("Intercepted _cluster/health")

	health := Health{
		ClusterName:                 "serverless",
		Status:                      "green",
		TimedOut:                    false,
		NumberOfNodes:               1,
		NumberOfDataNodes:           1,
		ActivePrimaryShards:         1,
		ActiveShards:                1,
		RelocatingShards:            0,
		InitializingShards:          0,
		UnassignedShards:            0,
		DelayedUnassignedShards:     0,
		NumberOfPendingTasks:        0,
		NumberOfInFlightFetch:       0,
		TaskMaxWaitingInQueueMillis: 0,
		ActiveShardsPercentAsNumber: 1.0,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(health)
	if err != nil {
		println("Could not encode info details")
	}
}
