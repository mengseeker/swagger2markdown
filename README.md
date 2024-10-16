# swagger2markdown

## Install
```shell
go install github.com/mengseeker/swagger2markdown@latest
```

## Usage
```text
transform swagger into markdown.
only support swagger 2.0

Usage:
  swagger2markdown [flags]

Flags:
  -h, --help                 help for swagger2markdown
  -i, --input string         input file, can be json or yaml format, default read from stdin
  -f, --inputFormat string   input file format, json or yaml, default auto detect
  -o, --output string        output file, default print to stdout
  -m, --template string      custom template file
  -t, --toggle               Help message for toggle
```

## Example
```txt
# api/api.proto (version not set)

**Schemes:** http https 

---

## GET /v1/fetch_job

### Summary
获取挖掘作业


### Parameters
#### Query parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| workerId | string | false | worker_id 工作节点id |
| host | string | false | 工作节点host |
| ip | string | false | 工作节点ip |
| version | string | false | 工作节点版本，用于判断当前工作节点是否同服务版本一致，不一致时，worker会自动退出 |

### Responses

#### 200
A successful response.

| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| job | object | false |  |
| job.info | object | false |  |
| job.info.mapS3Path | string | false | 地图路径 |
| job.info.s3SaveDir | string | false | 挖掘结果存储位置 |
| job.info.jobId | string(uint64) | false | job_id |
| job.info.deviceId | string(uint64) | false | 设备id |
| job.info.dataStartAt | string(int64) | false | 数据分片开始时间 |
| job.info.dataEndAt | string(int64) | false | 数据分片结束时间 |
| job.info.s3 | object | false |  |
| job.info.s3.bucket | string | false | s3 bucket配置 |
| job.info.s3.endpoint | string | false | s3 访问地址 |
| job.info.s3.accessKey | string | false | s3 ak |
| job.info.s3.secretKey | string | false | s3 sk |
| job.info.config | object | false | 挖掘任务配置 |
| job.info.config.dataGeoType | string | false | 数据geo类型，用于位置纠偏 |
| job.info.config.tags | array | false | 算子 |
| job.info.config.tags[] | object | false |  |
| job.info.config.tags[].Type | string | false | 算子类型 |
| job.info.config.tags[].Name | string | false | 算子名称 |
| job.timeout | string(int64) | false | 任务执行超时 |
| hasExpired | boolean | false | 节点是否过期 |

#### default
An unexpected error response.

| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| details | array | false |  |
| details[] | object | false |  |
| details[].@type | string | false |  |
| code | integer(int32) | false |  |
| message | string | false |  |

```