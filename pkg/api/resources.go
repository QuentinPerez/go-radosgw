package radosAPI

type apiError struct {
	Code string `json:"Code"`
}

type Entry struct {
	Buckets []struct {
		Bucket     string `json:"bucket"`
		Categories []struct {
			BytesReceived int    `json:"bytes_received"`
			BytesSent     int    `json:"bytes_sent"`
			Category      string `json:"category"`
			Ops           int    `json:"ops"`
			SuccessfulOps int    `json:"successful_ops"`
		} `json:"categories"`
		Epoch int    `json:"epoch"`
		Time  string `json:"time"`
	} `json:"buckets"`
	Owner string `json:"owner"`
}

type Summary struct {
	Categories []struct {
		BytesReceived int    `json:"bytes_received"`
		BytesSent     int    `json:"bytes_sent"`
		Category      string `json:"category"`
		Ops           int    `json:"ops"`
		SuccessfulOps int    `json:"successful_ops"`
	} `json:"categories"`
	Total struct {
		BytesReceived int `json:"bytes_received"`
		BytesSent     int `json:"bytes_sent"`
		Ops           int `json:"ops"`
		SuccessfulOps int `json:"successful_ops"`
	} `json:"total"`
	User string `json:"user"`
}

// Usage represents the response of usage requests
type Usage struct {
	Entries []Entry   `json:"entries"`
	Summary []Summary `json:"summary"`
}

// SubUsers represents the response of subuser requests
type SubUsers []struct {
	ID          string `json:"id"`
	Permissions string `json:"permissions"`
}

// KeysDefinition represents the response of key requests
type KeysDefinition []struct {
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key"`
	User      string `json:"user"`
}

// User represents the response of user requests
type User struct {
	Caps        []Capability   `json:"caps"`
	DisplayName string         `json:"display_name"`
	Email       string         `json:"email"`
	Keys        KeysDefinition `json:"keys"`
	MaxBuckets  int            `json:"max_buckets"`
	Subusers    SubUsers       `json:"subusers"`
	Suspended   int            `json:"suspended"`
	SwiftKeys   KeysDefinition `json:"swift_keys"`
	UserID      string         `json:"user_id"`
}

type Stats struct {
	Bucket      string `json:"bucket"`
	BucketQuota struct {
		Enabled    bool `json:"enabled"`
		MaxObjects int  `json:"max_objects"`
		MaxSizeKb  int  `json:"max_size_kb"`
	} `json:"bucket_quota"`
	ID        string `json:"id"`
	IndexPool string `json:"index_pool"`
	Marker    string `json:"marker"`
	MasterVer string `json:"master_ver"`
	MaxMarker string `json:"max_marker"`
	Mtime     string `json:"mtime"`
	Owner     string `json:"owner"`
	Pool      string `json:"pool"`
	Usage     struct {
		RgwMain struct {
			NumObjects   int `json:"num_objects"`
			SizeKb       int `json:"size_kb"`
			SizeKbActual int `json:"size_kb_actual"`
		} `json:"rgw.main"`
	} `json:"usage"`
	Ver string `json:"ver"`
}

type Bucket struct {
	Name  string `json:"name,omitempty"`
	Stats *Stats `json:"stats,omitempty"`
}

// Buckets represents the response of bucket requests
type Buckets []Bucket

// Policy represents the response of policy requests
type Policy struct {
	Acl struct {
		AclGroupMap []struct {
			Acl   int `json:"acl"`
			Group int `json:"group"`
		} `json:"acl_group_map"`
		AclUserMap []struct {
			Acl  int    `json:"acl"`
			User string `json:"user"`
		} `json:"acl_user_map"`
		GrantMap []struct {
			Grant struct {
				Email      string `json:"email"`
				Group      int    `json:"group"`
				ID         string `json:"id"`
				Name       string `json:"name"`
				Permission struct {
					Flags int `json:"flags"`
				} `json:"permission"`
				Type struct {
					Type int `json:"type"`
				} `json:"type"`
			} `json:"grant"`
			ID string `json:"id"`
		} `json:"grant_map"`
	} `json:"acl"`
	Owner struct {
		DisplayName string `json:"display_name"`
		ID          string `json:"id"`
	} `json:"owner"`
}

// Quotas represents the reponse of quotas requests
type Quotas struct {
	BucketQuota struct {
		Enabled    bool `json:"enabled"`
		MaxObjects int  `json:"max_objects"`
		MaxSizeKb  int  `json:"max_size_kb"`
	} `json:"bucket_quota"`
	UserQuota struct {
		Enabled    bool `json:"enabled"`
		MaxObjects int  `json:"max_objects"`
		MaxSizeKb  int  `json:"max_size_kb"`
	} `json:"user_quota"`
}

// Capability represents the reponse of capability requests
type Capability struct {
	Perm string `json:"perm"`
	Type string `json:"type"`
}
