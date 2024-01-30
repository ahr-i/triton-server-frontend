# Triton Server Frontend
This is the web UI frontend code utilizing the Triton server.  
Used the Text-to-Image feature of Stable Diffusion.   

## 1. Docker Start
### 1.1 Clone
```
git clone https://github.com/ahr-i/triton-server-frontend.git
```

### 1.2 build
```
cd triton-server-frontend
docker build -t triton-frontend .
```

### 1.3 setting
```
vim setting/setting.go
```
Modify the contents of the file.   
```
package setting

/* ----- Server Setting ----- */
const ServerPort string = "443"                      // Edit This

const ModelPath string = "./models/model_list.json"
const UrlPath string = "./urls/url_list.json"

/* ----- Triton Server Setting ----- */
const TritonUrl string = "localhost:2000"            // Edit This

const batchSize int = 1
const Samples int = 1
const Steps int = 45
const GuidanceScale float64 = 7.5
const seed int = -1
```

### 1.4 Run
```
docker run -it --rm --name triton_frontend --network host triton-frontend
```
