package logging

const REQ_RES_RECV_MSG_FORMAT = "status=${status}#method=${method}#path=${path}#latency=${latency}#reqHeaders=${reqHeaders}#requestId=${respHeader:X-Request-Id}#reqBody=${body}#resBody=${resBody}#error=${error}"
const REQ_RES_LOG_MSG = "[REQ-RES-LOG] %s %s %s %s" // [REQ-RES-LOG] 200 POST /health 2.3sec
