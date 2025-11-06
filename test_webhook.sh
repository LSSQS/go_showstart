#!/bin/bash

# 测试 Webhook 通知脚本
WEBHOOK_URLS=("https://hook.echobell.one/t/lzr4ohbdn38hbp1zp29c" "https://hook.echobell.one/t/fhb0hnji7lwo1396a8cw")

for WEBHOOK_URL in "${WEBHOOK_URLS[@]}"; do
  echo "🧪 测试：杭州氧气音乐节通知 -> ${WEBHOOK_URL}"
  curl -X POST "$WEBHOOK_URL" \
    -H "Content-Type: application/json" \
    -d '{
      "type": "timed",
      "artist": "梓渝",
      "title": "杭州氧气音乐节 11.16 梓渝",
      "showTime": "2025.10.31 12:20 开票",
      "siteName": "杭州氧气音乐节",
      "url": "https://wap.showstart.com/pages/activity/detail/detail?activityId=demo-oxygen-1116"
    }'
done

echo -e "\n\n✅ 测试完成！请检查你的通知接收端。"
