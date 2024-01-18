package handler

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
	"runtime"
	"strconv"
	"time"

	"github.com/ahr-i/triton-server-front-end/models"
	"github.com/ahr-i/triton-server-front-end/setting"
	"github.com/ahr-i/triton-server-front-end/src/errController"
	"github.com/gorilla/mux"
)

type TritonResponse struct {
	Outputs []struct {
		Data []float32 `json:"data"`
	} `json:"outputs"`
}

type RequestData struct {
	Prompt string `json:"prompt"`
	Seed   string `json:"seed"`
}

func (h *Handler) inferHandler(w http.ResponseWriter, r *http.Request) {
	_, fp, _, _ := runtime.Caller(1)

	var request RequestData
	err_ := json.NewDecoder(r.Body).Decode(&request)
	errController.ErrorCheck(err_, "REQUEST JSON DECODE ERROR", fp)
	defer r.Body.Close()

	vars := mux.Vars(r)
	model := vars["name"]

	log.Println(request.Prompt)
	if request.Prompt == "" || request.Prompt == " " {
		rend.JSON(w, http.StatusBadRequest, nil)
	}

	modelMap := models.GetModelList()
	version, err := modelMap[model]
	if !err {
		rend.JSON(w, http.StatusNotFound, nil)
	}

	seed, _ := strconv.Atoi(request.Seed)
	url := "http://" + setting.TritonUrl + "/v2/models/" + model + "/versions/" + version + "/infer"
	requestData := map[string]interface{}{
		"inputs": []map[string]interface{}{
			{
				"name":     "PROMPT",
				"datatype": "BYTES",
				"shape":    []int{1},
				"data":     []string{request.Prompt},
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
				"data":     []int{seed},
			},
		},
		"outputs": []map[string]string{
			{
				"name": "IMAGES",
			},
		},
	}

	requestJSON, err_ := json.Marshal(requestData)
	errController.ErrorCheck(err_, "JSON MARSHAL ERROR", fp)

	req, err_ := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	errController.ErrorCheck(err_, "HTTP REQUEST ERROR", fp)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err_ := client.Do(req)
	errController.ErrorCheck(err_, "HTTP RESPONSE ERROR", fp)
	defer resp.Body.Close()

	body, err_ := ioutil.ReadAll(resp.Body)
	errController.ErrorCheck(err_, "HTTP BODY READ ERROR", fp)

	var tritonResponse TritonResponse
	if err := json.Unmarshal(body, &tritonResponse); err != nil {
		log.Fatalf("RESPONSE JSON PARSE ERROR: %v", err)
	}

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

		w.Header().Set("Content-Type", "image/png")
		err_ = png.Encode(w, img)
		errController.ErrorCheck(err_, "IMAGE ENCODE ERROR", fp)

		currentTime := time.Now().Format("20060102-150405.999")
		fileName := "result-" + currentTime + ".png"
		file, err := os.Create("./result/" + fileName)
		if err != nil {
			log.Fatalf("이미지 파일 생성 실패: %v", err)
		}
		defer file.Close()

		if err := png.Encode(file, img); err != nil {
			log.Fatalf("이미지 저장 실패: %v", err)
		}

		return
	}
}
