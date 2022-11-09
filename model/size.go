package model

type SizeData struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SizeDetail struct {
	Title       string   `json:"title"`
	Name        string   `json:"name"`
	FrameSize   SizeData `json:"frameSize"`
	ImageSize   SizeData `json:"imageSize"`
	FrameOffset SizeData `json:"frameOffset"`
	ImageOffset SizeData `json:"imageOffset"`
}
