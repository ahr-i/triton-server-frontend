package setting

/* ----- Server Setting ----- */
const ServerPort string = "80"

const ModelPath string = "./models/model_list.json"
const UrlPath string = "./urls/url_list.json"

/* ----- Triton Server Setting ----- */
const GatewayUrl string = "localhost:2000"

const batchSize int = 1
const Samples int = 1
const Steps int = 45
const GuidanceScale float64 = 7.5
const seed int = -1
