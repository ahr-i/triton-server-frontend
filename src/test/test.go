package test

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type TritonResponse struct {
	Outputs []struct {
		Data []float32 `json:"data"`
	} `json:"outputs"`
}

func Infer() {
	// 모델과 서버 정보 설정
	model_name := "stable_diffusion/versions/1"
	url := "http://localhost:2000/v2/models/" + model_name + "/infer"
	//model_version := "1"

	// HTTP 클라이언트 생성
	//modelMap := models.GetModelList()

	client := &http.Client{}

	// 추론 요청을 위한 JSON 데이터 생성
	requestData := map[string]interface{}{
		// 요청 데이터에 맞는 필드와 값 추가...
		"inputs": []map[string]interface{}{
			{
				"name":     "PROMPT",
				"datatype": "BYTES",
				"shape":    []int{1},
				"data":     []string{"fire truck"},
			},
			{
				"name":     "SAMPLES",
				"datatype": "INT32",
				"shape":    []int{1},
				"data":     []int{1},
			},
			{
				"name":     "STEPS",
				"datatype": "INT32",
				"shape":    []int{1},
				"data":     []int{45},
			},
			{
				"name":     "GUIDANCE_SCALE",
				"datatype": "FP32",
				"shape":    []int{1},
				"data":     []float32{7.5},
			},
			{
				"name":     "SEED",
				"datatype": "INT64",
				"shape":    []int{1},
				"data":     []int{1},
			},
		},
		"outputs": []map[string]string{
			{
				"name": "IMAGES",
			},
		},
	}
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("JSON 마샬링 실패: %v", err)
	}

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Fatalf("HTTP 요청 생성 실패: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// HTTP 요청 전송
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("HTTP 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	// 응답 데이터 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("응답 데이터 읽기 실패: %v", err)
	}

	//log.Println(body)
	// 응답 JSON 파싱
	var tritonResponse TritonResponse
	if err := json.Unmarshal(body, &tritonResponse); err != nil {
		log.Fatalf("응답 JSON 파싱 실패: %v", err)
	}

	// 이미지 데이터 처리
	if len(tritonResponse.Outputs) > 0 && len(tritonResponse.Outputs[0].Data) > 0 {
		imgData := tritonResponse.Outputs[0].Data

		// 이미지의 크기를 가정 (예: 512x512)
		width, height := 512, 512
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		// imgData에서 픽셀 값 추출 및 이미지 생성
		for i := 0; i < len(imgData); i += 3 {
			x := (i / 3) % width
			y := (i / 3) / width
			r := uint8(imgData[i] * 255)
			g := uint8(imgData[i+1] * 255)
			b := uint8(imgData[i+2] * 255)
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}

		// 현재 시간을 밀리세컨드 단위로 포맷
		currentTime := time.Now().Format("20060102-150405.999")

		// 파일 이름 생성
		fileName := "result-" + currentTime + ".png"

		// 이미지 파일로 저장
		file, err := os.Create("./result/" + fileName)
		if err != nil {
			log.Fatalf("이미지 파일 생성 실패: %v", err)
		}
		defer file.Close()

		if err := png.Encode(file, img); err != nil {
			log.Fatalf("이미지 저장 실패: %v", err)
		}

		log.Println("이미지 저장 완료: result.png")
	}
}
