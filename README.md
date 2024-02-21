# Triton Server Frontend
This is the web UI frontend code utilizing the Triton server.  
Used the Text-to-Image feature of Stable Diffusion.   

## 1. Docker Start
### 1.1 Download
```
git clone https://github.com/ahr-i/triton-server-frontend.git
```

### 1.2 setting
```
ce triton-server-frontend
vim setting/setting.go
```
Modify the contents of the file.   
```
package setting

/* ----- Server Setting ----- */
const ServerPort string = "80" // Edit this

const ModelPath string = "./models/model_list.json"
const UrlPath string = "./urls/url_list.json"

/* ----- Triton Server Setting ----- */
const GatewayUrl string = "localhost:6000" // Edit this

const BatchSize int = 1           // Edit this
const Samples int = 1             // Edit this
const Steps int = 45              // Edit this
const GuidanceScale float64 = 7.5 // Edit this
```

### 1.3 Build
```
docker build -t triton-frontend .
```

### 1.4 Run
```
docker run -it --rm --name triton-frontend --network host triton-frontend
or
docker run -it --rm --name triton-frontend -p 80:80 triton-frontend
```
