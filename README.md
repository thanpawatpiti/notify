# Notify

[English](#english) | [ภาษาไทย](#thai)

---

<a name="english"></a>
## English

**Notify** is a professional, secure, and easy-to-use Go library for sending notifications to various messaging platforms. It is designed to be stateless (no hardcoded secrets) and supports rich content (images, colors, embeds) for beautiful notifications.

### Features
- **Multi-Provider Support**: LINE (Messaging API), Telegram, Discord, Microsoft Teams.
- **Secure**: No hardcoded tokens; credentials are passed during initialization.
- **Beautiful**: Supports rich messages (Embeds, Adaptive Cards, Flex/Image messages).
- **Professional**: Built using the Functional Options pattern for flexible configuration (timeouts, custom HTTP clients).
- **Tested**: Fully unit-tested.

### Installation

```bash
go get github.com/thanpawatpiti/notify
```

### Usage

#### 1. Import the library
```go
import (
    "context"
    "time"
    "github.com/thanpawatpiti/notify"
    "github.com/thanpawatpiti/notify/providers/line"
    "github.com/thanpawatpiti/notify/providers/telegram"
    "github.com/thanpawatpiti/notify/providers/discord"
    "github.com/thanpawatpiti/notify/providers/msteams"
)
```

#### 2. Create a Message
```go
msg := notify.Message{
    Title:    "Hello World",
    Content:  "This is a notification from Go!",
    ImageURL: "https://example.com/image.png",
    Color:    "#00FF00", // Supported by Discord
}
```

#### 3. Send Notifications

**LINE (Messaging API)**
```go
// Requires Channel Access Token and UserID/GroupID
lineProvider := line.New("YOUR_CHANNEL_TOKEN", "TARGET_USER_ID")
lineProvider.Send(context.Background(), msg)
```

**Telegram**
```go
telegramProvider := telegram.New("YOUR_BOT_TOKEN", "CHAT_ID")
telegramProvider.Send(context.Background(), msg)
```

**Discord**
```go
discordProvider := discord.New("YOUR_WEBHOOK_URL")
discordProvider.Send(context.Background(), msg)
```

**Microsoft Teams**
```go
teamsProvider := msteams.New("YOUR_WEBHOOK_URL")
teamsProvider.Send(context.Background(), msg)
```

#### 4. Advanced Configuration (Functional Options)
You can configure timeouts or use a custom HTTP client:

```go
// Set a 10-second timeout
p := line.New(token, userID, notify.WithTimeout(10*time.Second))

// Use a custom HTTP client
p := discord.New(webhookURL, notify.WithHTTPClient(myClient))
```

---

<a name="thai"></a>
## ภาษาไทย

**Notify** คือไลบรารีภาษา Go สำหรับส่งการแจ้งเตือนไปยังแพลตฟอร์มต่างๆ ที่มีความเป็นมืออาชีพ ปลอดภัย และใช้งานง่าย ออกแบบมาให้ไม่มีการฝัง Token ไว้ในโค้ด (Stateless) และรองรับการส่งข้อความแบบ Rich Content (รูปภาพ, สี, Embeds) เพื่อความสวยงาม

### คุณสมบัติ
- **รองรับหลายแพลตฟอร์ม**: LINE (Messaging API), Telegram, Discord, Microsoft Teams
- **ปลอดภัย**: ไม่มีการ Hardcode Token หรือ Key ต่างๆ ต้องทำการ Init ฝั่งที่เรียกใช้งาน
- **สวยงาม**: รองรับข้อความแบบ Rich Content เช่น Embeds, Adaptive Cards, และรูปภาพ
- **มืออาชีพ**: ใช้ Functional Options Pattern ในการตั้งค่า (เช่น Timeout, Custom HTTP Client)
- **เชื่อถือได้**: มี Unit Test ครอบคลุม

### การติดตั้ง

```bash
go get github.com/thanpawatpiti/notify
```

### วิธีใช้งาน

#### 1. นำเข้าไลบรารี
```go
import (
    "context"
    "time"
    "github.com/thanpawatpiti/notify"
    "github.com/thanpawatpiti/notify/providers/line"
    "github.com/thanpawatpiti/notify/providers/telegram"
    "github.com/thanpawatpiti/notify/providers/discord"
    "github.com/thanpawatpiti/notify/providers/msteams"
)
```

#### 2. สร้างข้อความ (Message)
```go
msg := notify.Message{
    Title:    "สวัสดีครับ",
    Content:  "นี่คือการแจ้งเตือนจาก Go!",
    ImageURL: "https://example.com/image.png",
    Color:    "#00FF00", // รองรับใน Discord
}
```

#### 3. ส่งการแจ้งเตือน

**LINE (Messaging API)**
```go
// ต้องใช้ Channel Access Token และ UserID/GroupID
lineProvider := line.New("YOUR_CHANNEL_TOKEN", "TARGET_USER_ID")
lineProvider.Send(context.Background(), msg)
```

**Telegram**
```go
telegramProvider := telegram.New("YOUR_BOT_TOKEN", "CHAT_ID")
telegramProvider.Send(context.Background(), msg)
```

**Discord**
```go
discordProvider := discord.New("YOUR_WEBHOOK_URL")
discordProvider.Send(context.Background(), msg)
```

**Microsoft Teams**
```go
teamsProvider := msteams.New("YOUR_WEBHOOK_URL")
teamsProvider.Send(context.Background(), msg)
```

#### 4. การตั้งค่าขั้นสูง (Functional Options)
คุณสามารถตั้งค่า Timeout หรือใช้ HTTP Client ที่กำหนดเองได้:

```go
// ตั้งค่า Timeout 10 วินาที
p := line.New(token, userID, notify.WithTimeout(10*time.Second))

// ใช้ HTTP Client ที่กำหนดเอง
p := discord.New(webhookURL, notify.WithHTTPClient(myClient))
```
