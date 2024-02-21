package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ahr-i/triton-server-frontend/models"
	"github.com/ahr-i/triton-server-frontend/setting"
	"github.com/gorilla/mux"
)

/* Response Struct */
type GatewayResponse struct {
	Outputs []struct {
		Data []float32 `json:"data"`
	} `json:"outputs"`
}

type ResponseData struct {
	Image string `json:"image"`
}

/* Request Struct */
type RequestData struct {
	Prompt string `json:"prompt"`
}

/* Inference Handler: Gateway Server에 Inference Request 및 Image 전달 */
func (h *Handler) inferHandler(w http.ResponseWriter, r *http.Request) {
	// Request Decode
	var request RequestData
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	model := vars["name"]

	// log.Println(request.Prompt)
	if request.Prompt == "" || request.Prompt == " " {
		rend.JSON(w, http.StatusBadRequest, nil)

		return
	}

	// Model, Version Check And Setting
	modelMap := models.GetModelList()
	version, err_ := modelMap[model]
	if !err_ {
		rend.JSON(w, http.StatusNotFound, nil)

		return
	}

	// Gateway Request
	gatewayResponse, err := requestGatewayServerResponse(request, model, version)
	if err != nil {
		log.Println("** (ERROR)", err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	// Uint8 Array To Image
	img, err_ := converUint8ToPng(gatewayResponse.Outputs[0].Data)
	if err_ {
		// Inference fail
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	// Encode the image in base64
	imgBase64, err := encodeImageInBase64(img)
	if err != nil {
		log.Println("** (ERROR)", err)
		return
	}

	// Response에 Image 추가
	rend.JSON(w, http.StatusOK, ResponseData{Image: imgBase64})

	// Image Local 저장
	saveImageToLocal(img)
}

/* Send an inference request to the Gateway Server and return the request */
func requestGatewayServerResponse(request RequestData, model string, version string) (GatewayResponse, error) {
	rand.Seed(time.Now().UnixNano())

	// Gateway Inference Request
	seed := rand.Intn(10001)
	url := fmt.Sprintf("http://%s/provider/%s/model/%s/%s/infer", setting.GatewayUrl, setting.Provider, model, version)
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
				"data":     []int{setting.Samples},
			},
			{
				"name":     "STEPS",
				"datatype": "INT32",
				"shape":    []int{1},
				"data":     []int{setting.Steps},
			},
			{
				"name":     "GUIDANCE_SCALE",
				"datatype": "FP32",
				"shape":    []int{1},
				"data":     []float32{float32(setting.GuidanceScale)},
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

	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		return GatewayResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return GatewayResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Gateway Server Response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GatewayResponse{}, err
	}
	defer resp.Body.Close()

	// Response Decode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GatewayResponse{}, err
	}

	var gatewayResponse GatewayResponse
	if err := json.Unmarshal(body, &gatewayResponse); err != nil {
		return GatewayResponse{}, err
	}

	return gatewayResponse, nil
}

/* Convert Uint8 To Image(PNG) */
func converUint8ToPng(imgData []float32) (*image.RGBA, bool) {
	if len(imgData) <= 0 {
		return nil, true
	}

	// Image의 크기 가정
	width, height := 512, 512
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// ImgData에서 픽셀 값 추출 및 Image 생성
	for i := 0; i < len(imgData); i += 3 {
		x := (i / 3) % width
		y := (i / 3) / width
		r := uint8(imgData[i] * 255)
		g := uint8(imgData[i+1] * 255)
		b := uint8(imgData[i+2] * 255)
		img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
	}

	return img, false
}

/* Encode the image in base64 */
func encodeImageInBase64(img *image.RGBA) (string, error) {
	var buffer bytes.Buffer

	if err := png.Encode(&buffer, img); err != nil {
		return "", err
	}
	imgBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return imgBase64, nil
}

/* Save image to local */
func saveImageToLocal(img *image.RGBA) error {
	currentTime := time.Now().Format("20060102-150405.999")
	fileName := "result-" + currentTime + ".png"
	file, err := os.Create("./result/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}
