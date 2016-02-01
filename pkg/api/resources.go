package radosAPI

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
