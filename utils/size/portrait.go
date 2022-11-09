package size

import "image/process/model"

func Portrait(data string) model.SizeDetail {

	xlFrameSize := model.SizeData{
		X: 1490,
		Y: 2180,
	}
	xlImageSize := model.SizeData{
		X: 1410,
		Y: 2100,
	}

	xlFrameOffset := model.SizeData{
		X: 1600,
		Y: 60,
	}
	xlImageOffset := model.SizeData{
		X: 1640,
		Y: 100,
	}

	lgFrameSize := model.SizeData{
		X: 1188,
		Y: 1866,
	}
	lgImageSize := model.SizeData{
		X: 1118,
		Y: 1786,
	}

	lgFrameOffset := model.SizeData{
		X: 1800,
		Y: 400,
	}
	lgImageOffset := model.SizeData{
		X: 1840,
		Y: 440,
	}

	mdFrameSize := model.SizeData{
		X: 990,
		Y: 1460,
	}
	mdImageSize := model.SizeData{
		X: 940,
		Y: 1410,
	}

	mdFrameOffset := model.SizeData{
		X: 1900,
		Y: 500,
	}
	mdImageOffset := model.SizeData{
		X: 1925,
		Y: 525,
	}

	smFrameSize := model.SizeData{
		X: 655,
		Y: 964,
	}
	smImageSize := model.SizeData{
		X: 615,
		Y: 924,
	}

	smFrameOffset := model.SizeData{
		X: 2150,
		Y: 800,
	}
	smImageOffset := model.SizeData{
		X: 2170,
		Y: 820,
	}

	if data == "sm" {
		return model.SizeDetail{
			Title:       "30x40",
			Name:        "Frame/30x40.png",
			FrameSize:   smFrameSize,
			ImageSize:   smImageSize,
			ImageOffset: smImageOffset,
			FrameOffset: smFrameOffset,
		}

	} else if data == "md" {
		return model.SizeDetail{
			Title:       "40x60",
			Name:        "Frame/40x60.png",
			FrameSize:   mdFrameSize,
			ImageSize:   mdImageSize,
			ImageOffset: mdImageOffset,
			FrameOffset: mdFrameOffset,
		}
	} else if data == "lg" {
		return model.SizeDetail{
			Title:       "50x80",
			Name:        "Frame/50x80.png",
			FrameSize:   lgFrameSize,
			ImageSize:   lgImageSize,
			ImageOffset: lgImageOffset,
			FrameOffset: lgFrameOffset,
		}
	} else {
		return model.SizeDetail{
			Title:       "60x90",
			Name:        "Frame/60x90.png",
			FrameSize:   xlFrameSize,
			ImageSize:   xlImageSize,
			ImageOffset: xlImageOffset,
			FrameOffset: xlFrameOffset,
		}
	}

}
