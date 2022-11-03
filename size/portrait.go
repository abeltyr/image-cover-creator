package size

import "image/processing/model"

func Portrait(data string) model.SizeDetail {
	xlFrameSize := model.SizeData{
		X: 1475,
		Y: 2180,
	}
	xlImageSize := model.SizeData{
		X: 1410,
		Y: 2115,
	}

	xlFrameOffset := model.SizeData{
		X: 1600,
		Y: 60,
	}
	xlImageOffset := model.SizeData{
		X: 1636,
		Y: 100,
	}

	lgFrameSize := model.SizeData{
		X: 1188,
		Y: 1866,
	}
	lgImageSize := model.SizeData{
		X: 1128,
		Y: 1806,
	}

	lgFrameOffset := model.SizeData{
		X: 1800,
		Y: 400,
	}
	lgImageOffset := model.SizeData{
		X: 1830,
		Y: 430,
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
		X: 625,
		Y: 928,
	}

	smFrameOffset := model.SizeData{
		X: 2150,
		Y: 800,
	}
	smImageOffset := model.SizeData{
		X: 2165,
		Y: 818,
	}

	if data == "sm" {
		return model.SizeDetail{
			FrameSize:   smFrameSize,
			ImageSize:   smImageSize,
			ImageOffset: smImageOffset,
			FrameOffset: smFrameOffset,
		}

	} else if data == "md" {
		return model.SizeDetail{
			FrameSize:   mdFrameSize,
			ImageSize:   mdImageSize,
			ImageOffset: mdImageOffset,
			FrameOffset: mdFrameOffset,
		}
	} else if data == "lg" {
		return model.SizeDetail{
			FrameSize:   lgFrameSize,
			ImageSize:   lgImageSize,
			ImageOffset: lgImageOffset,
			FrameOffset: lgFrameOffset,
		}
	} else {
		return model.SizeDetail{
			FrameSize:   xlFrameSize,
			ImageSize:   xlImageSize,
			ImageOffset: xlImageOffset,
			FrameOffset: xlFrameOffset,
		}
	}

}
