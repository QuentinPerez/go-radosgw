package radosAPI

type apiError struct {
	Code string `json:"Code"`
}

// Usage represents the response of usage requests
type Usage struct {
	Entries []struct {
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
	} `json:"entries"`
	Summary []struct {
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
	} `json:"summary"`
}

// User represents the response of user requests
type User struct {
	Caps []struct {
		Perm string `json:"perm"`
		Type string `json:"type"`
	} `json:"caps"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Keys        []struct {
		AccessKey string `json:"access_key"`
		SecretKey string `json:"secret_key"`
		User      string `json:"user"`
	} `json:"keys"`
	MaxBuckets int `json:"max_buckets"`
	Subusers   []struct {
		ID          string `json:"id"`
		Permissions string `json:"permissions"`
	} `json:"subusers"`
	Suspended int `json:"suspended"`
	SwiftKeys []struct {
		SecretKey string `json:"secret_key"`
		User      string `json:"user"`
	} `json:"swift_keys"`
	UserID string `json:"user_id"`
}
