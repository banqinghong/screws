package apollo

type ApplicationConfig struct {
	BaseInfo *BaseInfo     `json:"baseInfo"`
	Items    []*ConfigInfo `json:"items"`
}

type BaseInfo struct {
	ID            int64  `json:"id"`
	AppID         string `json:"appId"`
	ClusterName   string `json:"clusterName"`
	NamespaceName string `json:"namespaceName"`
}

type ConfigInfo struct {
	Item *Item `json:"item"`
}

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
