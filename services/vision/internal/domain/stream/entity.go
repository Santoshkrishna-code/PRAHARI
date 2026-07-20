package stream

// Config represents stream parameter protocols (RTSP, WebRTC, WebSocket).
type Config struct {
	ID        string `json:"id"`
	CameraID  string `json:"camera_id"`
	RTSPUrl   string `json:"rtsp_url"`
	FPS       int    `json:"fps"`
	Resolution string `json:"resolution"` // E.g., 1080p, 4k
}
