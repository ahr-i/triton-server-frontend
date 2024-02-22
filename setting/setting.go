package setting

/* ----- Server Setting ----- */
const ServerPort string = "80" // Edit this

const ModelPath string = "./models/model_list.json"
const UrlPath string = "./urls/url_list.json"

/* ----- Triton Server Setting ----- */
const GatewayUrl string = "localhost:6000" // Edit this

const Provider string = "ahri"    // Edit this
const BatchSize int = 1           // Edit this
const Samples int = 1             // Edit this
const Steps int = 45              // Edit this
const GuidanceScale float64 = 7.5 // Edit this
