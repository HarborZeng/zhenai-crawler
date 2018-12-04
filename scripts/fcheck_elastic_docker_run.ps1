# encoding: utf-8
# auther: ZengXiaogang email: s_zengxiaogang@wps.cn

$CONTAINER_NAME = "zhenai_crawler_elastic"

# 删除掉已存在的名为“zhenai_crawler_elastic”的container
"checking whether same name container exists"
$sameNameContainer = docker container ls --filter "NAME=^$CONTAINER_NAME$" -aq
if ($null -eq $sameNameContainer) {
    "No same name container exists"
} else {
    "deleting same name container"
    $DELETE_ID = docker container rm $sameNameContainer -f
    if ($LASTEXITCODE -ne 0) {
        "failed"
        return
    } else {
        "success: $DELETE_ID"
    }
}

# 创建一个新的名为“zhenai_crawler_elastic”的container
"creating container named $CONTAINER_NAME"
$CONTAINER_DATA = $CONTAINER_NAME + "_data"
$id = docker container run --name $CONTAINER_NAME `
-p 9200:9200 `
-e "discovery.type=single-node" `
-d `
elasticsearch:6.5.1

if ($null -ne $id -and $?) {
    "success: $id"
    return
}
"failed"